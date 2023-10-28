package main

import (
	"fmt"
)

// Any accept any type
func cloneV[V any](v []V) []V {
	var c = make([]V, len(v))

	for k, v := range v {
		c[k] = v
	}

	return c
}

func cloneM[K comparable, V any](m map[K]V) map[K]V {
	var c = make(map[K]V)

	for k, v := range m {
		c[k] = v
	}

	return c
}

// this for describe types that can be added
type addable interface {
	int | float32 | float64 | string
}

func add[T addable](a []T) T {
	var result T

	for _, v := range a {
		result += v
	}

	return result
}

func main() {
	var v = []float32{1, 2, 3, 4, 5}
	var c = cloneV(v)

	println("v", v)
	println("c", c)

	var m = map[string]int{
		"a": 1,
		"b": 2,
	}

	var cm1 = cloneM(m)

	fmt.Print(cm1)

	var cm2 = cloneM(map[int]string{1: "a", 2: "b"})

	fmt.Print(cm2)

	fmt.Println(add([]int{1, 2, 3, 4, 5}))
	fmt.Println(add([]float32{1.1, 2.2, 3.3, 4.4, 5.5}))
	fmt.Println(add([]string{"a", "b", "c", "d", "e"}))
}
