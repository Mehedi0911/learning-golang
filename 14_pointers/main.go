package main

import "fmt"

func changeNum(num *int) { // *int refers to the pointer/ memory location of argument
	fmt.Println("Before change:", *num)
	*num = 55 // dereferencing the pointer to change the value at that memory location
	fmt.Println("Inside function:", *num)
}

func main() {
	num := 10
	changeNum(&num) // passing the memory address of num using & operator
	fmt.Println("Outside function:", num)
}
