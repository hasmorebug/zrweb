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
	b byte
}

func (h *Human) Drive(from Location, to Location) {
	fmt.Print("Human Driver")
}

type Robot struct {
	i int
}

func (*Robot) Drive(from Location, to Location) {

}

func GoInterfaceExample() {
	h := Human{11}
	//r := Robot{11}
	//	//i := Driver(nil)

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

	fmt.Println("=== struct ===")
	fmt.Printf("[human] address:0x%x--->", getAddrPtr(ph))
	fmt.Printf(":%d\n", getAddrValue(ph))
	faddr := reflect.ValueOf(h.Drive).Pointer()
	fmt.Printf("[ func] address:0x%x\n", faddr)
	fmt.Println("--- interface ---")
	fmt.Printf("[ tab] address:0x%x--->", phi)
	fmt.Printf("0x%x\n", *(*int)(phi))
	fmt.Printf("[data] address:0x%x--->", getAddrPtr(pdata))
	fmt.Printf("0x%x\n", getAddrValue(pdata))

	fmt.Println("*** tab ***")
	ptab := unsafe.Pointer(uintptr(getAddrValue(phi)))
	ptpinfo := unsafe.Pointer(uintptr(ptab) + 8)
	pfunc := unsafe.Pointer(uintptr(ptpinfo) + 8)
	fmt.Printf("[interface_type_info] address:0x%x--->", ptab)
	fmt.Printf("0x%x\n", getAddrValue(ptab))
	fmt.Printf("[          type_info] address:0x%x--->", getAddrPtr(ptpinfo))
	fmt.Printf("0x%x\n", getAddrValue(ptpinfo))
	fmt.Printf("[               func] address:0x%x--->", getAddrPtr(pfunc))
	fmt.Printf("0x%x\n", getAddrValue(pfunc))

	fmt.Println("&&& InterfaceTypeInfo &&&")
	piti := unsafe.Pointer(uintptr(getAddrValue(ptab)))
	fmt.Printf("[tag] address:0x%x--->", getAddrPtr(piti))
	fmt.Printf("0x%x\n", getAddrValue(piti))

	fmt.Println("@@@ TypeInfo @@@")
	pti := unsafe.Pointer(uintptr(getAddrValue(ptpinfo)))
	fmt.Printf("[tag] address:0x%x--->", getAddrPtr(pti))
	fmt.Printf("0x%x\n", getAddrValue(pti))
}

func getAddrPtr(p unsafe.Pointer) *int {
	return (*int)(p)
}

func getAddrValue(p unsafe.Pointer) int {
	return *(*int)(p)
}
