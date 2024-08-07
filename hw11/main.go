package main

import (
	"fmt"
	"strings"
)

func main() {
	kon := Konkantenasiy("Hello", "Shodmon")
	fmt.Println(kon)
	Dlina("Shodmon")
	fmt.Println(Dlina)
}

// Напишите функцию, которая принимает две строки и возвращает их конкатенацию.
func Konkantenasiy(str1, str2 string) string {
	return str1 + str2
}

// Напишите функцию, которая принимает строку и возвращает её длину.
func Dlina(str string) int {
	return len(str)
}

// Напишите функцию, которая проверяет, содержит ли строка заданную подстроку.
func PodStrok(str, now string) bool {
	return strings.Contains(str, now)
}

//Напишите функцию, которая находит индекс первого вхождения подстроки в строке.
//	Верните -1, если подстрока не найдена.

func PodstrokPoindex(str, now string) bool {
	strings.Index(str, now)
}

// Напишите функцию, которая заменяет все вхождения подстроки в строке на другую подстроку.
func DrugiePodstroki(str, now, new string) string {
	return strings.ReplaceAll(str, now, new)
}

//Напишите функцию, которая удаляет пробелы в начале и в конце строки.

func UdalitProbelov(str, now string) string {
	return strings.TrimSpace(str)
}

// Напишите функцию, которая преобразует строку к верхнему и нижнему регистру.
func PreobStrokiVer(str, now string) string {
	strings.ToLower(str)
	strings.ToUpper(str)
}

//Напишите функцию, которая повторяет строку заданное количество раз.

func StrokaKolichestvo(str string, x int) string {
	strings.Repeat(str, x)
}

// Напишите функцию, которая разбивает строку по заданному разделителю и возвращает срез строк.
func StrokiRozdelite(str, now string) []string {
	return strings.Split(str, now)
}
