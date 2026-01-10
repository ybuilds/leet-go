package main

import "fmt"

func ascendingSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func descendingSorted(arr []int) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func maxProfit2(prices []int) int {
	if ascendingSorted(prices) {
		return prices[len(prices)-1] - prices[0]
	}
	if descendingSorted(prices) {
		return 0
	}
	profit := 0

	return profit
}

func maxProfit2Driver() {
	prices := []int{7, 6, 4, 3, 1}
	profit := maxProfit2(prices)
	fmt.Println(profit)
}
