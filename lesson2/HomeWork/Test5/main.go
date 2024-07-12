package main

import "fmt"

func main() {
	var a, b, c, s, v float32
	a = 10
	b = 11
	c = 12
	v = a * b * c
	s = 2 * (a*b + c*b + a*c)
	fmt.Println(v, s)

}
