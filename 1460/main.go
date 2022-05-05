//Leetcode - Problem 1460 - Make two arrays equal by reversing sub-arrays
package main

// 2 điều kiện để 2 mảng có thể giống nhau
// - Tổng phần tử của 2 mảng bằng nhau
// - XOR các phần tử của 2 mảng = 0
// O(N)
// O(1)
func canBeEqual(target []int, arr []int) bool {
  sumTarget := 0
  sumArr := 0
  xor := 0
  for i := 0; i < len(target); i++ {
    sumTarget += target[i]
    sumArr += arr[i]
    xor = xor ^ target[i] ^ arr[i]
  }
  if sumTarget == sumArr && xor == 0 {
    return true
  }
  return false
}
