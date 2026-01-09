package main

import "fmt"

func reversePart(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n

	reversePart(nums, 0, n-1)
	reversePart(nums, 0, k-1)
	reversePart(nums, k, n-1)
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
