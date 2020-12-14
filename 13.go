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
  t := (int64(input[0]))
  i := 1

  for {
    if (float64(t) > math.Pow10(i)) {
      fmt.Printf("Over %d\n", int64(math.Pow10(i)))
      i++
    }
    found := true

    for idx, time := range input {
      if time == 0 || idx == 0 {
        continue
      }

      if (t + int64(idx)) % int64(time) != 0 {
        found = false
        break
      }
    }

    if !found {
      t += int64(input[0])
    } else {
      break
    }
  }

  return t
}
