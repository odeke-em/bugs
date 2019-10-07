package main

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

type ts struct {
	Timestamp time.Time
}

func TestCrash(t *testing.T) {
	want := []*ts{{}}

	got1 := &ts{}
	if diff := cmp.Diff(got1, want[0], roundToMicros); diff != "" {
		t.Fatalf("mismatch got - want +\n%s", diff)
	}
}

// Round times to the microsecond for comparison purposes.
var roundToMicros = cmp.Transformer("RoundToMicros",
	func(t time.Time) time.Time { return t.Round(time.Microsecond) })
