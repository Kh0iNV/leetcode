package main

import "testing"

func TestIsPathCrossing(t *testing.T) {
  testcase := []struct {
    input     string
    expected  bool
  } {
    {"NES", false},
    {"NESWW", true},
  }
  for _, tc := range testcase {
    if output := isPathCrossing(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}

func TestIsPathCrossing2(t *testing.T) {
  testcase := []struct {
    input     string
    expected  bool
  } {
    {"NES", false},
    {"NESWW", true},
  }
  for _, tc := range testcase {
    if output := isPathCrossing2(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
