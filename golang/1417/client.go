package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"

	"google.golang.org/grpc"
)

func main() {
	addr := flag.String("addr", "", "The address of the server")
	flag.Parse()
	var before, after uint64
	var ms runtime.MemStats

	go func() {
		log.Fatal(http.ListenAndServe(":6060", nil))
	}()

	for {
		runtime.ReadMemStats(&ms)
		runtime.GC()
		before = ms.HeapAlloc
		perDial(*addr)
		runtime.GC()
		runtime.ReadMemStats(&ms)
		after = ms.HeapAlloc

		fmt.Printf("After: %d\nBefore: %d\nDifference: %d\n", after, before, after-before)
		<-time.After(7 * time.Second)
		runtime.GC()
	}
}

func perDial(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to dial to client: %v", err)
	}
        fmt.Printf("conn: %T\n", conn)
	if err := conn.Close(); err != nil {
		log.Fatalf("Failed to close: %v", err)
	}
}
