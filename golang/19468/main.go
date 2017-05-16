package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/DataDog/zstd"
	"github.com/Microsoft/go-winio"
)

func NewClient(name string) (net.Conn, error) {
	connectionTimeout := 2 * time.Second
	return winio.DialPipe(`\\.\pipe\`+name, &connectionTimeout)
}

type Terminator struct {
	srv        net.Listener
	signalsMtx sync.RWMutex
	signals    []func()
}

func NewTerminator(name string) *Terminator {
	ret := Terminator{}

	srv, err := winio.ListenPipe(`\\.\pipe\`+name, nil)
	if err != nil {
		panic(err)
	}
	ret.srv = srv

	go ret.worker()
	return &ret
}

func (this *Terminator) worker() {
	conn, err := this.srv.Accept()
	if err != nil {
		panic(err)
	}

	_ = conn.Close()

	this.Signal()
}

func (this *Terminator) RegisterSignalHandler(cb func()) {
	this.signalsMtx.Lock()
	defer this.signalsMtx.Unlock()

	this.signals = append(this.signals, cb)
}

func (this *Terminator) Signal() {
	this.signalsMtx.RLock()
	for _, cb := range this.signals {
		cb()
	}
	this.signalsMtx.RUnlock()
}

func start_workload_simulator(srcFile string, c chan<- struct{}) {
	// Simulate fopen / zstd(cgo) workload
	go func() {
		// Open a file
		fh, err := os.Open(srcFile)
		if err != nil {
			panic(err)
		}
		defer fh.Close()

		for {
			tmp := make([]byte, 4096)
			n, err := fh.ReadAt(tmp, 0)
			if err != nil {
				panic(err)
			}
			if n != 4096 {
				panic("incomplete read")
			}

			_, err = zstd.Compress(nil, tmp)
			if err != nil {
				panic(err)
			}

			c <- struct{}{}
		}
	}()
}

func main() {
	var srcFile string
	flag.StringVar(&srcFile, "src", "", "the source file")
	flag.Parse()

	t := NewTerminator("winiopanic")
	t.RegisterSignalHandler(func() {
		fmt.Println("\nserver: recieved signal, shutting down")
		os.Exit(1)
	})

	time.Sleep(500 * time.Millisecond)

	// Simulate workload
	work_recv := make(chan struct{}, 1024)
	start_workload_simulator(srcFile, work_recv)

	i := 0
	for _ = range work_recv {
		i++
		if i%100 == 0 {
			fmt.Printf("*")
		}

		// Start client thread after a little while
		if i == 3000 {
			fmt.Println("\nclient: making connection...")
			client, err := NewClient("winiopanic")
			if err != nil {
				panic(err)
			}
			fmt.Println("\nclient: connected")
			defer client.Close()
		}
	}
}
