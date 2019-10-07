package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"math"
	"net"

	"cloud.google.com/go/logging"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/api/option"
	logpb "google.golang.org/genproto/googleapis/logging/v2"
	"google.golang.org/grpc"
)

func main() {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("Failed to find an available port: %v", err)
	}
	defer ln.Close()
	srv := grpc.NewServer()
	logpb.RegisterLoggingServiceV2Server(srv, new(logServiceServer))

	go srv.Serve(ln)

	conn, err := grpc.Dial(ln.Addr().String(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to dial to the gRPC server: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	lc, err := logging.NewClient(ctx, "odeke-sandbox", option.WithGRPCConn(conn))
	if err != nil {
		log.Fatalf("Failed to create new client: %v", err)
	}
	defer lc.Close()

	lgr := lc.Logger("1150")

        max_uint64 := uint64(math.MaxUint64)
        log.Printf("On the client, max_uint64=%d aka %x\n", max_uint64, max_uint64)
	if err := lgr.LogSync(ctx, logging.Entry{Payload: map[string]interface{}{
		"max": max_uint64,
	}}); err != nil {
		log.Fatalf("Failed to logSync: %v\n", err)
	}
}

type logServiceServer int

var _ logpb.LoggingServiceV2Server = (*logServiceServer)(nil)

func (lgs *logServiceServer) DeleteLog(ctx context.Context, req *logpb.DeleteLogRequest) (*empty.Empty, error) {
	return nil, errors.New("unimplemented")
}

func (lgs *logServiceServer) WriteLogEntries(ctx context.Context, req *logpb.WriteLogEntriesRequest) (*logpb.WriteLogEntriesResponse, error) {
	blob, _ := json.MarshalIndent(req, "", "  ")
	log.Printf("Got request:\n%s\n", blob)
	return new(logpb.WriteLogEntriesResponse), nil
}

func (lgs *logServiceServer) ListLogEntries(ctx context.Context, req *logpb.ListLogEntriesRequest) (*logpb.ListLogEntriesResponse, error) {
	return nil, errors.New("unimplemented")
}

func (lgs *logServiceServer) ListLogs(ctx context.Context, req *logpb.ListLogsRequest) (*logpb.ListLogsResponse, error) {
	return nil, errors.New("unimplemented")
}

func (lgs *logServiceServer) ListMonitoredResourceDescriptors(ctx context.Context, req *logpb.ListMonitoredResourceDescriptorsRequest) (*logpb.ListMonitoredResourceDescriptorsResponse, error) {
	return nil, errors.New("unimplemented")
}
