package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other number")
	}

	// multiple values in case

	day := time.Now().Weekday()

	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("Today is", day, "It's the weekend!")
	default:
		fmt.Println("Today is", day, "It's a weekday.")
	}

}
