package main

// Dùng mảng byte
func restoreString(s string, indices []int) string {
  var result = make([]byte, len(indices))
  for i := range result {
    result[indices[i]] = s[i]
  }
  return string(result)
}

// Dùng mảng rune
func restoreString2(s string, indices []int) string {
  var result = make([]rune, len(indices))
  for i, rune := range s {
    result[indices[i]] = rune
  }
  return string(result)
}

// O(N)
// O(N)
