package main

func main() {
	var name = "gopher"
	f(&name)

	var nameCopy string
	g(&name, &nameCopy)

	var a, b int
	a = 10
	h(&a, &b)
	i(&b)
}

func i(a *int) {
	*a = *a
}

func h(i, j *int) {
	*j = *i
}

func g(s, t *string) {
	*t = *s
}

func f(s *string) {
	*s = *s
}
