package main

import "fmt"

func main() {
	fmt.Println(GetDiscount())

}
func GetDiscount() string {

	var amount int
	fmt.Scan(&amount)
	if amount > 1000 {
		if amount > 1000 {
			return "Скидки 10%"

		} else if amount > 500 && amount <= 1000 {
			return "Скидка 5%"
		} else {
			return "Скидка 0%"
		}

	}

}
