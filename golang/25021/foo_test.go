package main

import "testing"

func TestX(t *testing.T) {

	t.Run("A", func(t *testing.T) {
		t.Log("hi!")      // Never gets printed
		t.Fatal("hello!") // Never gets printed
	})

	t.Run("B", func(t *testing.T) {
		panic("goodbye!")
	})
}
