package main

import "fmt"

func main() {
	GreetUser("Behruz")

}
func GreetUser(Users string) {
	fmt.Println("Hello" + " " + Users)
}
