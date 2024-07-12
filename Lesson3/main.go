package main

import "fmt"

func main() {
	a := 6
	fmt.Println(inc(a))
	fmt.Println(sum(5, 6))
}
func sum(a, b int) int {
	c := a + b
	return c
}

func inc(a int) {
	a++
	fmt.Println(a)
}
