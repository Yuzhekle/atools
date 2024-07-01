package main

import "fmt"

func main() {
	orders := []int{1, 2, 3, 4, 5, 6, 7, 8}
	cutOrders := cutOrders(orders)
	fmt.Println(cutOrders)
}

func cutOrders(orders []int) [][]int {
	var cutOrders [][]int
	start, rowCount := 0, 2
	for {
		end := start + rowCount
		if end >= len(orders) {
			cutOrders = append(cutOrders, orders[start:])
			break
		}
		cutOrders = append(cutOrders, orders[start:end])
		start += rowCount
	}
	return cutOrders
}
