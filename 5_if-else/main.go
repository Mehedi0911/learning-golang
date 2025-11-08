package main

import "fmt"

func main() {
	age := 19

	if age > 18 {
		fmt.Println("You are an adult.")
	} else if age == 16 {
		fmt.Println("You are almost an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	var role = "admin"
	var hasPermission = true

	if role == "admin" && hasPermission {
		fmt.Println("Access granted.")
	} else if role == "admin" || !hasPermission {
		fmt.Println("Access denied: insufficient permissions.")
	}

	// declare variables withing the if else statement

	if hasAuth := true; hasAuth {
		fmt.Println("User is authenticated.")
	}

	if color := "blue"; color == "red" {
		fmt.Println("The color is", color)
	} else if color == "blue" {
		fmt.Println("The color is", color)
	}
}
