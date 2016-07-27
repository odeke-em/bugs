package main

func foo() (int, int) {return 2.3} // ERROR ""
func foo2() {return int(2), 2} // ERROR ""
func foo3() (a, b, c, d int) {return 1} // ERROR ""

type node struct {
  left, right *node
  data interface{}
  visitCount uint64
}

func (n *node) walk() {
        if n == nil {
                return 0
        }
        n.left.walk()
        n.right.walk()
        n.visitCount += 1
        return n.visitCount
}

func main() {
}
