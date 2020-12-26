package main

import (
  "fmt"
)

func main() {
  cardKey := 15733400
  doorKey := 6408062

  p1 := P1(cardKey, doorKey)
  p2 := P2("")

  fmt.Printf("P1: %d, P2: %d\n", p1, p2)
}

func P1 (cardPubKey int, doorPubKey int) int {
  /*
   * The card transforms the subject number of 7 according to the card's secret
   * loop size. The result is called the card's public key
   */
  cardKeyLoopSize := FindLoopSize(7, cardPubKey)
  /*
   * The door transforms the subject number of 7 according to the door's secret
   * loop size. The result is called the door's public key.
   */
  doorKeyLoopSize := FindLoopSize(7, doorPubKey)

  /*
   * The card transforms the subject number of the door's public key according
   * to the card's loop size. The result is the encryption key.
   */
  cardEncKey := Transform(doorPubKey, 1, cardKeyLoopSize)

  /*
   * The door transforms the subject number of the card's public key according
   * to the door's loop size. The result is the same encryption key as the card
   * calculated.
   */
  doorEncKey := Transform(cardPubKey, 1, doorKeyLoopSize)

  if cardEncKey != doorEncKey {
    fmt.Println("Keys do not match!")
  }

  return doorEncKey
}

func P2 (input string) int {
  ret := 0

  return ret
}

func FindLoopSize(subject int, target int) int {
  loopSize := 0
  value    := 1

  for {
    value = Transform(subject, value, 1)

    loopSize++
    if value == target {
      break
    }
  }

  return loopSize
}

func Transform(subject int, value int, iterations int) int {
  for i := 0; i < iterations; i++ {
    value *= subject
    value %= 20201227
  }
  return value
}
