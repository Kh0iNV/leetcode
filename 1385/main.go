//Leetcode - Problem 1385 - Find the distance value between two arrays
package main

import "sort"

// O(N^2)
// O(1)
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
  count := 0
  for i := 0; i < len(arr1); i++ {
    for j := 0; j < len(arr2); j++ {
      if abs(arr1[i] - arr2[j]) <= d {
        count++
        break
      }
    }
  }
  return len(arr1) - count
}

func abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

//Dùng tìm kiếm nhị phân
// O(NlogN)
// O(1)
func findTheDistanceValue2(arr1 []int, arr2 []int, d int) int {
  result := 0\
  // NlogN
  sort.Ints(arr2)

  Outer:
  for _, val := range arr1 {
    left := 0
    right := len(arr2) - 1

    for left <= right {
      mid := (right + left)/2
      diff := abs(val - arr2[mid])

      if diff > d {
        if val > arr2[mid] {
          left = mid + 1
        } else {
          right = mid - 1
        }
      } else {
        continue Outer
      }
    }
    result++
  }
  return result
}
