package main

import "fmt"

func main() {
	/*	var i int
					for i := 0; i <= 9; i++ {
						fmt.Println(i)
					}
					fmt.Println(i)

					for i := 1; i <= 100; i++ {
						if i%3 == 0 {
							break
						}
						fmt.Println(i)
					}

				var n int
				fmt.Scan(&n)
				s := 0.0
				for i := 1; i <= n; i++ {
					s = s + 1/float64(i)
				}
				fmt.Println(s)


			var n int
			fmt.Scan(&n)
			s := 1.0
			for i := 1; i <= n; i++ {
				s = s * (1 + float64(i)/10)

			}
			fmt.Println(s)


		s := 1.0
		var a float64
		for i := 1; i <= 10; i++ {

			fmt.Scan(&a)
			s = s * a
		}
		fmt.Println(s)
	*/
	var n int
	fmt.Scan(&n)
	var a float64
	var b float64
	for i := 1; i <= n; i++ {

		fmt.Scan(&a)
		b = a - float64(int(a))
		fmt.Println(b)

	}

}
