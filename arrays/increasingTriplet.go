package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	solution := false

	first, second := math.MaxInt32, math.MaxInt32

	for i := range nums {
		if nums[i] <= first {
			first = nums[i]
		} else if nums[i] <= second {
			second = nums[i]
		} else {
			solution = true
			break
		}
	}

	fmt.Printf("%d %d\n", first, second)
	return solution
}

func increasingTripletDriver() {
	nums := []int{2, 1, 5, 0, 4, 6}
	res := increasingTriplet(nums)
	fmt.Println(res)
}
