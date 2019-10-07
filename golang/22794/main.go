package main

type it struct {
	ID string
}

func main() {
	i1 := &it{"Foo"}
	if i1.Id != "" {
	}
	i2 := &it{id: "Bar"}
}
