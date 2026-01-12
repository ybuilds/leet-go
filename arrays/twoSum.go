package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	complement := make(map[int]int)

	for i := range nums {
		comp := target - nums[i]
		_, ok := complement[comp]
		if ok {
			res[1] = i
			res[0] = complement[comp]
			break
		}
		complement[nums[i]] = i
	}

	return res
}

func twoSumDriver() {
	nums := []int{3, 2, 4}
	target := 6
	res := twoSum(nums, target)

	for i := range res {
		fmt.Printf("%d ", res[i])
	}
	fmt.Println()
}
