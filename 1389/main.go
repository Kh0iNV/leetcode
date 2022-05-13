//Leetcode - Problem 1389 - Create target array in the given order
package main

// O(N^2)
// O(N)
func createTargetArray(nums []int, index []int) []int {
  var result []int
  for i := 0; i < len(nums); i++ {
    tValue := nums[i]
    tIndex := index[i]
    if tIndex == i {
      result = append(result, tValue)
    } else {
      result = insert(result, tIndex, tValue)
    }
  }
  return result
}

// O(N)
// O(1)
func insert(arr []int, index int, value int) []int {
  for i := index; i < len(arr); i++ {
    temp := arr[i]
    arr[i] = value
    value = temp
  }
  // Last element in array
  arr = append(arr, value)
  return arr
}
