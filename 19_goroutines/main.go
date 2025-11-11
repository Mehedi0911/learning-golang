package main

import (
	"fmt"
	"sync"
)

func task(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task", i, "is running")
}

func main() {
	var wg sync.WaitGroup

	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()

}
