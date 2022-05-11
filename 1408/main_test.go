package main

import "testing"

func TestStirngMatching(t *testing.T) {
   testcases := []struct{
    input       []string
    expected    []string
  } {
    {[]string{"mass","ass","hero","superhero"}, []string{"ass","hero"}},
    {[]string{"leetcode","et","code"}, []string{"et","code"}},
    {[]string{"blue","green","red"}, []string{}},
  }
  for _, tc := range testcases {
    output := stirngMatching(tc.input)
    if len(output) != len(tc.expected) {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
    for i := 0; i < len(output); i++ {
      if output[i] != tc.expected[i] {
        t.Errorf("got %v, wanted %v", output, tc.expected)
      }
    }
  }
}
