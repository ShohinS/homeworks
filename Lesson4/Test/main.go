package main

import "fmt"
func main(){
	var a int
	fmt.Scanf(&a)
	fmt.Println(a)
	switch a {
	case 1:
		fmt.Println("Январь 31 дней")
	case 2:
		fmt.Println("Февраль 28 дней")
	case 3:
		fmt.Println("Март 31 дней")

	case 4:
		fmt.Println("Апрел 31 дней")
	case 5:
		fmt.Println("май 30 дней")
	case 6:
		fmt.Println("Июнь 30 дней")
	case 7:
		fmt.Println("Июль 31 дней")
	case 8:
		fmt.Println("Август 31 дней")
	case 9:
		fmt.Println("сентябрь 30 дней")
	case 10:
		fmt.Println("Октябрь 31 дней")
	case 11:
		fmt.Println("Ноябрь 30 дней")
	case 12:
		fmt.Println("Декабрь 31 дней")

	}

}
