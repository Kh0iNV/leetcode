package main

import "testing"

func TestNumIdenticalPairs(t *testing.T) {
  var testcase = []struct{
    input     []int
    expected  int
  } {
    {[]int{1, 2, 3, 1, 1, 3}, 4},
    {[]int{1, 1, 1, 1}, 6},
    {[]int{1, 2, 3}, 0},
  }

  for _, tc := range testcase {
    if output := numIdenticalPairs(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}

func TestNumIdenticalPairs2(t *testing.T) {
  var testcase = []struct{
    input     []int
    expected  int
  } {
    {[]int{1, 2, 3, 1, 1, 3}, 4},
    {[]int{1, 1, 1, 1}, 6},
    {[]int{1, 2, 3}, 0},
  }

  for _, tc := range testcase {
    if output := numIdenticalPairs2(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
