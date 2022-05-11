//Leetcode - Problem 1413 - Minimum value to get positive step by step sum
package main

// O(N)
// O(1)
func minStartValue(nums []int) int {
  min := nums[0]
  for i := 1; i < len(nums); i++ {
    nums[i] += nums[i-1]
    if min > nums[i] {
      min = nums[i]
    }
  }
  if min >= 0 {
    return 1
  }
  return 1 - min
}
