package main

func main() {
	i := 0
        m := make(map[string]string)
	go func() {
		i = 1
                m = nil
	}()
        println(m == nil)
}
