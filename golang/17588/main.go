package main

type Foo func(b Bar)

func Bar(fn Foo) {
	func() {
		fn(nil)
	}()
}
