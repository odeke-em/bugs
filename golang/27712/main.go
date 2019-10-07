package main

func main() {
	var s uint
	var _ interface{} = 1.0 << s

        x := int(1)
        var _ interface{} = 2 << x
        var _ string = 1
}
