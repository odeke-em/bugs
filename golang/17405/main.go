package main

func main() {
	var s interface{} = ""
	if k := s.(error); k != nil {
	}
}
