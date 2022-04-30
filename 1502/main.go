//Leetcode - Problem 1502 - Can make arithmetic progression from sequence
package main

import "sort"

// Cách này chậm, do phải sort trước, độ phức tạp = độ phức tạp của sort.Ints() = O(N^2) insertion sort
func canMakeArithmeticProgression(arr []int) bool {
  sort.Ints(arr)
  diff := 0
  for i := 0; i < len(arr) - 1; i++ {
    if diff == 0 {
      diff = arr[i] - arr[i+1]
      continue
    }
    if diff != arr[i] - arr[i+1] {
      return false
    }
  }
  return true
}

// (Số lớn nhất - Số bé nhất) / (số phần tử - 1) == step
// O(N)
// O(N)
func canMakeArithmeticProgression2(arr []int) bool {
  min := arr[0]
  max := arr[1]
  M := make(map[int]bool)
  N := len(arr)

  for i := 0; i < N; i++ {
    if min > arr[i] {
      min = arr[i]
    }
    if max < arr[i] {
      max = arr[i]
    }
    M[arr[i]] = true
  }

  if (max - min) % (N-1) != 0 {
    return false
  }

  diff := (max - min) / (N - 1)
  for min < max {
    min += diff
    _, isExist := M[min]
    if !isExist {
      return false
    }
  }
  return true
}
