package pathing

import (
	"testing"
)

func Name() string {
	return "pathing"
}

func TestName(t *testing.T) {
	//var name string
	if name = Name(); name != "Pathing" {
		t.Error("unexpected name")
	}
}
