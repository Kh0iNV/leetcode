package main

import "testing"

func TestMinStartValue(t *testing.T) {
   testcases := []struct{
    input       []int
    expected    int
  } {
    {[]int{-3,2,-3,4,2}, 5},
    {[]int{1,2}, 1},
    {[]int{1,-2,-3}, 5},
  }
  for _, tc := range testcases {
    if output := minStartValue(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
