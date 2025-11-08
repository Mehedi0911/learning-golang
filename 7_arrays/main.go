package main

import "fmt"

func main() {
	var arr [5]int // array of 5 integers
	arr[3] = 10
	fmt.Println("Initial array:", arr)

	var arr2 = [...]string{"Go", "Python", "Java", "C++"}

	fmt.Println("Initial array:", arr2)

	arr3 := [4]float64{3.14, 1.59, 2.65}
	arr3[3] = 9.26
	fmt.Println("Initial array:", arr3)

	fmt.Println("Length of arr3:", len(arr3))

	// for i := 0; i < len(arr2); i++ {
	// 	fmt.Println("Element", i, "is", arr2[i])
	// }

	for index, value := range arr3 {
		fmt.Println(index, value)
	}
}
