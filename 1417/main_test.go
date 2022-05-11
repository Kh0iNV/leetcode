package main

import "testing"

func TestReformat(t *testing.T) {
   testcases := []struct{
    input       string
    expected    string
  } {
    {"a0b1c2d3", "0a1b2c3d"},
    {"leetcode", ""},
    {"123425345346", ""},
    {"xinchao987654", "o9x8i7n6c5h4a"},
  }
  for _, tc := range testcases {
    if output := reformat(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
