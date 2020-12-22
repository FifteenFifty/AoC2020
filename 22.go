package main

import (
  "fmt"
  "regexp"
  "strconv"
  "strings"
)

func main() {
  input := `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

input = `Player 1:
29
25
9
1
17
28
12
49
8
15
41
31
39
24
40
23
6
21
13
45
20
2
42
47
10

Player 2:
46
27
44
18
30
50
37
11
43
35
34
4
22
7
33
16
36
26
48
19
38
14
5
3
32`

  playerMatcher := regexp.MustCompile("Player 1:\n([0-9\n]+)\n\nPlayer 2:\n([0-9\n]+)")
  playerMatches := playerMatcher.FindStringSubmatch(input)

  p1Deck := []int {}
  p2Deck := []int {}

  for _, vStr := range strings.Split(playerMatches[1], "\n") {
    v, _ := strconv.Atoi(vStr)
    p1Deck = append(p1Deck, v)
  }
  for _, vStr := range strings.Split(playerMatches[2], "\n") {
    v, _ := strconv.Atoi(vStr)
    p2Deck = append(p2Deck, v)
  }

  p1 := P1(p1Deck, p2Deck)

  fmt.Println(p1Deck)
  fmt.Println(p2Deck)
  p2 := P2(input)

  fmt.Printf("P1: %d, P2: %d\n", p1, p2)
}

func P1 (p1Deck []int, p2Deck []int) int {
  ret := 0

  for {
    var p1 int
    var p2 int

    p1, p1Deck = p1Deck[0], p1Deck[1:]
    p2, p2Deck = p2Deck[0], p2Deck[1:]

    if p1 > p2 {
      p1Deck = append(p1Deck, p1, p2)
    } else {
      p2Deck = append(p2Deck, p2, p1)
    }

    if len(p1Deck) == 0 {
      ret = Score(p2Deck)
      break
    } else if len(p2Deck) == 0 {
      ret = Score(p1Deck)
      break
    }
  }

  return ret
}

func Score(deck []int) int {
  ret := 0
  for idx, v := range deck {
    ret += v * (len(deck) - idx)
  }
  return ret
}

func P2 (input string) int {
  ret := 0

  return ret
}

