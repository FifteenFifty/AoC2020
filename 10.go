package main

import (
    "fmt"
    "sort"
)

func main() {
    input := []int{
118,
14,
98,
154,
71,
127,
38,
50,
36,
132,
66,
121,
65,
26,
119,
46,
2,
140,
95,
133,
15,
40,
32,
137,
45,
155,
156,
97,
145,
44,
153,
96,
104,
58,
149,
75,
72,
57,
76,
56,
143,
11,
138,
37,
9,
82,
62,
17,
88,
33,
5,
10,
134,
114,
23,
111,
81,
21,
103,
126,
18,
8,
43,
108,
120,
16,
146,
110,
144,
124,
67,
79,
59,
89,
87,
131,
80,
139,
31,
115,
107,
53,
68,
130,
101,
22,
125,
83,
92,
30,
39,
102,
47,
109,
152,
1,
29,
86,
    }

//    input = []int {
//28,
//33,
//18,
//42,
//31,
//14,
//46,
//20,
//48,
//47,
//24,
//23,
//49,
//45,
//19,
//38,
//39,
//11,
//1,
//32,
//25,
//35,
//8,
//17,
//7,
//9,
//4,
//2,
//34,
//10,
//3,
//    }

//    input = []int {
//16,
//10,
//15,
//5,
//1,
//11,
//7,
//19,
//6,
//12,
//4,
//    }

    sort.Ints(input)

    p1 := 0
    p2 := 0

    d1 := 0
    d2 := 0
    d3 := 1
    prevJ := 0

    for _, jolts := range input {
      if jolts - prevJ == 1 {
        d1++
      } else if jolts - prevJ == 2 {
        d2++
      } else if jolts - prevJ == 3 {
        d3++
      }
      prevJ = jolts
    }

    p1 = d1 * d3
    memo := make(map[int]int)
    p2 = combosInRange(input, len(input) - 1, memo)

    fmt.Printf("P1: %d, P2: %d\n", p1, p2)
}

func combosInRange(input []int, start int, memo map[int]int) int {
  if m, ok := memo[start]; ok {
    return m
  }
  combos := 0
  if (input[start] <= 3) {
    combos++
  }

  if start > 0 && input[start] - input[start - 1] <= 3 {
    combos += combosInRange(input, start - 1, memo)
  }
  if start > 1 && input[start] - input[start - 2] <= 3 {
    combos += combosInRange(input, start - 2, memo)
  }
  if start > 2 && input[start] - input[start - 3] <= 3 {
    combos += combosInRange(input, start - 3, memo)
  }

  memo[start] = combos
  return combos
}
