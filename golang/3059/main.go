package main

func main() {
	i := 0
	go func() {
		i = 1
	}()
}
