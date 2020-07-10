package test

type Foo interface {
       Bar() *interface{Foo}
       Baz() *interface{Foo}
       Bug() *interface{Foo}
}
