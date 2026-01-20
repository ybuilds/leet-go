package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := make([]float64, len(nums1)+len(nums2))

	var i, j, k int

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			res[k] = float64(nums1[i])
			i++
		} else {
			res[k] = float64(nums2[j])
			j++
		}
		k++
	}

	for i < len(nums1) {
		res[k] = float64(nums1[i])
		i++
		k++
	}

	for j < len(nums2) {
		res[k] = float64(nums2[j])
		j++
		k++
	}

	fmt.Println(res)

	if len(res)%2 == 0 {
		elem1 := float64(res[len(res)/2])
		elem2 := float64(res[len(res)/2-1])
		return (elem1 + elem2) / 2.0
	} else {
		return res[len(res)/2]
	}
}

func findMedianSortedArraysDriver() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}

	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
}
