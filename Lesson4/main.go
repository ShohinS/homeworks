package main

import "fmt"

func main() {

	num := 15
	if num > 9 && num < 100 {
		fmt.Println("Двузначное")
	}
	n := 8
	if n%2 == 1 {
		fmt.Println("нечетное")
	} else {
		fmt.Println("Четное")
	}
	m := 78
	if m >= 0 && m < 10 {
		fmt.Println("Однозначное")
	} else if m >= 10 && m < 100 {
		fmt.Println("Двузначное")

	}

	a := -8
	b := -7
	if a > b {
		fmt.Println("Первое число")
	} else {
		fmt.Println("Второе число")
	}

	a:=2
	b:=3
	c:=5
	if a<b && b<c {
		fmt.Println(c+b)
	}else if b<a && a<c {
		fmt.Println(c+a)
	} else if b<c && c<a {
		fmt.Println(a+c)
	} else if c<b && b<a {
		fmt.Println(a + b)
	} else if c<a && a<b {
		fmt.Println(a + b)
	} else if a<b && b<c {
		fmt.Println(c + b)
	} else if a<c && c<b {
		fmt.Println(c+b)

	}
		{

	}
	}




	}

	}

	nn:=1
	switch nn{
	case 1:
		fmt.Println("Positive")

	case 2:
		fmt.Println("Negative")
		case 0:
			fmt.Println("Zero")

}

	var n int
	n=1200
	if n %100 ==0 && n%400!=0 {
fmt.Println("No")
}  else if n%400==0 {
	fmt.Println("Yes")
}

