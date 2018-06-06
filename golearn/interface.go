package golearn

import (
	"fmt"
	"unsafe"
)

type Location int32

type Driver interface {
	Drive(from Location, to Location)
}

type Human struct {
	b byte
}

func (h *Human) Drive(from Location, to Location) {

}

type Robot struct {
	i int
}

func (*Robot) Drive(from Location, to Location) {

}

func GoInterfaceExample() {
	h := Human{11}
	//r := Robot{11}
	//i := Driver(nil)

	hi := Driver(&h)
	//ri := Driver(&r)
	//pi := new(Driver)

	// Memory size
	//fmt.Printf("struct [Human] size:%d-type:%v\n", unsafe.Sizeof(h), reflect.TypeOf(h))
	//fmt.Printf("struct [Robot] size:%d-type:%v\n", unsafe.Sizeof(r), reflect.TypeOf(r))
	//fmt.Printf("interface [Driver] size:%d-type:%v\n", unsafe.Sizeof(i), reflect.TypeOf(i))
	//
	//fmt.Printf("interface [Human] size:%d-type:%v\n", unsafe.Sizeof(hi), reflect.TypeOf(hi))
	//fmt.Printf("interface [Robot] size:%d-type:%v\n", unsafe.Sizeof(ri), reflect.TypeOf(ri))
	//fmt.Printf("interface pointer [Driver] size:%d-type:%v\n", unsafe.Sizeof(pi), reflect.TypeOf(pi))

	/*  内存布局简易示意图
			phi ---> ------------
					 | tab_ptr  |----------------------> |--------------------|
				     |----------|                        | interface_info_ptr |-------------------------->|------|
					 | data_ptr |-----                   |--------------------|                           | tags | --->|-----|
			         |----------|    |                   | type_info_ptr      |----------------           |------|     |     |
			          (hi : Driver)  |                   |--------------------|               |                        |*****|
									 |                   | func_ptr ...       |               |                        |     |
			                         |                   | *******************|               |                        |*****|
			            ph ---> ----------                                                    |                        |     |
			                    |        |                                                    |                        |*****|
			                    |--------|                                                    |
			                       (h : Human)                                                |
		                                                                             |------------|
		           |------|<---------------------------------------------------------| member_ptr |
		           | tag  |    (array)                                               |------------|
			       |------|
	               | addr |
		           |******|
		           | tag  |
			       |******|
	               | addr |
		           |******|
	*/
	ph := unsafe.Pointer(&h)
	phi := unsafe.Pointer(&hi)
	pdata := unsafe.Pointer(uintptr(phi) + 8)

	fmt.Printf("[sh] address:0x%x--->", ph)
	fmt.Printf(":%d\n", *(*int)(ph))
	fmt.Println("---------")
	fmt.Printf("[ih] address:0x%x--->", phi)
	fmt.Printf("0x%x\n", *(*int)(phi))
	fmt.Printf("[pd] address:0x%x--->", pdata)
	fmt.Printf("0x%x\n", *(*int)(pdata))
}
