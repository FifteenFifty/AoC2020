package main

import (
    "fmt"
    "strings"
    "strconv"
)

func main() {
  input := `12,1,16,3,11,0`


  inputStrs := strings.Split(input, ",")
  inputs := make([]int, len(inputStrs))

  for idx, inputStr := range inputStrs {
    i, _ := strconv.Atoi(inputStr)
    inputs[idx] = i
  }

  fmt.Println(inputs)

  p1 := P1(inputs, 2020)
  p2 := P1(inputs, 30000000)

  fmt.Printf("P1: %d, P2: %d\n", p1, p2)
}

type NumTracker struct {
  lastSeen int
  prevSeen int
}

func P1 (input []int, maxTurn int) int {
  tracker := make(map[int]NumTracker)

  for idx, v := range input {
    tracker[v] = NumTracker { idx + 1, 0 }
  }

  prev := input[len(input) - 1]
  num  := 0

  fmt.Println(tracker)

  for i := len(input); i < maxTurn; i++ {
    t, ok := tracker[prev]
    num = 0

    if ok {
      // Seen before
      if t.prevSeen == 0 {
        num = 0
      } else {
        num = t.lastSeen - t.prevSeen
      }
      t.prevSeen = t.lastSeen
      t.lastSeen = i
      tracker[prev] = t
    }

    tN, ok := tracker[num]
    if ok {
      tN.prevSeen = tN.lastSeen
      tN.lastSeen = i + 1
      tracker[num] = tN
    } else {
      tracker[num] = NumTracker { i + 1, 0 }
    }

    prev = num
  }

  return num
}

func P2 (input string) int {
  return 1
}
