package golearn

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Location int32

type Driver interface {
	Drive(from Location, to Location)
}

type Human struct {
}

func (h *Human) Drive(from Location, to Location) {

}

type Robot struct {
	i int
}

func (*Robot) Drive(from Location, to Location) {

}

func GoInterfaceExample() {
	h := Human{}
	r := Robot{}
	i := Driver(nil)

	hi := Driver(&h)
	ri := Driver(&r)
	pi := new(Driver)

	// Memory size
	fmt.Printf("struct [Human] size:%d-type:%v\n", unsafe.Sizeof(h), reflect.TypeOf(h))
	fmt.Printf("struct [Robot] size:%d-type:%v\n", unsafe.Sizeof(r), reflect.TypeOf(r))
	fmt.Printf("interface [Driver] size:%d-type:%v\n", unsafe.Sizeof(i), reflect.TypeOf(i))

	fmt.Printf("interface [Human] size:%d-type:%v\n", unsafe.Sizeof(hi), reflect.TypeOf(hi))
	fmt.Printf("interface [Robot] size:%d-type:%v\n", unsafe.Sizeof(ri), reflect.TypeOf(ri))
	fmt.Printf("interface pointer [Driver] size:%d-type:%v\n", unsafe.Sizeof(pi), reflect.TypeOf(pi))
}
