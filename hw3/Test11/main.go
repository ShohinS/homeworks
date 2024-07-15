package main

import "fmt"

func main() {
	fmt.Println(IsEven(3))
	fmt.Println(IsEven(4))

}
func IsEven(x int) bool {
	if x%2 == 0 {
		return true

	} else {
		return false
	}

}
