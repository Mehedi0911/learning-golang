package main

import (
	"fmt"
)

func main() {
	//creating maps

	// m := make(map[string]string)
	// // fmt.Println(m)
	// m["name"] = "Golang"
	// m["type"] = "Programming Language"
	// m["origin"] = "Google"
	// fmt.Println(m)
	// fmt.Println(m["name"])

	// names := make(map[string]int)
	// names["Alice"] = 25
	// names["Bob"] = 30
	// names["Charlie"] = 35
	// fmt.Println(names)
	// fmt.Println(len(names))
	// delete(names, "Bob")
	// fmt.Println(names)

	m := map[string]string{
		"name":   "Golang",
		"type":   "Programming Language",
		"origin": "Google",
	}
	fmt.Println(m)

	students := map[string]int{"mehedi": 30, "jamal": 35, "kamal": 25}
	fmt.Println(students)

	k, ok := students["mehedi"]
	fmt.Println("mehedi's age:", k)
	if ok {
		fmt.Println("mehedi is present in the map")

	} else {
		fmt.Println("mehedi is not present in the map")
	}

	// fmt.Println(maps.Equal(m, students)) // can not compare maps with different value or key types

}
