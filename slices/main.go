package main

import (
	"fmt"
	"reflect"
)

/*
reference : https://go.dev/blog/slices-intro
*/

func main() {
	// an array is length and element type
	var a [5]int // all elements will be initialized to its default value which is 0 in this case
	fmt.Printf("a: %+v\n", a)

	// both c and c are equal. Just various ways of initializing array
	b := [2]string{"Penn", "Teller"}
	//here the compiler count the array elements for you
	c := [...]string{"Penn", "Teller"}
	if !reflect.DeepEqual(b, c) {
		panic("Both b, c should be equal")
	}

	// A slice literal is declared just like an array literal, except you leave out the element count
	_ = []string{"Penn", "Teller"}

	// The make function takes a type, a length, and an optional capacity.
	// When called, make allocates an array and returns a slice that refers to that array.
	s := make([]byte, 5, 5)
	// s == []byte{0, 0, 0, 0, 0}

	//A slice can also be formed by “slicing” an existing slice or array

	x := s[3:] //x slice referencing the storage of s
	fmt.Printf("x slice content : %+v, len : %v, cap : %v\n", x, len(x), cap(x))

	/*
		A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity
		type slice struct{
			pointerToArray // It can also point out to nth element in an existing array
			length
			maximumCapacity
		}
	*/

}
