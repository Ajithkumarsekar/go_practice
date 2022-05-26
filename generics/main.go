package main

import (
	"fmt"

	"github.com/go_practice/generics/mymath"
)

func main() {
	fmt.Printf("mymath.SumInt64(23849, 987) = %v\n", mymath.SumInt64(23849, 987))
	fmt.Printf("mymath.Sum(23849, 987) = %v\n\n", mymath.Sum(23849, 987))

	fmt.Printf("mymath.SumInt32(432, 947) = %v\n", mymath.SumInt32(432, 947))
	fmt.Printf("mymath.Sum(432, 947) = %v\n\n", mymath.Sum(432, 947))

	fmt.Printf("mymath.SumFloat64(0.4089, 3.084212) = %v\n", mymath.SumFloat64(0.4089, 3.084212))
	fmt.Printf("mymath.Sum(0.4089, 3.084212) = %v\n\n", mymath.Sum(0.4089, 3.084212))

}
