package main

import (
	"unsafe"
	"fmt"
	"unicode/utf8"
)

func ReadMemory(ptr unsafe.Pointer, size uintptr) []byte {
	out := make([]byte, size)
	for i := range out {
		out[i] = *((*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(i))))
	}
	return out
}

type goString struct {
	elements []byte // underlying string bytes
	len      int   // number of bytes
}

// 文
func main() {
	s := []byte("Hello world")
	var stringExample = "Hello world"
	var anotherStringExample = "Hello world"
	var goString = goString{s, 11}

	sz := unsafe.Sizeof(stringExample)
	fmt.Println(sz) // 16

	fmt.Println(unsafe.Pointer(&stringExample))
	fmt.Println(unsafe.Pointer(&anotherStringExample))

	stringExample = anotherStringExample

	fmt.Println(unsafe.Pointer(&stringExample))
	fmt.Println(unsafe.Pointer(&anotherStringExample))

	n := unsafe.Pointer(&goString.elements[0])
	fmt.Println(ReadMemory(n, 11))
	fmt.Println(string(ReadMemory(n, 11)))

	fmt.Println(utf8.RuneLen('文')) //3
	buf := []byte{0,0,0}
	utf8.EncodeRune(buf, '文')

	r, _ := utf8.DecodeRune(buf)
	fmt.Printf("%q", r)
}