package main

import "testing"

func TestBuildArray(t *testing.T) {
   testcases := []struct{
    input1      []int
    input2      int
    expected    []string
  } {
    {[]int{1,3}, 3, []string{"Push","Push","Pop","Push"}},
    {[]int{1,2,3}, 3, []string{"Push","Push","Push"}},
    {[]int{1,2}, 4, []string{"Push","Push"}},
    {[]int{2,3,4}, 4, []string{"Push","Pop","Push","Push","Push"}},
  }
  for _, tc := range testcases {
    output := buildArray2(tc.input1, tc.input2)
    if len(output) != len(tc.expected) {
      t.Errorf("got %v, wanted %v", output, tc.expected)
    }
    for i, val := range output {
      if val != tc.expected[i] {
        t.Errorf("got %v, wanted %v", output, tc.expected)
      }
    }
  }
}

// So sánh 2 mảng thì kiểm tra len 2 mảng trước
// output [Push Push], expected [Push Push Push] vẫn pass nếu như không check len
