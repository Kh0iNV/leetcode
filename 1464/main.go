//Leetcode - Problem 1464 - Maximum product of two elements in an array
package main

// O(N)
// O(1)
func maxProduct(nums []int) int {
  max := nums[0]
  secondMax := nums[0]
  for _, val := range nums {
    if val > max {
      secondMax = max
      max = val
    } else if val > secondMax {
      secondMax = val
    }
  }
  return (max-1)*(secondMax-1)
}
