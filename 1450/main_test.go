package main

import "testing"

// O(N)
// O(1)
func TestBusyStudent(t *testing.T) {
  testcases := []struct {
    input1     []int
    input2     []int
    input3     int
    expected    int
  } {
    {[]int{1,2,3}, []int{3,2,7}, 4, 1},
    {[]int{4}, []int{4}, 4, 1},
    {[]int{4}, []int{4}, 5, 0},
    {[]int{1,1,1,1}, []int{1,3,2,4}, 7, 0},
    {[]int{9,8,7,6,5,4,3,2,1}, []int{10,10,10,10,10,10,10,10,10}, 5, 5},
  }

  for _, tc := range testcases {
    if output := busyStudent(tc.input1, tc.input2, tc.input3); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
