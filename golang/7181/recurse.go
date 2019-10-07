package main 
 
import ( 
  "runtime/debug" 
) 
 
func recurse() { 
  debug.SetMaxStack(10000) 
  var f1, f2, f3, f4 func() 
  f1 = func() { f2() } 
  f2 = func() { f3() } 
  f3 = func() { f4() } 
  f4 = func() { f1() } 
 
  f1() 
} 
 
func main() { 
  recurse() 
}