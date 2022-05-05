//Leetcode - Problem 1455 - Check if a word occurs as a prefix of any word in a Sentence
package main

func isPrefixOfWord(sentence string, searchWord string) int {
  index := 1
  j := 0
  for i := 0; i < len(sentence); i++ {
    if sentence[i] == ' ' {
      index++
    }

    if sentence[i] == searchWord[j] {
      if i == 0 || sentence[i-1] == ' ' || j != 0 {
        j++
      }
    } else {
        j = 0
    }

    if j == len(searchWord) {
      return index
    }
  }
  return -1
}
