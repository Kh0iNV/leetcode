package main

import "testing"

func TestCountOdds(t *testing.T)  {
   var testcase = []struct {
    input1    int
    input2    int
    expected  int
  } {
    {3, 7, 3},
    {8, 10, 1},
  }

  for _, tc := range testcase {
    if output := countOdds(tc.input1, tc.input2); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
