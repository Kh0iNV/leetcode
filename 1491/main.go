//Leetcode - Problem 1491 - Average salary excluding the minimum and maximum salary
package main

// O(N)
// O(1)
func avarage(salary []int) (result float64) {
  min := salary[0]
  max := salary[0]
  total := 0

  for _, val := range salary {
    if val < min {
      min = val
    }
    if val > max {
      max = val
    }
    total += val
  }
  if len(salary) == 2 {
    return 0.0
  }
  result = float64((total - max - min)/(len(salary)-2))
  return
}
