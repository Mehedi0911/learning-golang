package main

func add(params ...int) int {
	total := 0

	for _, v := range params {
		total += v
	}
	return total
}

func main() {

	nums := []int{5, 2, 15, 50, 10}
	res := add(nums...)
	println("Sum:", res)
}
