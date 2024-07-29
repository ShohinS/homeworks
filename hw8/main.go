package main

import "fmt"

func main() {
	//Max()
	//Min()
	//KolPol()
	//SumElv()
	//SrdZnach()
	//UdalitVho()
	//Moltiplay()
	// Index()
	//Obedinit()
	//Kopia()
	//PominatMaxMin()
}

// Максимальное элемент в массиве
func Max() {
	//Максимальное элемент в массиве
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Println(max)

}

// Минимальное элемент в массиве
func Min() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	fmt.Println(min)

}

// Подсчитать количество положительных чисел в массиве
func KolPol() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	kol := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			kol++
		}
	}

	fmt.Println(kol)

}

// Найти сумму всех элементов
func SumElv() {
	var arr [5]float64
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	s := float64(0)
	for i := 0; i < len(arr); i++ {
		s = s + arr[i]
	}

	fmt.Println(s)

}

// Найти среднее значение всех элементов
func SrdZnach() {
	var arr [5]float64
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	s := float64(0)
	for i := 0; i < len(arr); i++ {
		s = (s + arr[i]) / float64(len(arr))
	}

	fmt.Println(s)

}

// Удалить все вхождения заданого числа из массива
func UdalitVho() {
	var arr [5]float64
	var n float64
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&n)
	var lan []float64
	for _, v := range arr {
		if v != n {
			lan = append(lan, v)
		}

	}
	fmt.Println(lan)

}

// Умножить все элементы массива на заданное число
func Moltiplay() {
	var arr [5]float64
	var n float64
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&n)
	var lan []float64
	for _, v := range arr {
		a := n * v
		lan = append(lan, a)

	}
	fmt.Println(lan)

}

// Найти все индексы заданного числа в массиве
func Index() {
	var arr [5]float64
	var n float64
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Scan(&n)
	var lan []float64
	for i, v := range arr {
		if v == n {
			lan = append(lan, float64(i))
		}

	}
	fmt.Println(lan)

}

// Создать копию массива
func Kopia() {

	var arr [5]float64
	//var n float64
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i])
	}
	//fmt.Scan(&n)
	var lan []float64
	for _, v := range arr {
		lan = append(lan, v)

	}
	fmt.Println(lan)

}

// объединить два массива
func Obedinit() {

}

//Поменять местами максимальное и минимальное

func PominatMaxMin() {

}
