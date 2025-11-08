package main

// func counter() func() int {
// 	count := 0
// 	return func() int {
// 		count++
// 		return count
// 	}

// }

func sayKalimah() func() string {
	kalimahs := []string{"La", "Ilaha", "Illallah"}
	kalimah := ""

	return func() string {
		for _, k := range kalimahs {
			kalimah += k + " "
		}
		return kalimah
	}
}

func main() {
	// increment := counter()
	// increment2 := counter()
	// fmt.Println(increment())  // 1
	// fmt.Println(increment())  // 2
	// fmt.Println(increment())  // 3
	// fmt.Println(increment2()) // 3
	kalimah := sayKalimah()
	println(kalimah()) // La Ilaha Illallah
	println(kalimah()) // La Ilaha Illallah La Ilaha Illallah .... and so on
}
