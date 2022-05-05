package main

import "testing"

func TestIsPrefixOfWord(t *testing.T) {
  testcases := []struct {
    input1      string
    input2      string
    expected    int
  } {
    {"i love eating burger", "burger", 4},
    {"this problem is an easy problem", "pro", 2},
    {"i am tired", "you", -1},
  }
  for _, tc := range testcases {
    if output := isPrefixOfWord(tc.input1, tc.input2); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
