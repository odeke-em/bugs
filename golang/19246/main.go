package main

import (
	"fmt"
	"log"
)

type A struct {
	Cs []*C
}

func (a *A) String() string {
	return fmt.Sprintf("%+v", *a)
}

type B struct{}

func (b *B) String() string {
	return fmt.Sprintf("B.STRING WAS CALLED %+v", *b)
}

type C struct {
	Name string
	B    *B
}

func (c *C) String() string {
	return fmt.Sprintf("%+v", *c)
}

func main() {
	a := A{
		Cs: []*C{{
			Name: "ae1",
		}},
	}
	want := "{Cs:[{Name:ae1 B:<nil>}]}"

	if g, w := a.String(), want; g != w {
		log.Fatalf("got %q want %q", g, w)
	}
}