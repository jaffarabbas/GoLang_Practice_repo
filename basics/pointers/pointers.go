package pointers

import "fmt"

func simplePointer() {
	// Simple pointer
	i := 42
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
}

func pointerWithStruct() {
	// Pointer with struct
	type person struct {
		name string
		age  int
	}
	p1 := person{"Alice", 23}
	p2 := &p1
	p2.age = 24
	fmt.Println(p1)
}

func pointerWithNew() {
	// Pointer with new
	p := new(int)
	fmt.Println(*p)
	*p = 21
	fmt.Println(*p)
}

func Pointer() {
	simplePointer()
	pointerWithStruct()
	pointerWithNew()
}
