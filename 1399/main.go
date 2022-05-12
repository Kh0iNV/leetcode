//Leetcode - Problem 1399 - Count largest group
package main

// O(Nlog(N))
// O(log(N))
func countLargestGroup(n int) int {
  freq := make(map[int]int, n)
  max := 0
  count := 0

  for i := 1; i <= n; i++ {
    sum := getSum(i)
    freq[sum]++
    if max < freq[sum] {
      max = freq[sum]
    }
  }

  for _, el := range freq {
    if el == max {
      count++
    }
  }
  return count
}

// O(log(N))
func getSum(n int) int {
  if n < 10 {
    return n
  }
  sum := 0
  for n != 0 {
    sum = sum + n % 10
    n = n / 10
  }
  return sum
}
