package main

import "fmt"

func main() {

	//sl1 := []int{1, 2, 3}
	//sl2 := []int{4, 5}
	//sl3 := append(sl2, sl1...)
	//fmt.Println(len(sl3), cap(sl3))

	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr[1])

	arr1 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr1[1])

	a := [4]bool{true, false, true, false}
	fmt.Println(a[0])

	var ar [10]int
	fmt.Println(ar)

	var ars [4]bool = [4]bool{false, true, false, true}
	fmt.Println(ars[1], ars[3])

	var arrs [3]bool = [3]bool{true, false, true}
	arrs[0] = false
	fmt.Println(arrs[0])

	var A [10]int = [10]int{2, 3, 4, 5, 6, 7, -2, -1, 10}
	idx := 0
	for i, v := range A {
		if A[0] < v && v < A[9] {
			idx = i
			break
		}

		fmt.Println(idx)

	}

}
