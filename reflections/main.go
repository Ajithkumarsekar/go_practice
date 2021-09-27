package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Foo struct {
	A int `tag1:"First Tag" tag2:"Second Tag"`
	B string
}

func main() {
	example2(Foo{
		A: 10,
		B: "myName",
	})
}

func example2(t interface{}) {
	rType := reflect.TypeOf(t)
	fmt.Printf("reflect type : %v\n", rType)
	rValue := reflect.ValueOf(t)
	fmt.Printf("reflect value : %v\n", rValue)
}

func example1() {
	sl := []int{1, 2, 3}
	greeting := "hello"
	greetingPtr := &greeting
	f := Foo{A: 10, B: "Salutations"}
	fp := &f

	gType := reflect.TypeOf(greeting)
	examiner(gType, 0)
	fmt.Println()

	grpType := reflect.TypeOf(greetingPtr)
	examiner(grpType, 0)
	fmt.Println()

	slType := reflect.TypeOf(sl)
	examiner(slType, 0)
	fmt.Println()

	fType := reflect.TypeOf(f)
	examiner(fType, 0)
	fmt.Println()

	fpType := reflect.TypeOf(fp)
	examiner(fpType, 0)
	fmt.Println()
}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}
