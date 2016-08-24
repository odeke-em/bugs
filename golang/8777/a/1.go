// a/1.go
package a
func foo(i interface{}, c1 complex64, c2 complex128) int {
    if i == nil {
        panic("nil")
    }
    return 0
}
func bar(b bool, s string, c byte, i int) (interface{}, complex64, complex128) {
    if b {
        panic("nil")
    }
    return nil, 0, 0
}
