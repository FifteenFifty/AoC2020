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
  p2 := P2(p1Deck, p2Deck)

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

func P2 (p1Deck []int, p2Deck []int) int {
  _, score := RecursiveCombat(p1Deck, p2Deck, 1)

  return score
}

func RecursiveCombat(p1Deck []int, p2Deck []int, depth int) (bool, int) {
  var p1Wins bool
  score := 0

  seenCombos := make(map[string]bool)

  for {
    /*
     * Before either player deals a card, if there was a previous round in this
     * game that had exactly the same cards in the same order in the same
     * players' decks, the game instantly ends in a win for player 1. Previous
     * rounds from other games are not considered. (This prevents infinite games
     * of Recursive Combat, which everyone agrees is a bad idea.)
     */
    combo := fmt.Sprint(p1Deck) + "|" + fmt.Sprint(p2Deck)
    _, seen := seenCombos[combo]
    if seen {
      p2Deck = []int {}
    } else {
      seenCombos[combo] = true

      /*
       * Otherwise, this round's cards must be in a new configuration; the players
       * begin the round by each drawing the top card of their deck as normal.
       */
      var p1 int
      var p2 int

      p1, p1Deck = p1Deck[0], p1Deck[1:]
      p2, p2Deck = p2Deck[0], p2Deck[1:]

      /*
       * If both players have at least as many cards remaining in their deck as
       * the value of the card they just drew, the winner of the round is
       * determined by playing a new game of Recursive Combat (see below).
       */
      if len(p1Deck) >= p1 && len(p2Deck) >= p2 {

        /*
         * To play a sub-game of Recursive Combat, each player creates a new deck
         * by making a copy of the next cards in their deck (the quantity of cards
         * copied is equal to the number on the card they drew to trigger the
         * sub-game)
         */
        p1Copy := make([]int, p1)
        p2Copy := make([]int, p2)

        copy(p1Copy, p1Deck)
        copy(p2Copy, p2Deck)

        p1WinsRecurse, _ := RecursiveCombat(p1Copy, p2Copy, depth + 1)
        if p1WinsRecurse {
          // P1 wins
          p1Deck = append(p1Deck, p1, p2)
        } else {
          // P2 wins
          p2Deck = append(p2Deck, p2, p1)
        }
      } else {
        /*
         * Otherwise, at least one player must not have enough cards left in their
         * deck to recurse; the winner of the round is the player with the
         * higher-value card.
         */
        if p1 > p2 {
          p1Deck = append(p1Deck, p1, p2)
        } else {
          p2Deck = append(p2Deck, p2, p1)
        }
      }
    }

    if len(p1Deck) == 0 {
      p1Wins = false
      score  = Score(p2Deck)
      break
    } else if len(p2Deck) == 0 {
      p1Wins = true
      score  = Score(p1Deck)
      break
    }

  }

  return p1Wins, score
}
