package main

import "fmt"

func main() {
	var message string = "Hello, Golang!"
	fmt.Println(message)
	var number int = 100
	fmt.Println(number)
	var pi float64 = 3.14159
	fmt.Println(pi)
	var isGoFun bool = true
	fmt.Println(isGoFun)
	var greeting string = "Go" + "Lang"
	fmt.Println(greeting)

	const name string = "Muaaz"
	fmt.Println(name)

	student := "Mehedi"
	fmt.Println(student)

	const (
		age     = 25
		country = "bangladesh"
	)

	fmt.Println("age => ", age, "country=> ", country)
}
