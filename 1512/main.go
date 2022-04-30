//Leetcode - Problem 1512 - Number of good pairs
package main

// O(N^2)
// O(1)
func numIdenticalPairs(nums []int) int {
  var count = 0
  for i, val := range nums {
    for j := i + 1; j < len(nums); j++ {
      if (val == nums[j]) {
        count += 1
      }
    }
  }
  return count
}

// Cách 2: tính số lần xuất hiện => suy ra số cặp
// O(N)
// O(N)
func numIdenticalPairs2(nums []int) int {
  var count = 0
  var freq = make(map[int]int)

  for _, num := range nums {
    val, exists := freq[num]
    if exists {
      count += val
      freq[num] = val + 1
    } else {
      freq[num] = 1
    }
  }
  return count
}
