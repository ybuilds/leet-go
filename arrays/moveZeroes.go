package main

func moveZeroes(nums []int) {
	i, j := 0, 0
	for i = 0; i < len(nums); i++ {
		if nums[i] != 0 {
			i++
			continue
		}
		if i != j {
			nums[i], nums[j] = nums[j], nums[i]
		}
		i++
		j++
	}
}
