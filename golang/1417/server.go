package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"google.golang.org/grpc"
)

func main() {
	ln, err := net.Listen("tcp", ":64574")
	if err != nil {
		log.Fatalf("Failed to find an available address: %v", err)
	}
	defer ln.Close()

	log.Printf("Server address: %s\n", ln.Addr())

	srv := grpc.NewServer()
	dialogflow.RegisterSessionsServer(srv, new(sessionsServer))
	log.Fatal(srv.Serve(ln))
}

type sessionsServer int

func (ss *sessionsServer) DetectIntent(ctx context.Context, req *dialogflow.DetectIntentRequest) (*dialogflow.DetectIntentResponse, error) {
	return new(dialogflow.DetectIntentResponse), nil
}

func (ss *sessionsServer) StreamingDetectIntent(stream dialogflow.Sessions_StreamingDetectIntentServer) error {
	for {
		recv, err := stream.Recv()
		if err != nil {
			return err
		}
		_ = recv
		if err := stream.Send(new(dialogflow.StreamingDetectIntentResponse)); err != nil {
			return err
		}
	}
	return nil
}
