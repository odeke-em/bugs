package b

var (
    w []*uint
    x = *w[0]
    y = x
    z = (uintptr)(y)
)

func main() {
}
