package main

type plugin string

func init() {
	Rule = "a"
}

func (p plugin) String() string {
	return "a"
}

var Rule plugin
