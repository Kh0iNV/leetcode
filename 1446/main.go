//Leetcode - Problem 1446 - Consecutive characters
package main

// O(N)
// O(1)
func maxPower(s string) int {
  result := 1
  count := 1
  for i := 1; i < len(s); i++ {
    if s[i] == s[i-1] {
      count++
    } else {
      count = 1
    }
    if result < count {
      result = count
    }
  }
  return result
}
