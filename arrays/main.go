package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)

	for i := range nums {
		fmt.Printf("%d ", nums[i])
	}
	fmt.Println()
}
