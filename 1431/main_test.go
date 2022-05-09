package main

import "testing"

func TestKidsWithCandies(t *testing.T) {
   testcases := []struct{
    input1      []int
    input2      int
    expected    []bool
  } {
    {[]int{2,3,5,1,3}, 3, []bool{true,true,true,false,true}},
    {[]int{4,2,1,1,2}, 1, []bool{true,false,false,false,false}},
    {[]int{12,1,12}, 10, []bool{true,false,true}},
  }
  for _, tc := range testcases {
    output := kidsWithCandies(tc.input1, tc.input2)
    if len(output) != len(tc.expected) {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
    for i, val := range output {
      if val != tc.expected[i] {
        t.Errorf("got %v, wanted %v", output, tc.expected)
      }
    }
  }
}
