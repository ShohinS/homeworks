package main

import "fmt"

func main() {

}
func PrintTrafficLight() {

	var lightColor string
	fmt.Scan(&lightColor)
	switch lightColor {
	case "красный":
		fmt.Println("Стоп")
	case "желтый":
		fmt.Println("Внимание")
	case "зелный":
		fmt.Println("Идите")

	}
}
