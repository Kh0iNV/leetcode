package main

import "testing"

func Test (t *testing.T)  {
  var testcase = []struct {
    input     string
    expected  string 
  } {
    {"1st Oct 2052", "2052-10-01"},
    {"12st Jan 2052", "2052-01-12"},
    {"20th Oct 2052", "2052-10-20"},
    {"6th Jun 1933", "1933-06-06"},
    {"26th May 1960", "1960-05-26"},
  }

  for _, tc := range testcase {
    if output := reformatDate(tc.input); output != tc.expected {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
  }
}
