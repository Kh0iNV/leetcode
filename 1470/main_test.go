package main

import "testing"

func TestShuffle(t *testing.T) {
  testcase := []struct {
    input1    []int
    input2    int
    expected  []int
  } {
    {[]int{2,5,1,3,4,7}, 3, []int{2,3,5,4,1,7}},
    {[]int{1,2,3,4,4,3,2,1}, 4, []int{1,4,2,3,3,2,4,1}},
    {[]int{1,1,2,2}, 2, []int{1,2,1,2}},
  }
  for _, tc := range testcase {
    output := shuffle(tc.input1, tc.input2)
    for i, val := range output {
      if val != tc.expected[i] {
        t.Errorf("got %v, wanted %v", output, tc.expected)
      }
    }
  }
}
