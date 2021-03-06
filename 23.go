package main

import (
  "fmt"
)

func main() {
  input := []int { 3,8,9,1,2,5,4,6,7, }
  input = []int { 3, 9, 8, 2, 5, 4, 7, 1, 6 }

  inputCopy := make([]int, len(input))
  copy(inputCopy, input)

  p1 := P1(inputCopy, 100)

  for i := 10; i <= 1000000; i++ {
    input = append(input, i)
  }

  p2 := P2(input, 10000000)

  fmt.Printf("P1: %d, P2: %d\n", p1, p2)
}

func P1 (input []int, moves int) int {
  ret := 0
  i := 0

  end := len(input)

  for round := 0; round < moves; round++ {
    currentCup := input[i % len(input)]

    fmt.Printf("-- move %d --\n", round + 1)
    fmt.Printf("cups: ")
    for _, cup := range input {
      if cup == currentCup {
        fmt.Printf("(%d) ", currentCup)
      } else {
        fmt.Printf("%d ", cup)
      }
    }
    fmt.Printf("\n")

    /*
     * The crab picks up the three cups that are immediately clockwise of the
     * current cup. They are removed from the circle; cup spacing is adjusted as
     * necessary to maintain the circle.#
     */
    var a int
    var b int
    var c int
    fmt.Println("i is : ", i)
    input, a = Remove(input, (i + 1) % len(input))
    input, b = Remove(input, (i + 1) % len(input))
    input, c = Remove(input, (i + 1) % len(input))

    fmt.Printf("pick up: %d, %d, %d\n", a, b, c)

    /*
     * The crab selects a destination cup: the cup with a label equal to the
     * current cup's label minus one. If this would select one of the cups that
     * was just picked up, the crab will keep subtracting one until it finds a
     * cup that wasn't just picked up. If at any point in this process the value
     * goes below the lowest value on any cup's label, it wraps around to the
     * highest value on any cup's label instead.
     */
    destCup := currentCup
    for {
      if destCup == currentCup || destCup == a || destCup == b || destCup == c {
        destCup--
        if destCup < 1 {
          destCup = end
        }
      } else {
        break
      }
    }

    fmt.Printf("destination: %d\n\n", destCup)

    /*
     * The crab places the cups it just picked up so that they are immediately
     * clockwise of the destination cup. They keep the same order as when they
     * were picked up.
     */
     for j := 0; j < len(input); j++ {
       if input[j] == destCup {
         input = Insert(input, j + 1, a)
         input = Insert(input, j + 2, b)
         input = Insert(input, j + 3, c)
         break
       }
     }

     for j := 0; j < len(input); j++ {
       if input[0] == currentCup {
         i = 1
         break
       } else {
         var e int
         input, e = Remove(input, 0)
         input = Insert(input, end - 1, e)
       }
     }
  }

  for {
    var e int
    if input[0] == 1 {
      input, e = Remove(input, 0)
      break
    } else {
      input, e = Remove(input, 0)
      input = Insert(input, end - 1, e)
    }
  }

  fmt.Println("-- final --")
  for _, v := range input {
    fmt.Printf("%d", v)
  }
  fmt.Println()

  return ret
}

type Cup struct {
  val  int
  next *Cup
}

func P2 (input []int, turns int) int {

  // Element zero is wasted but we don't need to remember to subtract
  cups := make([]Cup, len(input) + 1)
  var prev *Cup

  for _, val := range input {
    cups[val] = Cup { val, nil }
    if (prev != nil) {
      prev.next = &cups[val]
    }
    prev = &cups[val]
  }

  // Complete the loop
  current  := &cups[input[0]]
  prev.next = current

  for round := 0; round < turns; round++ {
    // Pick up three
    a := current.next
    b := a.next
    c := b.next

    current.next = c.next

    destCupVal := current.val - 1
    for {
      if destCupVal < 1 {
        destCupVal = len(input)
      }
      if destCupVal == current.val || destCupVal == a.val || destCupVal == b.val || destCupVal == c.val {
        destCupVal--
      } else {
        break
      }
    }

    destCup := &cups[destCupVal]

    // Place the cups to the right of the destination
    c.next = destCup.next
    destCup.next = a

    current = current.next
  }

  // Result is the two cups following 1, multiplied together
  pos1 := cups[1]
  x := pos1.next.val
  y := pos1.next.next.val

  return x * y
}


func Remove(from []int, idx int) ([]int, int) {
  a := from[idx]
  return append(from[:idx], from[idx + 1:]...), a
}

func Insert(into []int, idx int, value int) []int {
  into = append(into[:idx + 1], into[idx:]...)
  into[idx] = value
  return into
}
