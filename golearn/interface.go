package golearn

import (
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	id   int
	name string
}

type iface struct {
	itab uintptr
	data uintptr
}

type Stringer interface {
	String() string
}

type Printer interface {
	Stringer
	Print()
}

func (self *User) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}

func (self *User) Print() {
	fmt.Println(self.String())
}

func GoInterfaceExample() {
	////////////////////////////////////////////////////
	var o Printer = &User{1, "Tom"}
	var s Stringer = o
	fmt.Println(s.String())
	//////////////////////////////////////////////////
	//var o interface{} = &User{1, "Tom"}
	//
	//if i, ok := o.(fmt.Stringer); ok {
	//	fmt.Println(i)
	//}
	//
	////u := o.(User) error
	//u := o.(*User)
	//fmt.Println(u)
	//
	//switch v := o.(type) {
	//case nil:
	//	fmt.Println("nil")
	////case fmt.Stringer:
	////	fmt.Println("Stringer", v)
	//case func() string:
	//	fmt.Println("func() string", v)
	//case *User:
	//	fmt.Printf("*User %d, %s\n", v.id, v.name)
	//default:
	//	fmt.Println("unknow")
	//}
	///////////////////////////////////////
	//var a interface{} = nil
	//var b interface{} = (*int)(nil)
	//
	//ia := *(*iface)(unsafe.Pointer(&a))
	//ib := *(*iface)(unsafe.Pointer(&b))
	//
	//fmt.Println(a == nil, ia)
	//fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil())
	//////////////////////////////////////
	//u := User{1, "Tom"}
	//var vi, pi interface{} = u, &u
	//
	////vi.(User).name = "Jack"
	//pi.(*User).name = "Ma"
	//
	//fmt.Printf("%v\n", vi.(User))
	//fmt.Printf("%v\n", pi.(*User))
	/////////////////////////////////////
	//u := User{1, "Tom"}
	//var i interface{} = u
	//
	//u.id = 2
	//u.name = "Jack"
	//
	//fmt.Printf("%v\n", u)
	//fmt.Printf("%v\n", i.(User))
	/////////////////////////////////////
	//t := Tester{&User{2, "Ada"}}
	//fmt.Println(t.s.String())
	/////////////////////////////////////
	//var t Printer = &User{1, "Tom"}
	//t.Print()
}

/*****************************************************************************/
//type Tester struct {
//	s interface {
//		String() string
//	}
//}

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

func maybeError() {
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
