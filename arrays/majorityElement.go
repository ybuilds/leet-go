package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	candidate, count := nums[0], 0

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		} else if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func majorityElementDriver() {
	nums := []int{3, 2, 3}
	res := majorityElement(nums)
	fmt.Println(res)
}
