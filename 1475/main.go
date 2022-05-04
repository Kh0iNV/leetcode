//Leetcode - Problem 1475 - Final prices with a special discount in a shop
package main

// Ý tưởng 1: dùng 2 for
// O(N^2)
// O(1)
func finalPrices(prices []int) []int {
  for i := 0; i < len(prices)-1; i++ {
    for j := i+1; j < len(prices); j++ {
      if prices[i] >= prices[j] {
        prices[i] = prices[i] - prices[j]
        break
      }
    }
  }
  return prices
}
