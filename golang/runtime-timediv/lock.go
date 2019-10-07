package main

func main() {
	v := int64(12345*1000000000+54321)
	div := int32(1000000000)
	_ = withPlus(v, div)
}
