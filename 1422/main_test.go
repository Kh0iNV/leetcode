package main

import "testing"

func TestMaxScore(t *testing.T) {
   testcases := []struct{
    input       string
    expected    int
  } {
    {"011101", 5},
    {"00111", 5},
    {"1111", 3},
  }
  for _, tc := range testcases {
    if output := maxScore(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
