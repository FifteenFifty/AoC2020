package main

import (
    "strings"
    "fmt"
    "strconv"
    "math"
)

func main() {
  earliest := 1006697
  input    := "13,x,x,41,x,x,x,x,x,x,x,x,x,641,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,x,17,x,x,x,x,x,x,x,x,x,x,x,29,x,661,x,x,x,x,x,37,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,23"
  input = "17,x,13,19"

  strTimes := strings.Split(input, ",")
  times    := make([]int, len(input))

  for idx, timeStr := range strTimes {
    if timeStr == "x" {
      times[idx] = 0
      continue
    }

    time, _ := strconv.Atoi(timeStr)
    times[idx] = time
  }

  p1 := P1(times, earliest)
  p2 := P2(times)

  fmt.Printf("P1: %d, P2: %d\n", p1, p2)
}

func P1 (input []int, earliest int) int {
  ret          := 0
  smallestWait := 999999999999

  for _, time:= range input {
    if time == 0 {
      continue
    }

    if time - (earliest % time) < smallestWait {
      ret = time * (time - (earliest % time))
      smallestWait = time - (earliest % time)
    }
  }

  return ret
}

func P2 (input []int) int64 {
  //  - Always a multiple of the first
  largest    := 0
  largestIdx := 0

  // Look for a multiple of time that is idx off the beginning timestamp
  for idx, time := range input {
    if (time > largest) {
      time = largest
      largestIdx = idx
    }
  }



  return 1
}
