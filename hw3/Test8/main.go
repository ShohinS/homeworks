package main

import "fmt"

func main() {
	ToggleLight(true)
}
func ToggleLight(test bool) {
	if test == true {
		fmt.Println("Light")
	} else {
		fmt.Println("Lightoff")

	}
}
