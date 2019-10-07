package main

type mockprocessor struct {
	ProcessFunc func(int, int) bool
}

func main() {
	proc := &mockprocessor{
		ProcessFunc: func(in, out int) {
		},
	}
}
