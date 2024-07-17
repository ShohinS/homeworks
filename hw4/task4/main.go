package main

import "fmt"

func main() {
	fmt.Println(GetGrade())

}
func GetGrade() string {

	var Score int
	fmt.Scan(&Score)
	if Score > 85 && Score <= 100 {
		return "A"
	} else if Score <= 85 && Score > 75 {
		return "B"

	} else {
		return "C"
	}
}
