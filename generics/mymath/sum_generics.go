package mymath

import (
	"github.com/go_practice/generics/mycontstraints"
)

func Sum[T mycontstraints.Number](x, y T) T {
	return x + y
}
