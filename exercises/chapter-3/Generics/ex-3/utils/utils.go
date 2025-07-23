package utils

import (
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](a,b T) T {
	return max(a,b)
}

func Sum[T constraints.Float](a,b T) T {
	return a+b
}
