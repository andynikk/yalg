package main

import (
	"strconv"
	"strings"
)

// <template>
func maxProfit(res *[]byte) string {

	lines := strings.Split(string(*res), "\n")
	//n, _ := strconv.Atoi(lines[0])
	prices := strings.Split(strings.Replace(lines[1], "\r", "", -1), " ")

	profit := 0
	for i := 1; i < len(prices); i++ {
		priceBefore, _ := strconv.Atoi(prices[i-1])
		priceNow, _ := strconv.Atoi(prices[i])

		if priceBefore < priceNow {
			profit += priceNow - priceBefore
		}
	}

	return strconv.Itoa(profit) + "\n"
}
