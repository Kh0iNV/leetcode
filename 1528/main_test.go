package main

import "testing"

func TestRestoreString(t *testing.T) {
  var testcase = []struct {
    input1      string
    input2      []int
    expected      string
  } {
    {"codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}, "leetcode"},
  }
  for _, tc := range testcase {
    if output := restoreString(tc.input1, tc.input2); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
