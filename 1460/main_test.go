package main

import "testing"

func TestCanBeEqual(t *testing.T) {
  testcase := []struct{
    input1    []int
    input2    []int
    expected  bool
  } {
    {[]int{1,2,3,4}, []int{2,4,1,3}, true},
    {[]int{7}, []int{7}, true},
    {[]int{1,12}, []int{12,1}, true},
    {[]int{3,7,9}, []int{3,7,11}, false},
    {[]int{1,1,1,1}, []int{1,1,1,1}, true},
  }
  for _, tc := range testcase {
    if output := canBeEqual(tc.input1, tc.input2); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
