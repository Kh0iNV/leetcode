//Leetcode - Problem 1518 - Water Bottles
package main

// Với n là số chai nước đầu vào thì thuật toán không thể vượt qua n step => O(logN)
// O(1)
func numWaterBottles(numBottles int, numExchange int) int {
  var result int
  var emptyBottles int

  // first drink
  result = numBottles
  emptyBottles = numBottles

  for emptyBottles >= numExchange {
    // exchange
    numBottles = emptyBottles / numExchange
    emptyBottles = emptyBottles % numExchange

    result += numBottles
    emptyBottles += numBottles
  }
  return result
}

// Cách 2: Dùng đệ quy
func numWaterBottles2(numBottles int, numExchange int) int {
  // first drink
  return numBottles + calcNumWater(numBottles/numExchange, numBottles%numExchange, numExchange) 
}

func calcNumWater(exchangedBottles int, emptyBottles int, numExchange int) int {
  if (exchangedBottles == 0) { return 0 }
  nextEmptyBottles := exchangedBottles + emptyBottles
  return exchangedBottles + calcNumWater(nextEmptyBottles/numExchange, nextEmptyBottles%numExchange, numExchange)
}
