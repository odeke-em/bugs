package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/rpc"
)

const SERVER = "127.0.0.1:9889"

var (
	Debug = flag.Bool("d", false, "set the debug mode")
	RPMD  = flag.Bool("rpmd", false, "rpmd-160 digest")
	SHA1  = flag.Bool("sha1", false, "sha1 digest")
)

var (
	serverKey = "server.key"
	clientKey = "client.key"
)

var caCert, serverCert, clientCert string

func init() {
	flag.Parse()
	if *RPMD {
		caCert = "ca-rmd160.pem"
		serverCert = "server-rmd160.pem"
		clientCert = "client1-rmd160.pem"
	} else if *SHA1 {
		caCert = "ca-sha1.pem"
		serverCert = "server-sha1.pem"
		clientCert = "client1-sha1.pem"
	} else {
		log.Fatal("choose flag rpmd or sha1")
	}
}

func server(certPool *x509.CertPool) (net.Listener) {
	//cert, err := tls.X509KeyPair(certBlock, keyBlock)
	cert, err := tls.LoadX509KeyPair(serverCert, serverKey)
	if err != nil {
		log.Fatal("server: load keys: ", err)
	}

	ServerTLSConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs: certPool,
		//ClientAuth:   tls. /*RequireAndVerifyClientCert,*/ VerifyClientCertIfGiven,
	}

	listener, err := tls.Listen("tcp", SERVER, ServerTLSConfig)
	if err != nil {
		log.Fatalf("server: listener: %s", err)
	}
	return listener
}

func client(certPool *x509.CertPool) *tls.Conn {
	//cert, err := tls.X509KeyPair(certBlock, keyBlock)
	cert, err := tls.LoadX509KeyPair(clientCert, clientKey)
	if err != nil {
		log.Fatal("client: load keys: ", err)
	}

	ClientTLSConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	}

	conn, err := tls.Dial("tcp", SERVER, ClientTLSConfig)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	return conn
}

// == RPC

type Args struct {
	Name string
}

type Echo int

func (Echo) Say(args *Args, reply *bool) error {
	fmt.Println(args.Name)
	*reply = true

	return nil
}

// * * *

func main() {
	log.SetFlags(0)

	CACertBlock, err := ioutil.ReadFile(caCert)
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(CACertBlock); !ok {
		log.Fatal("CA certificate not valid")
	}

	// Server

	listener := server(certPool)
	defer listener.Close()

	var echo Echo
	rpc.Register(echo)

	go func() {
		for {
			// Wait for clients
			conn_, err := listener.Accept()
			if err != nil {
				log.Printf("server: accept: %s", err)
				continue
			}

			if *Debug {
				log.Printf("server: accepted from %s", conn_.RemoteAddr())
			}
			go handleClient(conn_)

			break
		}
	}()

	// Client

	conn := client(certPool)
	defer conn.Close()

	client := rpc.NewClient(conn)
	defer client.Close()

	args := &Args{"November Rain"}
	var reply bool

	err = client.Call("Echo.Say", args, &reply)
	if err != nil {
		client.Close()
		log.Fatal(err)
	}
	fmt.Println("result:", reply)
}

func handleClient(c net.Conn) {
	if *Debug {
		log.Print("server: conn: client connected")
	}
	rpc.ServeConn(c)
	c.Close()

	if *Debug {
		log.Print("server: conn: client closed")
	}
}
