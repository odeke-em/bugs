package b

func GetY() interface{} {
	return struct{ x, y int }{2, 3}
}