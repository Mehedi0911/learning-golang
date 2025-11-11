package main

import "fmt"

// func printSlice(items []int) {
// 	for _, v := range items {
// 		fmt.Println(v)
// 	}
// }
// func printSliceV2[T interface{}](items []T) { // [T any] // [T comparable]
// 	for _, v := range items {
// 		fmt.Println(v)
// 	}
// }
// func printSliceV3[T int | string](items []T) { // [T any]
// 	for _, v := range items {
// 		fmt.Println(v)
// 	}
// }

type Stack[T any] struct {
	elements []T
}

func main() {
	// printSlice([]int{1, 2, 3, 4, 5})
	// names := []string{"Alice", "Bob", "Charlie"}
	// printSliceV2(names)

	myStack := Stack[int]{
		elements: []int{1, 2, 3},
	}
	myStack2 := Stack[string]{
		elements: []string{"A", "B", "C"},
	}

	fmt.Println(myStack)
	fmt.Println(myStack2)
}
