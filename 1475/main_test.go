package main

import "testing"

func TestFinalPrices(t *testing.T) {
  testcase := []struct{
    input     []int
    expected  []int
  } {
    {[]int{8,4,6,2,3}, []int{4,2,4,2,3}},
    {[]int{1,2,3,4,5}, []int{1,2,3,4,5}},
    {[]int{10,1,1,6}, []int{9,0,1,6}},
  }
  for _, tc := range testcase {
    output := finalPrices(tc.input)
    for i, val := range output {
      if val != tc.expected[i] {
        t.Errorf("got %v, wanted %v", output, tc.expected)
      }
    }
  }
}
