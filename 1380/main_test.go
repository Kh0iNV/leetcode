package main

import "testing"

func TestLuckyNumbers(t *testing.T) {
  testcases := []struct {
    input     [][]int
    expected  []int
  } {
    {[][]int{{3,7,8},{9,11,13},{15,16,17}}, []int{15}},
    {[][]int{{1,10,4,2},{9,3,8,7},{15,16,17,12}}, []int{12}},
    {[][]int{{7,8},{1,2}}, []int{7}},
  }
  for _, tc := range testcases {
    output := luckyNumbers(tc.input)
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
