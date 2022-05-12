//Leetcode - Problem 1394 - Find lucky integer in an array
package main

// O(N)
// O(N)
func findLucky(arr []int) int {
  result := -1
  freq := make(map[int]int)

  for _, val := range arr {
    freq[val]++
  }

  for key, el := range freq {
    if el == key {
      if result < key {
        result = key
      }
    }
  }
  return result
}
