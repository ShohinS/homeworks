package main

import "fmt"

func main() {
	isPosite := Posite(-3)
	fmt.Println(isPosite)
}

func Posite(x int) bool {
	if x < 0 {
		return false
	} else {
		return true
	}
}
