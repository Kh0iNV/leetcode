//Leetcode - Problem 1441 - Build an array with stack operations
package main

// O(N)
// O(N)
func buildArray(target []int, n int) []string {
  var result []string
  j := 0
  for i := 0; i < n; i++ {
    if target[j] != i+1 {
      result = append(result, "Push", "Pop")
    } else {
      result = append(result, "Push")
      j++
    }
    if j == len(target) {
      break;
    }
  }
  return result
}

func buildArray2(target []int, n int) []string {
  var result []string
  size := len(target)
  j := 0
  for i := 1; i <= target[size-1]; i++ {
    if target[j] != i {
      result = append(result, "Push", "Pop")
    } else {
      result = append(result, "Push")
      j++
    }
  }
  return result
}
