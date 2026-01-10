package main

import (
	"fmt"
)

func findPeaks(mountain []int) []int {
	peaks := []int{}

	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			peaks = append(peaks, i)
		}
	}

	return peaks
}

func findPeaksDriver() {
	mountain := []int{1, 4, 3, 8, 5}
	peaks := findPeaks(mountain)

	for _, i := range peaks {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
