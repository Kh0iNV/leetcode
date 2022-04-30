package main

import "testing"

func TestXorOperation(t *testing.T) {
  testcase := []struct {
    input1    int
    input2    int
    expected  int
  } {
    {5, 0, 8},
    {4, 3, 8},
    {1, 7, 7},
  }
  for _, tc := range testcase {
    if output := xorOperation(tc.input1, tc.input2); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }

}