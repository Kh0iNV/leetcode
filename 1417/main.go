//Leetcode - Problem 1417 - Reformat the string
package main

import "fmt"

// O(N)
// O(N)
func reformat(s string) string {
  var result string
  var letter []string
  var digit []string

  for _, c := range s {
    if 48 <= c && c <= 57 {
      digit = append(digit, fmt.Sprintf("%c", c))
    }
    if 97 <= c && c <= 122 {
      letter = append(letter, fmt.Sprintf("%c", c))
    }
  }
  numOfLetter := len(letter)
  numOfDigit := len(digit)

  if (numOfLetter - numOfDigit >  1) || (numOfDigit - numOfLetter > 1) {
    return ""
  }

  var lenOfResult int
  if numOfLetter > numOfDigit {
    lenOfResult = numOfDigit
  } else {
    lenOfResult = numOfLetter
  }

  for i := 0; i < lenOfResult; i++ {
    result += digit[i]
    result += letter[i]
  }

  if numOfLetter > numOfDigit {
    result = letter[lenOfResult] + result
  }
  if numOfDigit > numOfLetter {
    result += digit[lenOfResult]
  }
  return result
}

// In ra mã ASCII của các kí tự
// fmt.Printf("%d\n", 'a')
// fmt.Printf("%d\n", 'z')

//Nối chuỗi s += char
//Nối mảng chuỗi arrS = append(arrS, "s")
