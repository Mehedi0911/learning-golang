package main

import (
	"fmt"
	"slices"
)

func main() {

	// var numbers []int
	// fmt.Println(numbers, len(numbers), cap(numbers), numbers == nil)

	var numbers = make([]int, 2, 5)
	// fmt.Println("before append => ", numbers, len(numbers), cap(numbers), numbers == nil)

	var numbers2 = []int{1, 2, 3, 4, 5}
	// var names = []string{"A", "B", "C"}

	numbers = append(numbers, numbers2...)
	// fmt.Println("after append => ", numbers, len(numbers), cap(numbers), numbers == nil)

	var nums = make([]int, 0, 3)
	nums = append(nums, 10)
	var nums2 = make([]int, len(nums))
	fmt.Println(nums2, nums)
	copy(nums2, nums)
	fmt.Println(nums2, nums)

	var nums3 = nums2

	nums2 = append(nums2, 15, 12, 15)

	fmt.Println("nums3:", nums3)
	fmt.Println("nums2:", nums2)

	fmt.Println("numbers:", nums2[0:3])
	fmt.Println("numbers:", nums2[:2])
	var copied = nums2[2:]
	fmt.Println("copied:", copied)

	var newSlice = []string{"Go", "Python", "JavaScript"}
	var newSlice2 = []string{"Go", "Python", "JavaScript"}

	fmt.Println(slices.Equal(newSlice, newSlice2))

	var numbs = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(numbs)
}
