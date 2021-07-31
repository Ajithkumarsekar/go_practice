package main

import "fmt"

/*
https://go.dev/tour/methods/15
*/
func main() {
	typeAssertion("Hello")
	printType("Hello")
	printType(3)
	printType(true)
}

func printType(s interface{}) {
	fmt.Printf("Checking `%v` type\n", s)
	switch s.(type) {
	case string:
		fmt.Println("It's a string")
	case int:
		fmt.Println("It's a int")
	default:
		fmt.Println("Unknown dataType")
	}
}

func typeAssertion(i interface{}) {
	defer func() {
		if recover() != nil {
			fmt.Println("Gracefully handling the panic!")
		}
	}()

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // will panic if ok is not accepted
	fmt.Println(f)
}
