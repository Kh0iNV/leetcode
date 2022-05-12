package main

import "testing"

func TestCountLargestGroup (t *testing.T) {
   testcases := []struct{
    input       int
    expected    int
  } {
    {13, 4},
    {2, 2},
    {15, 6},
    {24, 5},
  }
  for _, tc := range testcases {
    if output := countLargestGroup(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
