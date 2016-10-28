package main

func main() {
	go func() {
		i := 0
		i = 1
	}()
}
