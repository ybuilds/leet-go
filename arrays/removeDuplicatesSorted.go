package main

import "fmt"

func removeDuplicates(nums []int) int {
	write, read := 0, 1

	for read < len(nums) {
		if nums[write] != nums[read] {
			write++
			nums[write] = nums[read]
		}
		read++
	}

	return write + 1
}

func removeDuplicatesDriver() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	res := removeDuplicates(nums)
	fmt.Println(res)

	for i := range res {
		fmt.Printf("%d ", nums[i])
	}
	fmt.Println()
}
