package main

import "testing"

func TestMaxProduct(t *testing.T) {
  testcase := []struct{
    input     []int
    expected  int
  } {
    {[]int{3,4,5,2}, 12},
    {[]int{1,5,4,5}, 16},
    {[]int{3,7}, 12},
  }

  for _, tc := range testcase {
    if output := maxProduct(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
