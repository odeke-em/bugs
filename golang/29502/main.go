package main

func Bug() (e int64, err error) {
	var stack []int64
	stack = append(stack, 123)
	// Commenting the next 3 lines causes the function to work
	if len(stack) > 1 {
		panic("too many elements")
	}
	last := len(stack) - 1
	e = stack[last]
	stack = stack[:last]
	return e, nil
}

func main() {
	Bug()
}
