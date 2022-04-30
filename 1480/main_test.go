package main

import "testing"

func TestRunningSum(t *testing.T) {
  testcase := []struct {
    input     []int
    expected  []int
  } {
    {[]int{1,2,3,4}, []int{1,3,6,10}},
    {[]int{1,1,1,1}, []int{1,2,3,4,5}},
    {[]int{3,1,2,10}, []int{3,4,6,16,17}},
  }
  for _, tc := range testcase {
    output := runningSum(tc.input)
    for i, val := range output {
      if tc.expected[i] != val {
        t.Errorf("got %v, wanted %v", output, tc.expected)
      }
    }
  }
}
// Không thể so sánh 2 slice như này: out != tc.expected