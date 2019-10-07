package main

import (
    "reflect"
    "sync"
)

type R struct {
	Data interface{}
}

type S struct{Data int}

func main() {
	f()

	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1e5; j++ {
				f()
			}
		}()
	}
	wg.Wait()
}

var tx = &S{Data: 10}
func (p S) Run() R {
	return R{Data: tx.Data}
}

func f() interface{} {
    return reflect.ValueOf(&S{Data: 10}).MethodByName("Run").Interface().(func() R)()
}
