package main

import "errors"

type error int

func F() error {
    return errors.New("x")
}

func main() {}
