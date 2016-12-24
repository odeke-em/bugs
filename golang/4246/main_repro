package main

func f1(m map[string]string) {
	m["ok"] = "ok"
}

func f2(m map[string]string) {
	m["that part"] = "that part"
}

func main() {
	m := make(map[string]string)
	go f1(m)
	go f2(m)
}
