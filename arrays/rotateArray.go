package main

import "fmt"

func rotate(nums []int, k int) {
	for k > 0 {
		elem := nums[len(nums)-1]
		for i := len(nums) - 1; i >= 1; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = elem
		k--
	}
}

func rotateDriver() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3

	rotate(nums, k)

	for _, i := range nums {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
