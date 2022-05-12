package main

import "testing"

func TestFindLucky (t *testing.T) {
   testcases := []struct{
    input       []int
    expected    int
  } {
    {[]int{2,2,3,4}, 2},
    {[]int{1,2,2,3,3,3}, 3},
    {[]int{2,2,2,3,3}, -1},
    {[]int{5}, -1},
  }
  for _, tc := range testcases {
    if output := findLucky(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
