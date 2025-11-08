package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func lang() (string, string, bool) {
	return "Golang", "Python", true
}

func main() {
	res := add(5, 5)
	fmt.Println("Sum:", res)

	v1, v2, v3 := lang()
	fmt.Println("Languages:", v1, v2, v3)
}
