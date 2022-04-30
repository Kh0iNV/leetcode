//Leetcode - Problem 1480 - Running sum of 1D array
package main

// O(N)
// O(1)
func runningSum(nums []int) []int {
  for i := 1; i < len(nums); i++ {
    nums[i] += nums[i-1]
  }
  return nums
}

// Tận dụng input đầu vào, không cần tạo ra var resutl []int
