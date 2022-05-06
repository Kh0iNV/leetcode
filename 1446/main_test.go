package main

import "testing"

func TestMaxPower(t *testing.T) {
  testcases := []struct{
    input     string
    expected  int
  } {
    {"leetcode", 2},
    {"abbcccddddeeeeedcba", 5},
    {"triplepillooooow", 5},
    {"hooraaaaaaaaaaay", 11},
    {"tourist", 1},
  }
  for _, tc := range testcases {
    if output := maxPower(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
