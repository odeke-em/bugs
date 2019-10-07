package main

func bar() {
	i := 0
	func() {
            i = 5
	}()
}

type Foo struct {
}
