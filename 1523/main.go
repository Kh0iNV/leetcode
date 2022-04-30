//Leetcode - Problem 1523 - Count odd numbers in an interval range
package main

// O(1)
// O(1)
func countOdds(low int, high int) int {
  var count = (high - low)/2
  if (low % 2 == 0) && (high % 2 == 0) {
    return count
  }
  return count + 1
}
