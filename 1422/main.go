//Leetcode - Problem 1422 - Maximum scope after splitting a string
package main

// O(N)
// O(1)
func maxScore(s string) int {
  result := 0
  maxSumRight := 0
  for i := 0; i < len(s); i++ {
    if s[i] == '1' {
      maxSumRight++
    }
  }

  sumLeft := 0
  for i := 0; i < len(s); i++ {
    if s[i] == '0' {
      sumLeft++
    } else {
      maxSumRight--
    }
    if result < sumLeft + maxSumRight {
      result = sumLeft + maxSumRight
    }
  }
  return result
}
