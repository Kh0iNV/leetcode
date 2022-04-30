package main

import "testing"

func TestAvarage(t *testing.T) {
  testcase := []struct {
    input       []int
    expected    float64
  } {
    {[]int{4000,3000,1000,2000}, 2500},
    {[]int{1000, 2000, 3000}, 2000},
    {[]int{6000, 5000, 4000, 3000, 2000, 1000}, 3500},
    {[]int{8000, 9000, 2000, 3000, 6000, 1000}, 4750},
  }
  for _, tc := range testcase {
    if output := avarage(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
