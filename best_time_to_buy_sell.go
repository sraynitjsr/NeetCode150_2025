package main

import "fmt"

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    minPrice := prices[0]
    maxProfit := 0

    for i := 1; i < len(prices); i++ {
        profit := prices[i] - minPrice
        if profit > maxProfit {
            maxProfit = profit
        }
        if prices[i] < minPrice {
            minPrice = prices[i]
        }
    }

    return maxProfit
}

func main() {
    prices := []int{7, 1, 5, 3, 6, 4}
    profit := maxProfit(prices)
    fmt.Println("Maximum Profit:", profit)
}
