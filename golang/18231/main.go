package main

type T struct {
	f map[string]string
}

func main() {
	_ = T{
		f: {



			"a": "b",
		},
	}
}
