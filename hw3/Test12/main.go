package main

import "fmt"

func main() {
	adder := add(1, 6)
	fmt.Println(adder)
}
func add(a, b int) int {
	return a + b

}
