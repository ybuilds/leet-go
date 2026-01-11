package main

import (
	"fmt"
)

func findPeaks(mountain []int) []int {
	peaks := []int{}
	i := 1

	for i < len(mountain)-1 {
		for i < len(mountain)-1 && mountain[i] >= mountain[i-1] {
			i++
		}

		if i-1 > 0 && i < len(mountain) && mountain[i] < mountain[i-1] {
			peaks = append(peaks, i-1)
		}

		for i < len(mountain)-1 && mountain[i] < mountain[i-1] {
			i++
		}
	}

	return peaks
}

func findPeaksDriver() {
	mountain := []int{4, 4, 1}
	peaks := findPeaks(mountain)

	for _, i := range peaks {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
