package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			target := nums[i] - 1
			nums[target], nums[i] = nums[i], nums[target]
		}
	}

	for i := 0; i < n; i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}

	return n + 1
}

func firstMissingPositiveDriver() {
	nums := []int{1, 2, 0}
	res := firstMissingPositive(nums)
	fmt.Println(res)
}
