package main

import "fmt"

func maxProfit(price []int) int {
	profit := 0
	minPrice := 10000

	for _, p := range price {
		if p < minPrice {
			minPrice = p
		}
		if p-minPrice > profit {
			profit = p - minPrice
		}
	}

	return profit
}

func maxProfitDriver() {
	price := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(price)
	fmt.Println(profit)
}
