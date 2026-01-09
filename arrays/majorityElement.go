package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	frequency := make(map[int]int)

	for _, num := range nums {
		frequency[num] += 1
		if frequency[num] > (len(nums) / 2) {
			return num
		}
	}

	return -1
}

func majorityElementDriver() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
