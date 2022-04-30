//Leetcode - Problem 1496 - Path crossing
package main

import "fmt"

type Point struct {
  x int
  y int
}

// Ý tưởng 1: tạo ra mảng lưu lại tọa độ các điểm trong hành trình
// Duyệt trong mảng, nếu có 2 điểm trùng nhau => return true
// O(N)
// O(N)
func isPathCrossing(path string) bool {
  x := 0
  y := 0
  pointPaths := make([]Point, len(path)+1)
  pointPaths[0] = Point{x, y}
  for i, rune := range path {
    switch rune {
    case 'N':
      y++
    case 'E':
      x++
    case 'S':
      y--
    case 'W':
      x--
    default:
      break
    }
    pointPaths[i+1] = Point{x, y}
  }
  return checkDuplicate(pointPaths)
}

// Hàm kiểm tra phần tử trùng lặp trong mảng
func checkDuplicate(list []Point) bool {
  visited := make(map[Point]bool)
  for _, point := range list {
      if _, isExist := visited[point]; isExist {
          return true
      }
      visited[point] = true
  }
  return false
}

// Ý tưởng 2: tạo ra map lưu lại tạo đổ hành trình "00" true
// O(N)
// O(N)
func isPathCrossing2(path string) bool {
  x := 0
  y := 0
  visitedPoints := make(map[string]bool, len(path)+1)
  visitedPoints["00"] = true
  for _, rune := range path {
    switch rune {
    case 'N':
      y++
    case 'E':
      x++
    case 'S':
      y--
    case 'W':
      x--
    default:
      break
    }
    keyPoint := fmt.Sprintf("%d%d", x, y)
    if _, isExist := visitedPoints[keyPoint]; isExist {
      return true
    }
    visitedPoints[keyPoint] = true
  }
  return false
}
