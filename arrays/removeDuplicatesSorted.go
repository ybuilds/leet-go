package main

import "fmt"

// func removeDuplicates(nums []int) int {
// 	i, j, count := 0, 1, 0

// 	for i < len(nums) && j < len(nums) {
// 		if nums[i] == nums[j] {
// 			j++
// 			continue
// 		}
// 		nums[i+1], nums[j] = nums[j], nums[i+1]
// 		i++
// 		j++
// 		count++
// 	}

// 	return count + 1
// }

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

	for _, i := range nums {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
