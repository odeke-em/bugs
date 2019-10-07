package main

import (
        "runtime"
        . "testing"

        radix "github.com/mediocregopher/radix.v3"
)                                                      
                                                       
type nilAction struct{}                                
                                       
func (nilAction) Keys() []string {            
        return nil        
}                         
                                                
func (nilAction) Run(radix.Conn) error {
        return nil                              
}                                                                                                                 
                                                                                                                  
func BenchmarkPanic(b *B) {                                                                                       
        parallel := runtime.GOMAXPROCS(0)                                                                         
                                                                                                                                                  
        pool, err := radix.NewPool("tcp", "127.0.0.1:6379", parallel)
        if err != nil {                
                b.Fatal(err)                  
        }

        do := func(b *B, fn func()) {
                b.SetParallelism(parallel)
                b.RunParallel(func(pb *PB) {
                        for pb.Next() {
                                fn()
                        }
                })
        }

        b.Run("not-panicking", func(b *B) {
                do(b, func() {
                        pool.Do(nilAction{})
                })
        })

        b.Run("panicking", func(b *B) {
                do(b, func() {
                        pool.Do(radix.Cmd(nil, "SET", "foo", "bar"))
                })
        })
        
}