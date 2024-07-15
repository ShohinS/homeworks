package main

import "fmt"

func main() {
	fmt.Println(Concat("Horov", "Jamshed"))

}

func Concat(name, Surname string) string {
	return name + Surname
}
