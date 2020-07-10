package main

import "testing"

func TestIt(t *testing.T) {
	go func() {
		t.Fatalf("Ending here")
	}()

        go func() {
		t.Skip("Ending here")
        }()
}
