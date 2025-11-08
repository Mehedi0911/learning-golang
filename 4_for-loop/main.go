package main

import "fmt"

func main() {

	// while loop using for construct
	var i int = 0
	for i < 5 {
		fmt.Println("Iteration number:", i)
		i++
	}

	// traditional for loop
	for j := 0; j <= 5; j++ {
		if j%2 == 0 {
			continue // skip even numbers
		}
		fmt.Println("Count:", j)
	}

	//using range to iterate over a slice

	for i := range 3 {
		fmt.Println("Value from range:", i)
	}

}
