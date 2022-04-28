package main

import "testing"

func TestCanMakeArithmeticProgression(t *testing.T) {
  var testcase = []struct {
    input     []int
    expected  bool
  } {
    {[]int{3, 5, 1}, true},
    {[]int{1, 2, 4}, false},
  }

  for _, tc := range testcase {
    if output := canMakeArithmeticProgression(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}

func TestCanMakeArithmeticProgression2(t *testing.T) {
  var testcase = []struct {
    input     []int
    expected  bool
  } {
    {[]int{3, 5, 1}, true},
    {[]int{1, 2, 4}, false},
  }

  for _, tc := range testcase {
    if output := canMakeArithmeticProgression2(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
