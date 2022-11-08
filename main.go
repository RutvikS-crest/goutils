package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Max(10,20))
	fmt.Println(Max(5.55,5.0))
}
