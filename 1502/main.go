package main

import "sort"
import "fmt"

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
func canMakeArithmeticProgression2(arr []int) bool {
  min := arr[0]
  max := arr[0]
  step := 0

  for i := 0; i < len(arr); i++ {
    localDiff := 0
    if min > arr[i] {
      localDiff = min - arr[i]
      min = arr[i]
    }
    if max < arr[i] {
      localDiff = arr[i] - max
      max = arr[i]
    }
    fmt.Println(localDiff)
    if localDiff != 0 && localDiff < step {
      step = localDiff
    }
  }
  fmt.Println(max, min, step)
  if step == 0 {
    return false
  }
  return (max - min) / (len(arr)-1) == step
}

func main() {
  a := canMakeArithmeticProgression2([]int{3, 9, 1, 5, 7})
  fmt.Println(a)
}
