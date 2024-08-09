package main

import "fmt"

func main() {

	m1 := map[string]int{
		"apple":  1,
		"banana": 2,
	}
	//fmt.Println(PustoeMap(m1))
	InvertMap(m1)
	fmt.Println(m1)
	//UpMap(m1, "apple", 20)
	//fmt.Println(m1)
	//UdalitMap(m1, "apple")
	//fmt.Println(m1)
	//fmt.Println(ExsistKey(m1, "apple"))

	//Создание и вывод map
	//Описание: Создайте map для хранения строк и их длин, добавьте несколько элементов и выведите содержимое.

	//m := map[string]int{
	//	"Привет":    len("Привет"),
	//	"Шодмоншох": len("Шодмоншох"),
	//}
	//for key, value := range m {
	//	fmt.Println(key, value)
	//
	//}

}

//Проверка наличия ключа
//Описание: Создайте map с несколькими элементами и напишите функцию, которая проверяет наличие заданного ключа.

func ExsistKey(m map[string]int, key string) bool {
	_, exsist := m[key]
	return exsist
}

// Обновление значения по ключу
// Описание: Напишите функцию для обновления значения в map по заданному ключу.
func UpMap(m map[string]int, key string, newvelue int) {
	m[key] = newvelue

}

//Удаление элемента из map
//Описание: Создайте функцию, которая удаляет элемент из map по заданному ключу.

func UdalitMap(m map[string]int, key string) {
	delete(m, key)

}

//Проверка пустого map
//Описание: Напишите функцию, которая проверяет, пустой ли map.

func PustoeMap(m map[string]int) bool {
	return len(m) == 0

}

// Инвертирование ключей и значений
// Описание: Напишите функцию, которая инвертирует ключи и значения в map.
func InvertMap(m map[string]int) map[int]string {
	n := make(map[int]string)
	for k, v := range m {
		n[v] = k
	}

	return n
}
