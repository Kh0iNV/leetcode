//Leetcode - Problem 1403 - Minimum subsequence in non-increasing order
package main

import "sort"

// O(Nlog(N))
// O(N)
func minSubsequence(nums []int) []int {
  var result []int
  sum := 0
  for _, val := range nums {
    sum += val
  }
  // độ phức tạp trung bình của các thuật toán sắp xếp Nlog(N)
  sort.Slice(nums, func(i, j int) bool {
    return nums[i] > nums[j]
  })
  temp := 0
  for i := 0; i < len(nums); i++ {
    if temp <= sum - temp {
      result = append(result, nums[i])
      temp += nums[i]
    } else {
      break
    }
  }
  return result
}


/*
                Best      Average     Worst  
Selection Sort  n^2       n^2         n^2    
Bubble Sort     n         n^2         n^2    
Insertion Sort  n         n^2         n^2    
Heap Sort       nlog(n)   nlog(n)     nlog(n)
Quick Sort      nlog(n)   nlog(n)     n^2    
Merge Sort      nlog(n)   nlog(n)     nlog(n)
Bucket Sort     n+k       n+k         n^2    
Radix Sort      nk        nk          nk     
Count Sort      n+k       n+k         n+k    
*/
