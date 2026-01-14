package main

import (
	"fmt"
	"math"
)

func firstMissingPositive(nums []int) int {
	max := math.MinInt

	for _, i := range nums {
		if i > max {
			max = i
		}
	}

	if max < 0 {
		return 1
	}

	exist := make([]bool, max+1)

	for _, i := range nums {
		if i != math.MaxInt && i > 0 {
			exist[i] = true
		}
	}

	for i := range exist {
		if !exist[i] {
			return i
		}
	}

	return len(exist)
}

func firstPositiveDriver() {
	nums := []int{1, 2, 0}
	res := firstMissingPositive(nums)
	fmt.Println(res)
}
