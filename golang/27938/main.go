package p

/*
type t struct {
        F sync.Mutex
}
*/

type s struct {
        sync.Mutex
}

var _ io.Writer
