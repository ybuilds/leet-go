package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)

	hashmap := make(map[int]int)

	for i := range nums {
		complement := target - nums[i]
		if _, ok := hashmap[complement]; ok {
			res[0] = i
			res[1] = hashmap[complement]
			break
		}
		hashmap[nums[i]] = i
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
