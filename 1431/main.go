//Leetcode - Problem 1431 - Kids with the greatest number of candies
package main

// O(N)
// O(1)
func kidsWithCandies(candies []int, extraCandies int) []bool {
  n := len(candies)
  result := make([]bool, n)
  max := candies[0]
  for i := 1; i < n; i++ {
    if max < candies[i] {
      max = candies[i]
    }
  }
  for i := 0; i < n; i++ {
    if candies[i] + extraCandies >= max {
      result[i] = true
    }
  }
  return result
}
