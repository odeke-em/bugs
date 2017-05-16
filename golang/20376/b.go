package main

type plugin string

func init() {
	Rule = "b"
}

func (p plugin) String() string {
	return "b"
}

var Rule plugin
