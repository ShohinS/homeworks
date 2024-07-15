package main

import "fmt"

func main() {
	concatenator := Conc("holov", "Mukim")
	fmt.Println(concatenator)
}
func Conc(a, b string) string {
	return a + b
}
