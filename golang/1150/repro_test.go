package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"testing"

	"cloud.google.com/go/logging"
        structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/api/option"
	logpb "google.golang.org/genproto/googleapis/logging/v2"
	"google.golang.org/grpc"
)

func TestIssue1150(t *testing.T) {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("Failed to find an available port: %v", err)
	}
	defer ln.Close()
	srv := grpc.NewServer()

	var logsSvc logsService
	logpb.RegisterLoggingServiceV2Server(srv, &logsSvc)

	go srv.Serve(ln)

	conn, err := grpc.Dial(ln.Addr().String(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to dial to the gRPC server: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	lc, err := logging.NewClient(ctx, "odeke-sandbox", option.WithGRPCConn(conn))
	if err != nil {
		t.Fatalf("Failed to create new client: %v", err)
	}
	defer lc.Close()

	lgr := lc.Logger("1150")

	key := "max"
	value := uint64(math.MaxUint64)
	if err := lgr.LogSync(ctx, logging.Entry{Payload: map[string]interface{}{
		key: value,
                key: uint64(18446744073709551615),
	}}); err != nil {
		t.Fatalf("Failed to logSync: %v\n", err)
	}

	if len(logsSvc) == 0 {
		t.Fatal("Expected at least one log")
	}
	headEntry := logsSvc[0]
	if headEntry == nil || len(headEntry.Entries) == 0 {
		t.Fatal("Could not retrieve a single entry")
	}
	payload := headEntry.Entries[0].GetPayload()
	jsonPayload, ok := payload.(*logpb.LogEntry_JsonPayload)
	if !ok {
		t.Fatalf("Got %T, expected a *logpg.LogEntry_JsonPayloaa", payload)
	}

	st := jsonPayload.JsonPayload
	if st == nil {
		t.Fatal("Unexpectedly got a nil JSON payload")
	}
	gotValue := st.Fields[key]
        var asString string

        switch gv := gotValue.Kind.(type) {
        case *structpb.Value_NumberValue:
            asString = fmt.Sprintf("%f", gv.NumberValue)
        case *structpb.Value_StringValue:
            asString = gv.StringValue
        }

	wantString := fmt.Sprintf("%d", value)
	if asString != wantString {
		t.Fatalf("Value as string mismatch\nGot:  %q\nWant: %q", asString, wantString)
	}
}

type logsService []*logpb.WriteLogEntriesRequest

var _ logpb.LoggingServiceV2Server = (*logsService)(nil)

func (lgs *logsService) DeleteLog(ctx context.Context, req *logpb.DeleteLogRequest) (*empty.Empty, error) {
	return nil, errors.New("unimplemented")
}

func (lgs *logsService) WriteLogEntries(ctx context.Context, req *logpb.WriteLogEntriesRequest) (*logpb.WriteLogEntriesResponse, error) {
	*lgs = append(*lgs, req)
	blob, _ := json.MarshalIndent(req, "", "  ")
	log.Printf("Got request:\n%s\n", blob)
	return new(logpb.WriteLogEntriesResponse), nil
}

func (lgs *logsService) ListLogEntries(ctx context.Context, req *logpb.ListLogEntriesRequest) (*logpb.ListLogEntriesResponse, error) {
	return nil, errors.New("unimplemented")
}

func (lgs *logsService) ListLogs(ctx context.Context, req *logpb.ListLogsRequest) (*logpb.ListLogsResponse, error) {
	return nil, errors.New("unimplemented")
}

func (lgs *logsService) ListMonitoredResourceDescriptors(ctx context.Context, req *logpb.ListMonitoredResourceDescriptorsRequest) (*logpb.ListMonitoredResourceDescriptorsResponse, error) {
	return nil, errors.New("unimplemented")
}
