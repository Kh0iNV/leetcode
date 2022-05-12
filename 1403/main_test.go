package main

import "testing"

func TestMinSubsequence(t *testing.T) {
   testcases := []struct{
    input       []int
    expected    []int
  } {
    {[]int{4,3,10,9,8}, []int{10,9}},
    {[]int{4,4,7,6,7}, []int{7,7,6}},
    {[]int{6}, []int{6}},
  }
  for _, tc := range testcases {
    output := minSubsequence(tc.input)
    if len(output) != len(tc.expected) {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
    for i := 0; i < len(output); i++ {
      if output[i] != tc.expected[i] {
        t.Errorf("got %v, wanted %v", output, tc.expected)
      }
    }
  }
}
