package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b uint
	a = 12
	b = 3
	s := math.Sqrt(float64(a * b))
	fmt.Println(s)
}
