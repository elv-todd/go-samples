package main

import (
	"bytes"
	"fmt"
)

type path []byte

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func (p path) ToUpper() {
	for i, b := range p {
		if 'a' <= b && b <= 'z' {
			p[i] = b + 'A' - 'a'
		}
	}
}

func Extend(slice []int, element int) []int {
	n := len(slice)
	if n == cap(slice) { // Slice is full; must grow.
		newSlice := make([]int, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

// Append : actually, use append(slice []T, ...int) builtin
func Append(slice []int, items ...int) []int {
	for _, item := range items {
		slice = Extend(slice, item)
	}
	return slice
}

// see also https://github.com/golang/go/wiki/SliceTricks

func main() {
	pathName := path("/usr/bin/tso") // Conversion from string to path.
	pathName.TruncateAtFinalSlash()
	fmt.Printf("Trunc: %s\n", pathName)

	pathName.ToUpper()
	fmt.Printf("ToUpper: %s\n", pathName)

	slice := make([]int, 0, 5)
	for i := 0; i < 10; i++ {
		slice = Extend(slice, i)
		fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
		fmt.Println("updated address of 0th element:", &slice[0])
	}

	slice = []int{0, 1, 2, 3, 4}
	fmt.Println(slice)
	slice = Append(slice, 5, 6, 7, 8)
	slice2 := []int{55, 66, 77}
	slice = Append(slice, slice2...)
	fmt.Println(slice)

}
