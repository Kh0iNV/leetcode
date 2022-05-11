//Leetcode - Problem 1408 - String matching in an array
package main

import "strings"

// O(N^2)
// O(N)
func stirngMatching(words []string) []string {
  var result []string
  for i := 0; i < len(words); i++ {
    for j := 0; j < len(words); j++ {
      if i == j { continue }
      if strings.Contains(words[i], words[j]) {
        result = append(result, words[j])
        break
      }
    }
  }
  return result
}
