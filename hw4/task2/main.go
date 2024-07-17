package main

import "fmt"

func main() {
	PrintWeather()
}
func PrintWeather() {
	var weatherTupe string
	fmt.Scan(&weatherTupe)
	switch weatherTupe {
	case "солнечно":
		fmt.Println("Солнечно")
	case "облачно":
		fmt.Println("Облачно")
	case "дождливо":
		fmt.Println("Дождливо")

	}

}
