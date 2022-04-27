package leetcode

func countOdds(low int, high int) int {
  var count = (high - low)/2
  if (low % 2 == 0) && (high % 2 == 0) {
    return count
  }
  return count + 1
}
