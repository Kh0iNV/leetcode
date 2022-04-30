//Leetcode - Problem 1486 - XOR operation in an array
package main

// O(N)
// O (1)
func xorOperation(n int, start int) (result int) {
  for i := 0; i < n; i++ {
    result ^= start + 2*i
  }
  return
}
