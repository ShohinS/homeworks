package main

import "fmt"

func main() {
	PrintGreeting()
}

func PrintGreeting() {

	var dayTupe string
	fmt.Scan(&dayTupe)
	switch dayTupe {
	case "утро":
		fmt.Println("Доброе Утро")
	case "день":
		fmt.Println("Доброе День")
	case "вечер":
		fmt.Println("Добрый Вечер")

	}

}
