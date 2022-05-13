package main

import "testing"

func TestFindTheDistanceValue(t *testing.T) {
   testcases := []struct{
    input1      []int
    input2      []int
    input3      int
    expected    int
  } {
    {[]int{4,5,8}, []int{10,9,1,8}, 2, 2},
    {[]int{1,4,2,3}, []int{-4,-3,6,10,20,30}, 3, 2},
    {[]int{2,1,100,3}, []int{-5,-2,10,-3,7}, 6, 1},
  }
  for _, tc := range testcases {
    if output := findTheDistanceValue2(tc.input1, tc.input2, tc.input3); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
