package main

import (
  "fmt"
  "strconv"
)

// O(M*N)
// O(M+N)
func luckyNumbers(matrix [][]int) []int {
  var result []int
  m := len(matrix)
  n := len(matrix[0])
  luckByRow := make([]string, m)
  luckByCol := make([]string, n)

  for i := 0; i < m; i++ {
    min := matrix[i][0]
    index := fmt.Sprintf("%d0", i)
    for j := 1; j < n; j++ {
      if matrix[i][j] < min {
        min = matrix[i][j]
        index = fmt.Sprintf("%d%d", i, j)
      }
    }
    luckByRow[i] = index
  }

  for j := 0; j < n; j++ {
    max := matrix[0][j]
    index := fmt.Sprintf("0%d", j)
    for i := 1; i < m; i++ {
      if matrix[i][j] > max {
        max = matrix[i][j]
        index = fmt.Sprintf("%d%d", i, j)
      }
    }
    luckByCol[j] = index
  }

  for _, row := range luckByRow {
    for _, col := range luckByCol {
      if row == col {
        i, _ := strconv.ParseInt(fmt.Sprintf("%c", row[0]), 10, 64)
        j, _ := strconv.ParseInt(fmt.Sprintf("%c", row[1]), 10, 64)
        result = append(result, matrix[i][j])
      }
    }
  }
  return result
}
