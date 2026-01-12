package main

import "fmt"

func zeroFilledSubarray(nums []int) int64 {
	i, j := 0, 0
	var count int64

	for i < len(nums) && j < len(nums) {
		for j < len(nums) && nums[j] != 0 {
			j++
		}
		i = j
		for i < len(nums) && nums[i] == 0 {
			i++
		}
		var diff int64 = (int64)(i - j)
		count += diff * (diff + 1) / 2

		j = i
	}

	return count
}

func zeroFilledSubarrayDriver() {
	nums := []int{0, 0, 0, 2, 0, 0}
	res := zeroFilledSubarray(nums)
	fmt.Println(res)
}
