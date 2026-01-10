package main

import "fmt"

func maxProfit2(prices []int) int {
	i, profit := 0, 0
	valley, peak := 0, 0

	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		valley = prices[i]

		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]

		profit += peak - valley
	}

	return profit
}

func maxProfit2Driver() {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit2(prices)
	fmt.Println(profit)
}
