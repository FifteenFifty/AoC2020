package main

import (
    "fmt"
    "strconv"
    "strings"
    "regexp"
)

func main() {
  input := `(9 * 3 + 4) * (7 + (2 + 3 + 6 * 2) + 3) * 4 * 5
  (8 + 4 + 9 + (7 * 9 + 4 + 6 + 9) * 3) + 6 + 5 + ((8 * 2) + 2 * 9) * 5
  5 * 6 * 8 * (9 + 7 * 8 * 9) + 6 + 2
  9 * 9 + (4 * (6 * 7 * 7 + 4 * 6) * 5 + 6 * (2 * 5)) * (8 + 7 + 6 + 2 * 7) + 5
  5 + (7 + 3 * 7 + 9 * (7 + 7 + 4 + 3 * 3)) + 8 * 4
  (6 + 2 + 5 + 4) * 2 * 9 * 2 * 8
  7 + 8 + ((4 + 6 + 6) * 6 + 9 + 3 * 8 + 9) + 5
  3 + ((6 * 6 * 7 * 6) * 6 + 6 * 4) + (2 + 5 * 6 + 6 + (2 * 4) + 6) * 9
  (2 * (4 * 9 * 3) * 8 + 6 + (7 + 2 + 6)) * 6 * (5 * 4 + 8 + 7 + (7 + 2 * 3))
  8 * ((7 + 4 + 9) * 5 + 6 + (4 + 9 + 8 * 3 * 8 + 3) * 5 * (5 + 8 + 7 + 3 * 7 * 4))
  6 * 5 * 5 + (8 * 5 * 8 * 5 + 7) + ((2 + 2 * 6 * 3) * 4) * 6
  2 + 2 * 9 + 8 * (3 * 7 * 7 * (2 + 5) * 3 + 2)
  6 + 2
  9 * 2 + ((6 * 7 * 7 * 2) + 8) + (7 * (4 * 6 + 7) * 2 * 2 * 9) + 4 * 8
  (7 * 9 + 9 * 2 * 4 * 9) + 4 * 2 * 9 + 6
  6 * 3 + 7 + 9 + 9 + ((6 + 6 + 7 * 5) * 6 * 8 * 5)
  ((8 * 8 * 6 + 6 * 4) + 3 + 3 + (3 * 5 + 3 * 5 + 4 + 7)) * 8 + (6 + 7 * 2 + 2) + (9 + 9)
  7 + 8 + 3 * 5 + (5 * 5 + 6 + 6) * (9 * 6)
  5 + 5 + 2
  ((3 + 3) * 2 + (5 + 8 * 2 + 4 * 3)) * 8
  ((5 + 6 + 5 * 8 + 4) + 9 + (5 + 5)) * (7 * 2 + (8 + 7)) + 9
  9 * (9 * 8) + (3 * 2 * (9 + 2 * 9 * 5) * 4 * 2) + 5
  4 * 4 + (2 + 5 + 5 * (3 + 6 * 4 + 7)) * (2 * 2) + 6
  6 + 2
  6 + 5 + (4 + 8) + 5
  6 * (4 * 6) + (4 * 5 + 4 + 5 + 9) + 7 * (2 + 7 + 9)
  9 * (3 * (5 + 9 * 4 * 5 + 3 * 9)) + ((4 * 4 * 3 * 6) * (9 * 8 * 3) * 9 * 9) + 3
  5 + (6 * (3 + 9 + 7 * 6 + 2) * (2 + 4 * 2 * 4 * 6) * 9 + 3) + 6 + 8 + 5 * 3
  8 + 9 * (8 + (3 + 2 + 4) * 4 + 3 * 4 * 5)
  9 + 5 + 8 + 9 * (8 + (4 * 2 * 9 + 8 + 7) * (8 + 6 * 3) * 7 * 7)
  6 * ((9 + 2) * (7 + 8 + 5 + 9 * 7) * 4) * 4 * ((7 + 7 + 4 * 5 * 2) + 2) + 9
  (4 + 9 * 8) * (6 + (7 * 6) * 9 + 8 * 2 * 7) + 4 + 9 + 3 * (6 + 8 + (8 * 4 + 7 * 3) * 8)
  (5 * 9 * 2 * 7 * 8 + 6) + (5 + 5)
  6 + 7 * 9 + 3
  ((7 + 8 * 3 * 7 + 5) * 7 * 5 + 6 * 6 + 5) + 9 + 4
  (5 * 2 * 4 * 9 * 2 + (8 + 6 * 7 + 6)) + ((4 * 4 * 6 * 9 + 7) + 9 * 7)
  (3 + (5 * 8 + 4 + 4 * 7)) * 7 + 8 + (7 + (3 * 5 * 4 + 6 + 9)) + ((7 + 2 + 7 * 9 + 6 * 8) + 3 + 3 + (2 * 9 * 7) + 4 + (6 + 3 * 5)) * 9
  5 * 4 + ((6 + 8 * 5 * 6) * 8) + 7 + 7
  5 * (4 * (2 * 9 * 7 * 7) * 5) + ((6 * 2 * 9 + 2) * 3 * 2 + 7 + 3) + 9 + 8
  (7 + 3 * 3 * 6 * 2 * 3) * 6 + 3 * ((3 * 4 * 3 * 9 * 7) + 9 + 8) * 9
  ((4 + 8 * 9) + 5 * 4 * 2 * 8 * 4) * (6 + 2 * (9 + 5) * 9 * 7 + 8)
  9 + 8 + (3 * 3 * (9 + 2 + 8 * 7 * 9))
  (5 * 8 * (7 * 2 + 8)) + 2 * 8 + 8 + 6 * 5
  2 + (2 * 9 * 5 + 4) * (5 + 9 * 3 * 4) * 4 + 2
  8 * 2 * (7 * 7 * 4 + 5 + 9 + 4)
  4 * 9 + 2 * 8
  6 + 2 + ((5 + 3 * 5 + 7 + 9) + 4 * 6 + 6) * 5
  (8 * 3 * 7) + (4 * 5 + 8 * 7 + 6)
  3 * 7 * ((3 + 3 * 5 + 9 + 7) + 9 * 2 + (4 + 3))
  (4 + 2 * 2 + 3 * 2 + 2) * 4 + 2 * 4 + 6 + 3
  ((9 * 7 * 7 + 8) * (2 * 5 + 2 + 3) * 2 * 4 * 7) * 9 + 5 + 5
  4 + 9 * 2 + 5 + (3 + 5 + (4 + 3 * 3 * 7 * 5 * 8) + 9 * 2 * 2) * (4 + (2 * 6) + 7 * 5 * 7 + 2)
  6 + 2 * 4 * (6 + 8) * (4 + 2) * 3
  3 + ((8 + 8 + 6) * 5) + 9 + 7 + ((7 * 2 * 9 + 5 * 3 + 5) + 3 + (2 * 5) + 3 + 5)
  4 + 2 * 6 + 6 + 6
  (5 * (4 * 9 * 8 + 5 * 8) * 6 * 9 + 3) * ((2 * 7 * 3 + 3 * 4 * 3) + 9 + 8 + 8) + 4
  4 * 4 + 9 * 9 + (2 * (6 * 5 + 9 + 7 * 3) + (2 * 8 + 3 + 2 + 3))
  6 * 4 + 5 * (6 * (6 * 7 + 2 * 6 + 4 * 7)) * 4 + ((5 * 9 * 8 * 9 * 5 * 5) * 5)
  3 + (5 * 4 + (3 + 6 * 2 + 8 + 7 * 8) + 3 + (9 + 8 * 6)) * (5 + 6 * (8 * 9 + 4 * 9 + 5) * 7 * 5) + 5
  9 + 2 + 9 + (3 + 9) + (4 * 4 + 7 * 5) + 5
  5 + 6
  6 + (7 + 2 * 6 * 9 * 7 + 5) * 9 + (4 * 5 * 8) + 2 * 3
  8 + 7 * (9 + (9 + 3 * 3 * 7) * (5 + 6 * 3 * 7 + 8 * 2)) * 7
  (5 * 7 + 5 + 3 + 6) * 8
  5 + ((9 + 5) + 9 + 4 * (5 * 4 * 7 + 3) * 3)
  5 * 8 * 7 + 8 * ((8 + 9 + 3 * 6 * 5) * 4 * 9 + 4)
  8 * (6 + (6 + 9 * 7 * 8 * 7) + 7 * 2) + ((2 * 9) * 7) + 4
  6 + ((4 + 5 + 6 + 7) * 7)
  8 * (5 * 6 * 4 * 8 + 6 * 6) + 2 * 9
  7 + 3 + 2 * (6 + 8)
  9 * ((4 + 4 * 9 + 6 * 8 * 7) + 9 + 2 + 6 * (4 * 9 + 7 + 6) * 3) * 5 + 7
  ((2 + 3 * 2 * 9 + 6) + 4 + 8 * 2 * 5 + 8) * 5 + 5 + (9 * 6) + 9 + (3 * 2 * 8)
  4 + (6 + 2) * 9 + (5 + 4 * (8 * 3 + 7 + 6) * 4) * ((7 + 2 + 3) * 6 + 9 * (6 + 7 * 2))
  6 + 2 + ((2 * 4 * 8) * 2 * 4 * 7 + 9) * (7 + 8 + (6 + 7 * 9 + 9 * 7) + 4) * 8
  9 * 9 + 7 + 9 * (9 * 4) * 9
  (7 * 3 + (3 * 7 + 4 * 4 * 8) * 2 * 2 * (3 * 3)) + 8
  9 * 6 + 3 * (2 * 5 * 7 * 7 + 7) * 8 + (2 + 4)
  (8 * 9) + 8 * 6
  (3 * 2 + 6 + (7 * 8 * 5 * 4)) + 5 * 6
  4 + 3 + (9 + 3 + (2 * 9 * 6) + 7) * 8 + 3 * 8
  9 + 6 * 8 + 4 + 9 * (9 * (8 * 4 + 9) + 3)
  5 + 4 * 6 * (8 * (2 + 9 + 6) * (6 * 3 + 2) + 6) * (8 * 8 * 3 * (6 * 9) * 3 + (4 + 4 + 2 * 2 * 6)) + 6
  4 + (6 + 6 * (4 * 7 * 7) + 5) * 5 + 6 * (7 + 2 * 4 * 2 + 9 + 9)
  ((8 + 7) * (7 + 5 * 9) + (9 + 2 * 6 * 2) + 6 * 7) * (3 + 3 * 2 + 3 + 7) * 4 + 7 * 7
  (6 + 9 * 7 * 4 + 7) * 6 + 7 + 5 + 3 + 8
  9 + 2 + (9 * (8 * 9 * 2 * 6 + 2 + 5)) + 6 + 4 + 8
  6 + 4 * ((2 * 9) * 6 * 5) + 2 * 9
  (6 + 8 * 2 + (5 + 2 * 7 * 7) * 5 + 4) * (2 * 3 + 8 * (9 + 6 * 4 + 3 + 6 + 4) * (2 + 5 * 8 * 2 + 4))
  3 + (9 * (2 * 7 * 6) * 9 + 7 + 3) + (7 * (5 * 6 + 7) + 3 * (5 * 7 * 5 * 4) * 8 + 6) * 4 * 5 + 9
  (6 + 6 + 4) + 3 * 4 + 3
  9 * (2 * 7 + 9 + (4 + 2) * 5) + (2 + 5 + (7 + 6)) * 8
  2 * (6 * 6 * 9 + 8) + 4 * 2
  ((4 * 3) + 8) + 5 + 5
  5 + 8 * 7 + 6 + 6 * ((9 * 2 * 5 + 4 * 7) + 8 + 6 + 5 + (3 + 7))
  8 + (8 * 8 * 8 * (3 * 7 * 8) * 4) * 8 + 5 * 8 * 3
  (8 * 3 * 8) * 4 * 3 * 3 * 6
  5 + 6 * (9 + (5 + 6 + 8 + 3 * 4) * (6 * 3 + 3 + 4 * 6 + 2))
  3 * (2 * (5 * 9) * 5 + 8 + 5 * 6) * 8 + (8 + 5 + 7 * (6 + 5 + 9 * 2 * 9) + 2 * (3 * 2 * 2 + 9 * 4))
  (2 * 8 + 5 + 7 + (8 + 2 + 7 * 5 * 5 + 4) * 2) + 5 * ((8 * 7 + 5) * 3 * 4 + (4 * 7 * 6) * (7 + 6 + 4 + 7 * 7) + (8 * 7 + 2 * 2 + 9 + 3)) + ((4 + 8) + 9)
  3 * (8 + (2 + 9) + 2 * 8 + (9 * 4 * 9)) * 2 * 3 * 4
  (4 + (9 * 4) + 5) * 8 * (3 + 7 * 3 * 7 + 4) * 7 * 2
  2 * 4 * (2 * 8 + 9 + 7 + (8 * 3 * 6 + 3 + 5) + 9) + (5 * 8 * 8 + (4 * 5 * 2 + 4 * 2 + 5) * 4) + 2 * 6
  7 + 8 + ((5 + 7 + 2 + 9 * 4 * 3) * 4) + 7 + 6 + 8
  ((4 + 4 * 4 * 9 * 6 + 9) * (7 + 9) * 2 + 8) * 8 * 6 + 9
  (2 * (6 + 7 * 3 + 6 * 7) + 4 * 2 + (3 + 8 + 9) * 8) * 8 + 8 + 8 + (2 + 8 * 5 * 9)
  (7 + 5 + (7 + 9 + 5) + 2) * 4 + 9
  3 * 6 + 8 + 2 * 7
  6 + 2 + ((7 * 6 + 3 + 4) * 8 * (2 + 6 * 5) * 6) + 4 * 2
  (4 * 8) + 7 * 3 + 3 * 7 + 5
  5 + 6 + 3 + 7 * (2 * 7) + (5 + (3 + 4 + 3 + 4 + 7) + 8)
  2 + ((4 + 7 * 6) + 8 + 9 * 6 * 9)
  (6 * 2) * (7 + 5 * (7 * 6 + 7 + 3 + 8 + 7) + 2 + 5 + 8) + 6
  (3 * 2 * 8 + 7) + (5 * 4 * 9 + 6) + 4 * 3 + 4
  8 + (8 + 2 * 3 + 2 + 2)
  (7 + 8 * 5) * (3 * (9 + 5 + 6 + 9 + 3 + 9) * 3) + 8 + 7 * (5 + 9 + 3 * 2 + 4 + 5) + 4
  7 * ((6 * 7 * 2 * 9 + 8) + 8 * 6 * 8 * 5) * 2 + 3 + 8 * 3
  (4 + (5 * 7 + 3 + 4 + 9)) * 7
  7 * 5 * 2 * (4 + 4 + 6 + 3 * 7) + (8 * (6 + 2 + 8 * 9 + 6) * 3 * 6 + 3 + (4 + 5 + 9)) * 7
  3 * 4 + 2
  ((6 + 9 + 5 + 4 + 9 * 7) * 9 + 8 + (9 * 2 * 6 * 4 * 7 * 3)) + (8 + 6 + 7 + 4) * 9 * (8 * 5) + 7
  (9 * 8 + 4) + (3 * 4 + (9 * 5 * 3 * 4) * 9)
  (7 + 8 * 3 + 6 * (4 + 5 + 7 * 8 + 9) + 4) * 3
  ((7 * 4) + (6 * 7 * 6) * 7 + 9 * 3) + 2 * (4 + 8) + (4 * 5 + 8 + 8 + 5) * (4 + 9) * 5
  (2 + 3 + 7) * 6 + 8 + (4 + 7 * 6 * (3 + 4 + 4 + 7 * 6) * 8 + 3) * 4
  6 * (8 * 3 + 4 * (6 + 9 * 2 + 3 + 8 * 5) * 8 + 6)
  8 + ((2 * 8 + 5 * 5 * 5 * 3) * 9 + 5) + 9
  (8 + 4 * 9 * 8) + 5 * 5 * (7 + 7 + 2 * 9 + 8 * 6) + 6 * (6 * 6 + 6 + 3 + 9 * 8)
  5 * 9 + (5 + (9 * 5 + 7) + (2 + 2 * 7) * 4) * (9 * 9 + 5 * 4) + 9
  7 + 7 * (2 * 9 * 6 + 3 + 5)
  2 + (8 * 9 * 7 * (4 * 9 + 9 * 4 + 5) + 4 * 4) + 5 + 2 + 8
  (2 + (4 + 9 + 4 * 7 * 3) + 4 + 8) + (2 + 4 * 5 * (3 + 4) + 8 * 5)
  2 + ((7 * 8 + 6 * 6) * 4 + 6 + 7 * 6 * 9) + ((7 + 9 * 5 * 5) + 4 * 9 + 6 + 7) * ((6 + 9 + 4 * 7 * 2) * 7 + 9 * (7 + 2 + 3 * 3 + 5) * 5)
  (3 * 8 + 2 * 2 + (3 * 7 + 8 * 5 * 6) + 6) + 4 + 5
  8 + 5 * 8 + 7 * 2
  (3 + (6 * 9 * 8 + 4 + 8) * 4 + 2) + 8
  (7 + 7) + 5 * (9 * 2 + 9) + ((4 * 3 * 3 * 4 + 4 * 9) * 8 * 8 + (3 * 9 + 7 + 8) * 3 + 9)
  (6 + (5 + 7 * 5) * 8 * 9) * 6 + 3
  3 + 2 * (5 * 8) + 5 + 6
  7 * 4 + 7 + 7 * 4 * (4 + 9 * 6 * 5)
  3 + 4 + (6 * 9) * (6 + (5 * 5 + 6 * 8) + 9 * (7 * 3 + 6)) + 8
  4 + 4 * 6 * ((4 * 3 * 2) * 4 + 3 + (7 + 8 + 7 * 6 * 9 * 3)) * (5 + 8)
  (6 + 7 * 8 + 3 + 9) * 5 * 3 * 3
  8 + 5 + 3 * 8 * 2 + 9
  (9 * (5 + 8 + 4 + 4 + 9 + 4) * 2 + 7 * (6 * 4 + 7 + 4 + 8)) + 9 + 5
  7 + 7 * 3 + 8 * ((5 + 3 * 4 * 9) * 4 * 6 * 8)
  2 + 5 + 8 + (3 + 3 * 4 + 8)
  5 + 4 * 3 * 9 + ((5 * 2 * 2) * 5 + 4 + (4 * 2))
  (3 * 7 + 2) * 3
  8 + 7 * 3 + 9 + 9 * 5
  2 + (2 * 6 + 3 + (4 + 3 + 4 + 2 + 7 * 9)) + 8 + 2 * 7 * 9
  (4 * 4 + 4 + 7) * 8 + 5
  2 * (5 * (8 * 5 * 2 * 6 * 8) * 6 * 3 + 8 * 8)
  8 + 3 * 9 + 5 * 9 * (5 * 7 + 6 + 5 * 3 + 7)
  2 + (3 + 7) * 8 * 9
  7 + 9 + 6
  4 + (6 + 7 + 8 + 4 * 5) + 2 * 6 * 8 * 2
  (9 + 3 * 7 * 3) * 8 + (9 + 5 * 2 * 8 + 3 * 4) * 8 + 2 * 6
  4 + 2 * 3 * 6 * ((3 * 5 * 3) * 2 + 6 + 5 * 4)
  3 + 9 * (7 * (8 + 7) * 9 * 7 + 7)
  6 + 3 * 6 + 4 * (7 * 7 + 2 * 9 + 4 + 2) + (6 + 6 + 6 + 7)
  (5 * 4) * 6 * (5 + 4 + 5 + 3 * (8 * 6) * 5) * 7
  (8 * 4) + 2 * (2 * 2 * (9 * 9 + 8) + 8 * 9 + 6) * 7
  3 + 5 * (4 + 9 + 3) + 9 * ((5 * 2 * 4 + 2) + 6 + 7 + 7 + 5)
  (3 + (9 + 7 * 3 + 6 * 6)) + 9 * 7 * 4 + (9 * 4 * 9 + 7 * 7 * 7)
  (3 * 7) + 4
  4 + (2 + 5 * 5) + 6 + 2 * 8 + 9
  8 * (6 * (7 * 3) + 8 + 5 + 8 + (8 * 2 + 6 + 9 + 4 * 3))
  7 * 8 * (5 * (2 * 6 * 9) * 3 * 6 * (6 + 8 + 5))
  (5 + 9 * 4 * 7 * 7 + 8) + 4
  ((8 * 9 * 6 + 4 * 6 + 3) * 6 + 2 * (4 + 8) * 7) + (4 + 4) + (7 + 7 * 7 * (4 + 5 + 2))
  (4 + (8 + 4 * 4 * 2 + 4) + 7 + 6 * 2 + 6) * 9 + 3 * 3 * (9 * 9)
  (9 * 5) * 6 + ((5 + 3) + 9 + 4)
  7 + ((8 * 9 * 5 + 2 + 2 * 7) * 9 * 3 * 7 * 2) * 6 * 2 + 6 * (9 * 4 * 2 + (2 + 7 + 5) * 4 + 7)
  8 + (5 * 4 * 7 + 2) * 6 + (6 * 3) * 4
  7 + (9 + 9 * 4 + 3) * 8 * (9 + (6 + 9 + 2 + 6 + 8) * 9 * 7 * (2 * 8 * 5 * 8)) * (4 + 7 + 8 + 2)
  6 + 3 * 5 * 8 + ((9 + 5 + 6) * 9 * (3 + 7 * 8 + 6) * (2 * 9 + 5)) + (7 + 6)
  5 * (9 + 7) + 2 * 4 + 3
  7 + 2 * (5 + 6 + 9 + 4 * (7 + 6 * 5 * 6 + 8) * 4) * 6 + 4
  6 * 8 * 8 * 4 + 8 * (2 * (9 * 4 * 7 + 9) * 5)
  8 * 4 + (4 + 7 * 3 + 5 + 8) + 4 * 7 + ((4 * 9 + 3 * 8) * 5 * 6 + 7 + 5 * 4)
  (9 * 3 + 3) + 9 + ((9 + 9 + 6 + 8) * 9 + 7) * 4 + 9 * ((8 + 7 + 8 * 2 + 8) * 5 + 8 + 3 * (3 + 5 + 2 * 6) * 8)
  (8 * 4 * 3) + (4 * 3 * 8 + (5 * 9) * 3 * 9) + 9 + (4 * 2) + 9 + 4
  2 * 5 + (9 + (5 + 2 * 2 * 7 * 3 * 9))
  (4 * 2 + 8) + (4 + 5 + 3 + 9 + (9 * 2 * 4) * (2 * 6 + 6))
  9 * 5 + 2 * ((3 * 3 + 3 + 9 * 7 + 9) * 3 * 3 + 6) + 7
  8 + (5 + 4 + 9) + ((6 + 9 * 6) * 8 + 3 + 6 * 8 * 6) * ((6 + 4) * (3 + 5 * 7 * 9 * 4 + 3) + (2 * 4)) + 4
  5 + 7 + (5 + 2 + (5 + 6 * 2) * 7) + 4 * 4
  (7 * 8 * 8) * (3 * (2 * 3 * 4 + 3 + 5 + 6) + 2 + 9 + (7 + 5 + 7 * 9) + 9) + 2
  2 * 8 + ((4 + 9 + 6 + 8 * 2 * 9) * 9 * (5 * 7 + 7 * 4 + 3 + 9) * 3 + 2 * 8) * (3 + 7 * 7 * 5 * 3) * 9 * 8
  6 * 8 * (5 * 7 + 3 * 2 * 3 + 4) + 7 + (4 + 2 * 8 * 8) + ((9 * 4 + 5) * (7 + 6 + 2 * 4))
  4 + 9 + (9 + 5 + 3 * 8) * 5 * 9
  5 + 2 + ((6 + 8 + 3 * 7) * 6 * 5 + 8 + 3 * 9) + 6
  3 + (7 + (3 * 6 * 7) + 9 * 2 * 2 + (9 * 4)) * 8 * 8 * 6
  ((5 + 9 + 8 + 6) + 4 + (3 * 4 + 8) * (2 + 5 * 2 * 6)) + 9 + 8 * 2
  8 * 9 + 2
  (3 * (3 * 5 * 2 * 3 * 3) * 3 * 7) * 3 * 6 * 5 + 8
  6 * 8 + 6 + 9 * 9 * (7 * 2 * 5 * 3 * 4)
  (8 + 2 + 4 * 3 + (4 * 6 + 7 * 2)) + (6 + 8) + 8
  9 + 4
  ((4 + 8) + 8 * (4 * 5 + 8 + 2 + 8 * 9) + 7 * (3 * 7 * 4 * 8) + 4) * 6 * 7 * (2 * 3 + 6)
  6 + (4 + 5 * 9) * 6 + 9
  (4 * (5 * 3 * 3 * 7 * 2 + 7) + 4) + 4 * 8 * 4 * 3
  8 + (9 + 8 + 9) * 5
  9 + 8 + 8 + 9 * (8 * 3)
  (6 + 8 * 5 * (5 + 7 * 8 + 4 + 3) * 6 + 4) * (2 * 4 * 5) + 7 * 8 + (9 * 6)
  ((9 + 2 + 5 + 5 + 5) + (9 + 6 + 7) * (7 + 5 + 9 + 5 * 6)) * (4 + (7 + 8 + 9 + 7) * (7 + 7 + 6 * 4 + 2) + 5 + 5 + 6) * 3
  7 * (6 + 8 * 6 * (5 + 6 + 5 * 4) * (9 * 2 * 6 * 4 + 8)) * 2 + 9
  (7 * 3 * 3 * 9 * 4) + 9 * (9 * (8 * 8 + 2 * 8 + 3 + 6) * 5 + (4 + 3 * 3 * 4 + 5 * 8)) + 9 * 4
  (7 + 3 + 4) * (2 + (4 * 7 + 5 * 3 * 7) * 4 * (3 + 5) + 8 * (8 * 9 + 2 + 7))
  7 + (8 * 7 * 2 + (7 + 6 + 8) * 3 * 3) + (4 * (7 * 7 + 4 * 7 * 4) + 7 * 5 + (6 * 6 + 4))
  5 + 3 + 4 + (2 + 4 * 5 * 2 + 8)
  (3 + 3 + 2 + 9 * 4) + 7 + (2 * 8 * 7 * 3 + (3 + 9 + 9 * 2)) + 7 * (9 * 7) + 2
  7 * 5 * (4 * 8 * 4 * (7 + 6 * 5 * 5 + 7 * 9)) + 9
  (2 * 5 + 2 + (6 + 3 + 4)) + 2 + ((6 + 8 + 9) * 4) + 2
  8 * (2 * 9 + 3 + 2) + 9
  (5 + (5 + 9 * 5 + 5 * 4 * 9) * 7) + (8 * 8 * 9 + 8 * 3) + 3 + 6 * 5
  8 + (7 * (7 + 9) + 3)
  (7 + 4 * 5 + (8 + 2 + 7) * 7 + 8) + (9 * (9 * 8 + 7 + 9 * 4)) * (4 + 9 * 5) + 2
  (7 * 5 + 6 * 5 * (2 * 4 * 3 + 4 + 7)) * 5 * (7 * (7 + 3 * 4) * 9)
  4 * (6 + 9 * 3) + 9 + 5 * 9
  (9 + 3 + 8 * 7 * 2 * 6) * (8 + 7 + 2 + 7 + 6) + (5 * 6 * 6 + 3) + 2 + 9
  5 * 5 + 6 + (3 * 7 + 6)
  2 * 7 + (5 * (5 + 8 + 9 * 5 + 3) + (3 * 2) * 3 + 7 * 2) + 2 + 5 + 3
  (9 * 5 * (4 * 9 * 6) + 3 * 8) * 2 * 8 * 4 * ((3 * 7 * 9 + 5) + 6 * 7)
  5 * 9 + ((9 + 4 + 5 * 7) + 4 * 3 + 6 * (2 * 2)) + 7 + 2
  2 * (6 + 2) * ((6 + 3 * 8 * 3 + 7) * 7 * 5 + (9 + 2 + 3 * 9 * 8 * 8) * 7) + (9 * 9 + 8 * 7 + (9 * 3 * 8) + 9)
  4 + 2 * 2 * 5 + 5 * 4
  4 + 4 * 8 * 7 + 3
  9 + (3 + 8 * (4 * 2 + 3 + 4) + 5 + 8) * 8 + 4 * 7 * (4 + 3 * 9 * 4)
  6 + (3 + 2 * 2 + 3 + 7 * 8) * 4 * 8 + 2
  (7 * 4 + 5 * 2 * (5 * 8 + 3 + 6 * 2)) + (7 * 2) * 9 * (7 + 7 + 5) * (8 + 2 * 9) * 6
  3 + 9 * (7 + 2 * 6 + 2 + 5) * 4
  9 + (8 * 2 + (8 + 8 + 8)) * 8 + 7
  6 * 7 * 7 * 4
  2 * (2 + (4 + 7 * 7 + 6 * 8)) + 4
  (9 + 8 + (5 + 9 + 4 + 9 * 6 + 3)) * 4
  6 + ((9 + 2 + 4 * 2 + 9) * 7 * 3 + 5 * 6 * 6) + 7 * 4 * 9
  (5 + (5 + 9 + 6 + 7)) * 5 + 4 + 6
  (5 * 8 * 6 + 8 + 7) * 2 + 7 + 8 + 5
  (7 * 5 + 6 + 5 * (6 + 4 * 9)) + 4 * 8 + 3
  4 * 9 * (8 * 5 * 3 * 9 * (4 * 2 * 3 * 4) * 3) * ((5 + 5 * 3) + 9 * 4 + (9 + 6) * 2 * 6)
  7 * ((3 + 8 * 3 * 6 * 8 * 6) + (8 * 3 * 7 + 4 * 2 * 5) * 5 * 8 * (7 + 9 * 4 + 6 * 2 * 8) * 4) + 7 + 7
  (7 * (7 + 4 + 4 * 5 * 5 + 2) * 2 + 5 * 6) + 2 * 5 + 3
  (2 + 7) * 6 * (3 + (9 * 4 + 8 * 8 + 5) + (2 + 7 * 5 * 5) * 5 + (3 * 6) * 2) + (8 + 5 * 9 + 9 + 8)
  6 * 5 * 3 + 7 * (4 + 2 + 7 + 9)
  (9 + 6 + (2 + 2 + 5 + 6 * 5 * 8)) + 7 * 2
  ((7 + 7 * 9) + (5 + 8 * 2 + 4 * 7 * 7) + 7 * 3) * (8 + 2) + 7
  3 + 6 * (8 * 7 * 3 * 7) + (5 + 4 * 3) + 6
  (7 + 7) + 9 + (5 * (5 + 9 * 3 + 3 * 2 + 3) * 4 * (8 * 7) * 2 * (7 * 9 * 6 + 4 * 6 + 2)) + 6
  ((6 * 2 * 6 + 3 * 8) + (9 + 8 + 9 + 2 + 3 + 5) * 7 * 9) + 8 + 6 * 3 + 6 + 7
  ((9 * 8 + 2) * 8 + 6 * (4 + 8 + 4 * 9 * 9) + 9) * 7 * 5
  8 + 6 + (4 + (6 + 6 + 3 + 9)) + 3 * 3 + 4
  8 + (3 + 5 * 9 + 3) * 9 + 8 + 8 + 7
  5 + (3 * 4 + (8 + 4 + 3 + 4) * 8 + 7)
  (9 * (6 * 7 * 9 + 5) * (7 + 6 * 2 * 4 + 2 + 8)) + 4 + 8
  3 * 5 * (9 + 2 * (9 * 6 * 7 + 9) + 2) * 7 + 6
  4 + (3 * 7)
  (3 + 7 * 9) * 6 * 6 + 3 + 3 + (2 + (9 * 4 + 6 + 7 + 9))
  (4 * 2 * (9 + 5 * 9 + 5 * 9 + 9) + 8) * 7 * 7 * 3 + (9 * 4 * 7)
  3 * (4 * 5) + 9 + ((7 * 8 * 4 * 8) + 7 * 3)
  8 + 9 + 4 * 4 * ((5 + 3 * 6) * 5) * 6
  9 + 8 + (9 * 9 + 4) + 5
  9 * ((8 * 7 * 2 + 4 + 8) * 9 + 5 + (2 + 4 + 6 * 7 * 3 + 3) + (7 * 9 * 2 + 7 * 9) * 3) + (6 * (9 + 7 * 3 + 7 + 6) + 6 + 6)
  4 * ((5 * 7 * 9 * 9 * 4) + 2 + (5 * 5) * 2 * 7 * 9) * 4 * 7 + 2 + 9
  9 + (7 + 5) + 4 + (3 + 2 * 4) * 8 + 7
  (8 * 7 + 7 + 6 + 7 * 6) * 5 + 7 + 5
  9 + 5 * 5 * (3 + 8 * 4 + 8 + 4 * 5) + 3 + 4
  4 * 5 * 6 + 6 + (7 * 7 + 6 + 4 + 4 * 3)
  (4 * 2 + 9) + (4 * 3 + 9) * 9 + 8
  9 * 6 + ((7 + 2 + 6) + (2 + 2 * 3)) * 6
  ((5 + 6 + 6 * 6) + 8 + (3 * 8 * 3 * 3 * 6) * 4) * (2 + 6 * 7 * 6 * (3 + 5)) * 6 * 2
  (5 * 3) + (2 + 9) + 2
  4 * 8 + 6
  2 + 6 + 2 + 4
  5 * 8 * (9 * 2 * 8) + 8 + 9 + 4
  5 * (4 + 7 * (5 * 7 * 6 * 5 * 8) + (6 + 3 * 6 + 3 + 7) * (7 * 5 * 8) * 4)
  5 * 9 + 5 + 5 * (9 * 2 + 4 + 9 * 7 + 4) * 7
  5 + (4 + 7 + (9 + 7 + 5) + 9 * 6 * 3) * 5 * 3 + 5
  (7 * 8 * 7) + 6 + 2 * 4 + 6 + 4
  (2 + (4 + 4 + 6) + 9 + 5 * (2 * 7 * 3) * 2) + 7 + ((9 + 5) + (5 * 3 + 8 + 5 * 6 + 8) * 2 + 7) * 9 * 3
  (2 * (9 * 5) + 9 + 4) + 8 * 7 + 8 + 8 + 8
  2 + 3 * 3 + (7 * (9 + 5) + (8 + 6 + 8 + 2) * 7 * 2 + (7 + 8 + 9 * 3)) + 5 + 9
  5 + 4 + (7 * 3 * (2 * 6 + 2 * 3) + 3) * 7
  9 + (3 + (7 * 3) + 7 + 3) + 5 + (7 * 3 + 7)
  7 + 4 * 2 + 4 + (8 + 3 + 7) * 5
  9 * 4 + (6 + 9 + 3)
  2 + 3 + 7 + 7 + 8 + ((4 * 7 + 9 + 8 + 6 * 3) * 3 * 7 * 9 * 6 * 7)
  (7 * 9 * 4 * 2) + (3 + (5 + 2 + 4 + 6 * 8 * 5) + 5 + 5) + 4
  (6 * 2 + 8 + 3) * ((8 * 3 * 9 * 2 * 7) * 4 + 9 + 6) * 6 + 8
  (3 * 5 * 8 * 4 + (2 * 5) + 2) + 5 * 8 * 5 * 3 * 4
  4 + (9 + 7 * (6 + 3 + 9 + 7))
  (9 * 2 + 3 + 7) * 7 * (8 * 2 * 7) * 3 + 4
  6 * (8 * 2 * 8 + 8 + 9 + (4 * 7 * 4 + 9 + 5 + 5)) * 4
  5 * (2 * 4) * 6
  4 + 4 * 4 + 6 + ((5 + 3 + 2 * 2 + 4) * 3 * 2 * (7 + 7 + 8 + 6 + 3) + 2)
  7 * 6 * 8 + (9 * 9 + 3 + 9 + 3 * 7) + 3
  8 * (2 * 8 * (2 * 4 + 6 * 4 * 8)) * (9 * 3 + 8 * 3 * 8 * (6 + 8 + 6 * 7 + 3)) * 6 * 6 * 7
  (3 + (7 * 7 * 4 + 2 * 4 + 7)) + 6 * 5 + 9 + 4
  (2 + 2 * 3) + (3 * 3 + 2 * 3 * 5 + 9) + 8 * (8 + 9 * (9 + 9 + 2 * 9) + 9 + (6 + 3 * 7) * 5) * 4
  4 * 8 * 3 + 5 * (6 * (5 * 4 + 9 + 4 * 6 * 3)) + 9
  4 * ((7 + 4 * 8) * 2) * 6 * 6 + (7 * 4) * 2
  7 + 3 * 3 + 6
  2 * (3 + (6 + 5) + (3 * 2 + 6) + 4) + 6 * 7 * (5 * 2 + 3 + 9 * 9)
  (7 * 5 * (8 + 4 + 6 + 3 * 5)) + 2 * 9
  3 * 4 * 7 * (7 + (5 + 7)) + 9
  9 * ((9 + 7) + 3)
  5 + ((6 * 2 + 8) * 7)
  7 + 2 * (6 + 6 + 6 + 3) * 2
  (9 + 3 * (4 + 8 + 5 + 5 + 7 * 9)) * ((7 * 5 * 9 * 2 * 6) + 5 + 6 * 3 + 9) * 8 + 7 * (5 + 6)
  3 + ((8 + 5 * 9 * 6 + 6) * 2 * 5 + 9 * (8 * 5 * 8 + 4 + 9 * 2) * 8) * 8 + 7
  3 + 2 * 3 + ((9 + 3) * 9)
  8 * (2 + (6 + 2 * 8)) * 2 * 7 + 8
  (4 + 9 * (8 + 4 * 2 + 9 * 6 + 8) * 6 + (8 * 9) + 3) * 6
  3 * (3 * 9 * 5 * 6 + 2 + 4) * 9 * (8 * 4 * 7 * 7 * 8 * 3) * 5
  8 * 2 + (7 + 4 + 5 + 2 * 4 * 8) + (2 * 4 + 7 + 4) * 4
  (7 + 9 + 9 * 6) * 6 + 9 * 8
  5 + 8 + 3 * 3 + (6 + (9 * 8 + 8 * 5) + (3 * 5 * 4))
  (2 * 9 + 3 * 9 * 8) * (6 + 6 * 8) * 9 * 4
  2 * (8 * 9 + 4)
  2 + 5 * (5 * 4 * 9 * 3 * 8 * 3) * 8
  9 * 9 * (7 * 3 * 4) * 8 * 8
  4 + (9 + 8 + 6 + 8) * 9 + (6 + 6 + 9 * 5) + 9 + (8 + 4 * (5 + 6 + 3 * 6 + 6) + 5 * (8 * 3 * 8 * 3 * 7 * 2))
  (5 * 8 * 7 * 5 * 2 + 4) * 6 + 8 * ((4 * 5 * 9 + 3 + 2) * 4 + 4)
  2 * 9 + (9 * 7 * 8 * 6 * 7 * 9) + 4 + (7 + (3 + 9 * 5 + 5 + 2) + (6 * 3 + 2 + 7 * 7)) * 7
  7 * 3 * 7 + (2 * 2 * 4) * 5
  (2 + 8 * 6 * 3 * (6 * 5 + 9 + 3 * 8 + 8) + (3 + 6 + 9 + 4 * 3 * 6)) + 4 + (3 * 4 + 2 + (5 * 3 * 7 + 5 + 9)) * ((2 + 3 + 6 + 6 * 4 * 2) + 4 * 9 + 7 * (6 + 2)) + (6 + (8 * 7 + 2 + 2) + 8) * (7 + (9 + 3 + 9))
  7 * 4 * (7 + 5 * 3 + (9 * 2 * 5 * 2) + 9 + 2) * (5 + 4 * (8 + 8 * 6) + 2)
  7 * (9 * (7 + 4 * 3 + 8 + 3) + 8 + 8) + 9
  ((8 + 6 + 4) + (9 + 6 * 7 + 4) + (6 + 9 * 5 + 4) + 3) + 5 + 9 + (9 + 9 * 7) * 3 + 2
  8 + 4 + (9 * 9 + 2 * 8 + 4) + 5 * 4
  9 * (3 * (7 + 8 * 7 * 8) + 5 * 7 * (5 * 6 + 7 * 4 + 6)) + 9
  (8 + 3 * 6 + 8 + 8) * 5 * 3
  ((5 + 7 + 4 + 2 * 6) + (6 * 4 + 6 * 3 + 8) * 5 + 8 + 8 * (5 * 6 + 2)) + (2 * 4 + 9 * 7 * 6 * (9 * 6 + 6 * 8)) * 6 * 9 + 3
  ((3 * 7) * 7 * 2 * 8 + 2 * (2 + 6 + 7 * 3 + 2 * 5)) + (6 * 3) * (4 * 5 + 4 + 9) + 6 + (9 * 7 * 5 + 5 * 9) + 3
  6 + 4 + 5 + 8 + (2 * (2 * 2 + 8) * 9 * 7) + 6
  5 * 4 + 7 + (2 + 6 * 3 + (2 * 2 * 2))
  (8 * 3 + (2 * 6 * 4 * 2 * 9) + 5 + 9 * 8) * 5 + 7
  5 * 5 * (3 + 4 * 3) + ((7 * 3 + 5 + 8) + 2) + 6
  ((9 * 5) + 4) * 6 + 4 * ((5 * 2 * 9) + 5 + 3) + 6 + 3
  9 * (2 + 4 + 4 * (5 * 9 * 3 + 4 * 5 + 8) * (3 * 7) * 6) + 6
  (3 + 2 * 5 + 5) * (9 * 3 + 5 * 7) + 3 * 9
  6 * ((7 + 9 + 9 * 2 * 4) * 8) + 9 + (6 * 8) * 3
  4 + 9 * 8 + (3 * 4 * 4 * 7 * (7 * 6 + 7 * 2 * 7 + 8))
  (7 + 7) + 7 * ((7 + 3 * 4 + 8 + 8 + 6) + 9 + 2 * 6 * 9 + 2)
  9 + 7 + 9 * 9 + 5 + 5
  8 + 2 + 3 + 2 + 6
  8 + (4 * (3 * 2) + 9 + 6)
  9 * 8 * 2 + (6 + 6 * 9) + 8 * 8
  7 + 4 * 2 * 5 + 4 + 7
  5 * 8 * 4 * 2 * (6 + 5 + 2)
  3 * 5 * (3 + 9 * 2) * (3 + 5)
  3 * 6 * 7 * (2 + 8 * (8 + 6 + 3 * 9 + 9 * 5) * 3 + 8 * (6 + 3 * 5 + 8 + 4)) * (9 + 5 + 9) * 5
  5 * 3
  (3 * 2 + 7 + 8) * 2
  ((7 + 8 * 4) + 6 + 9 * (6 + 5 + 4 + 9 + 7 * 6) + 9 + 5) + 5 * 4
  3 * 9 * 6 * 4 * ((4 + 8 + 6) + (4 + 6 * 2 + 9 + 7 * 2) + 7 * 2 + (5 + 9 * 2 + 2))
  (8 + 6 + 5) * 4 * 8
  7 * 4 + 8 + (4 + 9 + 3 + 5 + 2 + 6) + (7 + 2 * 5 * 7) + 3
  6 + 7 * 4 + (8 * 3 + 8 * (7 * 5 + 3 * 4 * 5) * 9 * 2) * 2 + 3
  9 + 6 * (6 * (4 + 7)) + 4 * 7
  (4 + 5 + 8) * 8 + (8 * 7 + 9 * 5) * 9
  5 * 3 * (9 * 9 + 5 * 6 * 7 + 6) + 7 + 5
  7 + 3 * 7 * ((4 + 6 * 4 + 8) + 5 + 9 + 7 * 5)
  9 * (8 + 9 + 3 * 5 * (4 + 6 * 4 + 7 + 7) + (3 * 8 * 2 + 2 + 8)) * 8 + 8 * (3 * 9) + 6
  2 + (7 * 9 * 3) + (5 * 8 * 8 * (2 + 2 * 2 * 6 * 7 + 6)) * (9 * 6 + 2)
  8 * (4 + (8 + 3 * 4 * 2 * 5) * (8 * 2) + (2 + 2 + 7 * 6 + 3 * 2) + (2 * 7)) * 7
  3 * 7 * 3 + 4 + (2 + (6 * 3) + (2 + 5 + 8 + 3 * 3) + 2 + (4 * 7 + 2)) * 4
  (7 * 7 + 5 * 8) * 3 + 2 + 9 + (2 * 8 * 4 + 7 + 4 * 8) * (7 + 8 * 5)
  2 * ((7 * 5 * 4) * (4 + 7 * 9 + 3 + 3 * 6) + (5 + 5 + 9 + 7 * 8) * (5 * 2 * 8 + 4 * 6) + (5 + 4 + 3 * 9 * 4 + 9)) * 6 * 8 * 5
  (8 + 2) * 3 + 8 * ((5 * 9 * 7 * 7) * 7) + 8
  (3 + 7 + 9 + 7) * 9 * (2 + 5 * 8 * 9) + (5 * 5 + 9) + (5 + (4 * 4 * 7 * 9) + 4 * (2 * 2 + 5))
  3 + (4 + 8 + 3 + 4 + 7 + 6) + 4 + 3 * 4 + ((5 * 6) + 2 * 5 * 2 + 8 * 3)`

  p1 := P1(input)
  p2 := P2(input)

  fmt.Printf("P1: %d, P2: %d\n", p1, p2)
}

func P1 (input string) int {
  ret := 0

  for _, line := range strings.Split(input, "\n") {
    v, _ := P1Calc(line, 0)
    ret += v
  }

  return ret
}

func P1Calc(input string, offset int) (int, int) {
  total   := 0
  current := 0
  op      := "+"

  for i := offset; i < len(input); i++ {
    switch input[i] {
    case ' ':
      if (op == "" || current == 0) {
        continue
      }
      total = Calc(total, op, current)
      current = 0
      op = ""

    case '+':
      op = "+"

    case '*':
      op = "*"

    case '(':
      current, i = P1Calc(input, i + 1)

    case ')':
      total = Calc(total, op, current)
      return total, i

    default:
      // Number
      v, _ := strconv.Atoi(string(input[i]))
      current *= 10
      current += v
    }
  }

  total = Calc(total, op, current)

  return total, 0
}

func P2 (input string) int {
  ret := 0

  tokenRegex := regexp.MustCompile("(?:[0-9]+|[()+*])")

  for _, line := range strings.Split(input, "\n") {
    tokens := tokenRegex.FindAllString(line, -1)
    v, _ := P2Calc(tokens, 0)
    ret += v
  }

  return ret
}

func P2Calc(input []string, offset int) (int, int) {
  stack   := []int {}
  op      := ""
  prevOp  := ""
  total   := 1
  end     := false
  endIdx  := len(input)

  for i := offset; i < len(input) && !end; i++ {
    switch input[i] {
    case "(":
      v, end := P2Calc(input, i + 1)
      stack  = append(stack, v)
      i      = end
      prevOp = op

    case ")":
      end    = true
      endIdx = i
      break

    case "+":
      op     = "+"

    case "*":
      op     = "*"

    default:
      prevOp = op
      v, _ := strconv.Atoi(input[i])
      stack = append(stack, v)
    }

    if prevOp == "+" && len(stack) > 1 {
      // Pop the last 2 elements off the stack, sum them, and put them back
      a := stack[len(stack) - 1]
      b := stack[len(stack) - 2]
      stack = stack[:len(stack) - 2]
      stack = append(stack, a + b)
      prevOp = ""
    }

  }

  // Since we do addition at the time, we only have multiplication left
  for _, v := range stack {
    total *= v
  }

  return total, endIdx
}

func Calc(a int, op string, b int) int {
  ret := 0
  if op == "+" {
    ret = a + b
  } else if a == 0 {
    ret = b
  } else {
    ret = a * b
  }
  return ret
}
