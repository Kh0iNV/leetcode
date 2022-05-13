package main

import "testing"

func TestCreateTargetArray(t *testing.T) {
   testcases := []struct{
    input1      []int
    input2      []int
    expected    []int
  } {
    {[]int{0,1,2,3,4}, []int{0,1,2,2,1}, []int{0,4,1,3,2}},
    {[]int{1,2,3,4,0}, []int{0,1,2,3,0}, []int{0,1,2,3,4}},
    {[]int{1}, []int{0}, []int{1}},
  }
  for _, tc := range testcases {
    output := createTargetArray(tc.input1, tc.input2)
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
