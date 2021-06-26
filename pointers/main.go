package main

import (
	"fmt"
	"reflect"
)

/*
* is called dereferencing operator
& is called address operator
*/

func main() {
	example1()
	example2()
	example3()
	example4()
	example5() //important
}

func example1() {
	x := 7
	// y is a pointer to x
	y := &x
	fmt.Println(x, y)

	// * (asterisk) - deference operator
	// point to the value store in the value (address) stored in y
	*y = 10
	fmt.Println(x, y)

	panicOnNilPointerDereference()
}

func panicOnNilPointerDereference() {

	defer func() {
		if recover() != nil {
			fmt.Println("Gracefully handling the panic!")
		}
	}()

	// trying to dereference a nil address would panic with
	// `invalid memory address or nil pointer dereference`
	var p *int
	p = nil
	fmt.Println(*p)
}

func changeValue(str *string) {
	// str is a pointer which means it contains a memmory address as it's value
	// so we are de-referencing it with * and modifying the actual value in the
	// memory location
	*str = fmt.Sprintf("string value - %s, is changed by changeValue function", *str)
}

func example2() {
	name := "Ajith kumar"
	fmt.Println(name)
	changeValue(&name)
	fmt.Println(name)
}

func example3() {
	x := 7
	var y *int = &x // y := &x

	// int & *int are to different data types
	fmt.Printf("Type of x : %v, Type of y : %v\n", reflect.TypeOf(x), reflect.TypeOf(y))
	// memory address of x and value of y should be same
	fmt.Printf("Memory address of x : %v, Memory address of y : %v, value of variable y : %v\n", &x, &y, y)

}

type Profile struct {
	name string
	age  int
}

func changeProfile(profile Profile) {
	// go always pass by value so the name value won't change
	profile.name = "Moni"
}

func changeProfile1(profile *Profile) {
	profile.name = "Moni"
}

func example4() {
	profile := Profile{
		name: "Ajithkumar",
		age:  25,
	}
	changeProfile(profile)
	fmt.Println(profile)
	changeProfile1(&profile)
	fmt.Println(profile)
}

type doError struct {
}

func (d doError) Error() string {
	return "error"
}

func do() *doError {
	return nil
}

// error is an interface so when we assign a *doError which is nil to an interface (error),
// it won't become nil because the interface will store the dataType, and it's value. Here the
// dataType will be doError and the value will be nil
func wrapDo() error {
	return do()
}

// do not return concrete error types
func example5() {
	err := wrapDo()
	// because error interface store the dataType as *doError but its value is nil
	fmt.Println(err == nil) // false
	fmt.Println(err)        //<nil>

	err1 := do()
	fmt.Println(err1 == nil) // true
	fmt.Println(err1)        //<nil>
}
