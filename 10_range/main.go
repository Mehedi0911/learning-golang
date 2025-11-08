package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50}

	sum := 0

	for i, num := range nums {
		fmt.Println(num, i)
		sum += num
	}

	// fmt.Println("Sum:", sum)

	// m := map[string]int{
	// 	"Alice":   30,
	// 	"Bob":     25,
	// 	"Charlie": 35,
	// }

	// for key, value := range m {
	// 	fmt.Println(key, value) // key and value
	// }
	// for k := range m {
	// 	fmt.Println(k) // only keys
	// }

	for i, c := range "golang" {
		fmt.Println(i, c) // index/starting byte of rune and rune (unicode code point rune)
		fmt.Println(string(c))
	}
}
