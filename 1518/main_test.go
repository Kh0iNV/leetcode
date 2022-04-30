package main

import "testing"

func TestNumWaterBottles(t *testing.T) {
  var testcase = []struct {
    input1    int
    input2    int
    expected  int
  } {
    {9, 3, 13},
    {15, 4, 19},
    {120, 17, 127},
  }

  for _, tc := range testcase {
    if output := numWaterBottles(tc.input1, tc.input2); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}

func TestNumWaterBottles2(t *testing.T) {
  var testcase = []struct {
    input1    int
    input2    int
    expected  int
  } {
    {9, 3, 13},
    {15, 4, 19},
    {120, 17, 127},
  }

  for _, tc := range testcase {
    if output := numWaterBottles2(tc.input1, tc.input2); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
