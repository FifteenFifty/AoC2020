package main

import (
  "fmt"
  "strings"
  "regexp"
  "strconv"
)

func main() {
  input1 := `0: 8 11
1: 41 41 | 48 48
2: 48 85 | 41 128
3: 48 118 | 41 25
4: 131 48 | 70 41
5: 41 110
6: 48 100 | 41 132
7: 41 41
8: 42
9: 89 41 | 100 48
10: 108 48 | 53 41
11: 42 31
12: 41 18 | 48 98
13: 43 41 | 69 48
14: 9 48 | 93 41
15: 41 20 | 48 63
16: 48 5 | 41 45
17: 62 48 | 26 41
18: 48 48
19: 48 72 | 41 18
20: 61 48 | 57 41
21: 48 83 | 41 51
22: 103 48 | 18 41
23: 84 41 | 27 48
24: 48 109 | 41 38
25: 48 81 | 41 7
26: 41 72 | 48 81
27: 100 41 | 7 48
28: 41 132 | 48 120
29: 116 48 | 3 41
30: 89 41 | 18 48
31: 48 133 | 41 127
32: 41 96 | 48 95
33: 120 41
34: 48 41 | 48 48
35: 123 48 | 91 41
36: 48 81 | 41 100
37: 41 7 | 48 120
38: 2 48 | 77 41
39: 41 67 | 48 41
40: 81 48 | 7 41
41: "a"
42: 41 68 | 48 105
43: 52 48 | 112 41
44: 101 48 | 114 41
45: 7 48
46: 48 29 | 41 82
47: 120 41 | 81 48
48: "b"
49: 81 41 | 7 48
50: 7 41 | 7 48
51: 41 1 | 48 81
52: 41 73 | 48 104
53: 54 48 | 119 41
54: 41 36 | 48 55
55: 48 1 | 41 18
56: 34 48 | 39 41
57: 41 92 | 48 80
58: 41 22 | 48 111
59: 41 47 | 48 125
60: 30 41 | 86 48
61: 41 37 | 48 6
62: 100 48 | 98 41
63: 41 59 | 48 17
64: 48 44 | 41 74
65: 16 48 | 58 41
66: 41 126 | 48 121
67: 48 | 41
68: 64 41 | 24 48
69: 106 41 | 32 48
70: 87 48 | 115 41
71: 67 34
72: 48 48 | 41 67
73: 1 48 | 103 41
74: 41 107 | 48 124
75: 89 41 | 39 48
76: 48 14 | 41 35
77: 40 48 | 47 41
78: 41 65 | 48 66
79: 67 1
80: 120 48 | 98 41
81: 48 48 | 67 41
82: 23 48 | 21 41
83: 67 89
84: 18 48
85: 120 41 | 89 48
86: 48 89 | 41 39
87: 103 48 | 72 41
88: 7 48 | 89 41
89: 41 41 | 41 48
90: 81 41 | 34 48
91: 34 48
92: 48 103 | 41 34
93: 110 41 | 34 48
94: 41 4 | 48 76
95: 34 41 | 103 48
96: 100 48 | 103 41
97: 41 120 | 48 34
98: 41 48
99: 97 48 | 92 41
100: 48 41 | 67 48
101: 129 48 | 93 41
102: 48 62 | 41 50
103: 67 67
104: 48 1 | 41 81
105: 46 41 | 78 48
106: 88 41
107: 79 41 | 33 48
108: 102 41 | 60 48
109: 48 113 | 41 99
110: 48 41 | 41 41
111: 103 67
112: 48 71 | 41 123
113: 49 41 | 56 48
114: 41 117 | 48 73
115: 67 132
116: 41 93 | 48 122
117: 48 120 | 41 39
118: 89 41 | 132 48
119: 48 12 | 41 19
120: 48 41
121: 48 49 | 41 47
122: 48 120 | 41 89
123: 110 48 | 98 41
124: 83 48 | 130 41
125: 41 120 | 48 132
126: 75 48 | 27 41
127: 41 94 | 48 10
128: 48 72 | 41 1
129: 48 110 | 41 120
130: 103 48 | 89 41
131: 90 41 | 28 48
132: 41 48 | 48 41
133: 48 15 | 41 13

bbaaabbbbaabbaababaababbbabbbaaa
bbbbbaabbabaaaaabbaaaabbabaababbbbabbbbababaabaabaaaaaabbbaaaaab
baabbaabababbabbaaaaabba
ababbbbbbbabaaaaabaabbabbbbaaaabbabbaaba
abbbbabaabbbabbbabbaaaaaaaaaabba
bbaaaaaabaabbbababbaabbabbabaaaabbbaababaaaabbbaabababaa
bbbbbaabbabbbbabaabbbbaaaaaaabba
bbbaabbaaababbbbabbbaabb
abbbbaaabbbaabababaababa
aaaababbabaaabaaabbbbbaaaabaaabaaababbbbbaabaabababbbbbbbaabaaaa
babaababaabbbbabbabbbbabbababbaa
babbaaabababbaaabbbaabab
aaaabaabbbababbbbbbbbbbb
aabbbbbabbbbabbbbbbbbababbbabababaaaababbbaaaabaaaabaaaaabbaabbaaababaaaabaaabab
aaabbaabaaaabbbababbbbbbbabaaabaababaabb
abbaaaaabbaabaabbbabbbab
baabbbaabbabbabaabbabbbb
abababbbaaaaabbbbbbbbabaabaababbabbbbbbbbbabbabbaaabaaba
abbaaaaaaabbabbaaaababab
baabababbaaaaaaabbbbabab
babbabaaaaabaabbbbbaaaab
aabaaaababbaaabbbbbbaaaa
babaaaaabaaaababbbabbabb
abbabaaaabbababbaaabbaba
aaabaaaaaabaaaababaaaabaaaaabaabaababbabaabababa
ababbabbaaabaabababaabbaaabbabababaabbaa
baaaabbbbbabaaaabbbabaabbabbbaaa
abaabaaaababababbaabaabb
aabaabbbbabaaabbaababbaabbbbbbabbbbaabababbabbaa
abaaaaaabbaaaabaaabbaaab
aaaabbbbbabaabaaababaaabbbbbaaabbbbbabbbaaaabaaabababaaa
bbabbbbaaababbbbaaababbbbbbaabbaababbabbaaaabbbaaaaabbbaaabbabbbaaababab
babaaaaaaabaaabaabbbabab
aaaabbaaabaaabaabaabbbabaaaabbbbabbaaabbbbabbbbbbbaabbabbbbbababababaabb
abbbabbabbabbbbaaabaaaaababbaabbaabaaabb
abbaaabbbbaaababbbbaaabbabbbbbbbbbbbbaababbbbbbbabbbbabbabbabbbb
babbbaababbaabaaaabbbbbb
baaaabbbbbbbaabbbabbbababaaaabba
aaabbbaababbaaaabaaaabbbbaaabaaa
abbbbaaaaababbabbbabaaabbbaaabba
babbabaabaaaabbbbbbaaaab
abbababaabbbbaababaaabba
aababbbbbababaababbaababababbbbbbaaabbbbaaabaaab
baabbaabaaabaabaababbaaaabaaaabaaaababbbbbbabbbbbbbbabab
bbababbaaaabaaaaabbabbaa
bbababbbbababbabbbaababbbabaabbaabbaaabaabababaa
baababbabababbabaababbab
abbababbbaababababaaabaaababbbabbbbaaaab
abbbbbaaaabaaaaaabababbb
abbbbaaabaaabbbaababbbaa
bbababbbbabaabaabbbbabaaaaaaaabb
aaaabbaaaabbbbaaaaababba
abaaabaaaababbabbaabaaba
babbabababbaabaaabaaabbb
babababbbaaaabbbbabaababbbbaaaaaaabbaababbaababbaaabaaababaaaaabbbabbbab
baaabbabbaababbaaabbbaaa
bbbababbbabababbaaaabaabababaabaaabbabbb
bababbabbabbbbabaababbbbabaabbabaababaabbbaaaabaaaabbaabaaabbaaa
bbabbbaabaaaabbbbbabaaabababaababaaababa
aabbbbaabaaaaaaaaabbbaabbbbabaababbaaabbbabbabbbaaabababbaabbaba
bbbbaabbbabbabaaabbbbbaaabaaabab
baabbabababaabbbabbabbbaabaabaabaabbbbaabbabbabbabbbaaaabaaaaabaabbbaaba
ababbabbbbbaaabbbbbbaaaa
aaaaaabbababbbabbaababbb
ababbabbaaabaababaababaabaababaaababbababaaaabaa
baaaababaabaaaaabaabaaba
bbabbbbaabbbaabbbabbaaabbababaababbbbbabbbabbbbb
bbbababbbaabbbaaabbbbabb
bbbababbbbbbbaababbbaabbbababbbabbabbabb
abaaaaaaaaababbbaaaabaaa
baaaababbbababbaaaabaaab
babaaabbbbbabbaaaaababbbabbbabaaabbabbaa
bbaaaabbbaaabaaaabababbbbaabaaabbbbbbbbabaaaaabbababaaaaaaaabbabbbaaabaababbbaab
aabaaaabbbabbbaaaaaaabaa
ababaaabaabbbbabbbbbaabbaabbabbababbbbbbaabbbbbb
bababaabbbbabaabbababbba
bbbbbbabababbabbbbaaaaaabaabaaaaaababaaabbbbbbaabbbabbbaaaaaabaa
abbbbaaaababaaabaaabaabaabbbbbbbbabbbaaa
abaabaabbabbaabbbabaabaabbbabbbbabbbaaabaabbabaaabbabbaabbaaabba
abbbbbaaababbaaaabbaabaabbabbaaabbababbbaaaaabaaababaabb
abbaaababbbabbbbabbbbababaabababbbbbbbabababbababbbbabbb
aababaaaaabaaaababbbaaabbbabbaabababbabababbabbb
bbbababaaabababbabaabaaabbabbbab
bbabababbbaabbababbababaabbaaaaaaaabbaba
aaabaabbbbaaabbbabababbb
abababababbbaaaaabbababbaabaaaaabbaaababaabbbababbabbbabaaaaabaa
bbbbabaaabaabaabbbaaaabaaaababbbbbaababababababa
aaaabbaaabaababbbaaabbbb
baaabbabbabbababaaaabbabaabaaaabbbbaabbbabaaaabb
aabaaaaaaaababbbabbbaabaabbbbaabbbbbbbabaababababaababaa
babaabaaaaaaabababbaaabbbbbaabaabbababab
aababbbbbbbabaaaaaabaabbbbbabbaabbabbababbbbaaba
bbbabbaababbbaabaabbaaaa
abbbaabbbbbabaaaaaababab
babbababbbbabbbbaabaaabaabbaabbaabbabaab
abbaaaabbbabbbaaabbbbbab
aabbabbbbabababbaaabbbbaaababaabbabbbbbbbbaabbbaabbbabaababbbaabbbaabbaabbbaabbb
bbabaabaabaabaaaabbababaaabaabaa
bbbbbaabaaaaaabbaabbbaababaabaaabbabaaabbbabbaaaabaaaabb
aababbababaaaaaaabbbbbbbbaabaaab
aaabbbaabbbaabbababaabbabbbbabaabbaabbaaaabababbabbbbbba
bbbbbababbabbaaabbbabaabaabbbbbb
aababaabaaaaaaabaabbaabaaaabbaaaaababbbbabbbabbbbaaabbab
baabbaaabbaaaabbaaabbbbbabaabbabbbbabbba
bababaabbabbbbbabaabbbbabbbaaaab
ababbaaabbbaabbababbbbab
babaaaababbbbbbbaaaaaababaabbbba
aaaabbabbaaabbaabaabbaba
bbbbaabbbbbbbbababababbabababbaa
baaabbaaabbaaaaababababa
aaabaaaabbbbbaababbbbababbabbbaaaabbabbaababaabbbaaabaabababaababbabaabb
babaabbababaaabbbababbbb
baabababbbaaaabbbbabbaaaabbabbaabbabbbbb
babbbbababaaaabaabbbbbaaababbaabaabbbbbabbaabaaaaabababb
babbbaaabaabaaabaabaaabb
aaabbbbbbbaaabbbbaaaaaba
baaaaaabaabbbbbbabaababa
ababaaabbbbaaabbababaaaaababbaababbbbaaaaababaaabbabbbaababababbabaabaabbabbabbbbaabaabb
abaaaaaaabbaababbbaababbbbbaabbbbbaabbaa
ababbbabbabbbaabbbbbbbbb
abbaaaabbaabbbababaaaabb
bbababbbababbbababbabbbb
baabaaaababbabaaababaaabbbbaaabbbbbbbbbaaabbaaaaaaabbbba
baaabbabbaaaaaaababaaaabbaabbbba
abaaabbbbaabbaababababbaaabbbaaabaababab
bbbaabbaaabbabaababbababaabbabbbbabbaabaaaabbaabaabaabaabaaaabbb
aabbaabababaababaaababbbbabaabbb
abbaabbabaaabbabbaabaaab
bbbbabaababbbabaabaabbbb
baabaaaaababbbabbbbaaaaabbbabababaababaa
baaaababbbaaababbbabbaabaaabbaaa
babbabaabbaabbabbbabbabaabaaabab
aababaabbabaabbabbbabbba
aaaaaaaaabbaabaaaabaabbbabbabbab
babbbaabbbabbaaabbabbbbabaaabbbb
abbaaaaaabbbabbababbbbbb
babbabbabaababbaabbbabbbaaababaa
bbabbbaaaabaabbbaababaabaababaabaaaaabba
babbabaaaaabbbbbbbbabbba
baaaaaaabbbbbababbbbbbba
bbbbabaaaaaaabbbababbbba
aababbabbabaababbaaaabbbbaaabaaa
babaabbabbaaaabbbaabaaba
abbbaaabababababbbabbbbb
bbbbaabbababbabbababababbbbaaababbbaaaba
aabaaabaabbabababbaaaabbababaabbbbaabbba
bbbaabbabababbabbbbbbbbb
abbababbbbbbbaabbaaabbababbbbaabaabbaaaaabbbbbab
ababbabbbbaaabbbababbabbaaaaaaba
bbabaabababbbbabbababaabaabaabbaaabaabab
bbaaabbbbabaabbaaabbbaba
baaaabbbbaabbaaaabaaabaabaababaabbbaabaa
aababaababbaaaaaabaaaabb
bbbbabaababaababaaaaababbbbbabaababaababbbbaaaab
abababbbbababbbbbaaaabbbaabaaaaaababbabababbbbbbbbaabbba
bbaaaaaaabbbbaaabaaaaabb
bbabaaaaabbabaaababaaaabaaabbaab
bbbbbababaabbaaaaaaaaaaaabaaaabaabaabaabbaabaaab
baabbbabaaaaaaaabbbabaaaababbabbaaababab
aabbaabaabbabaaaaabbbbaabaaabbababbabbba
aaaabaababaaabaabbaaaaab
abaabbabababbbbbabbbabbabaabbaababababaa
bbbbaabbbbbbaabbbaabbbaaaabbbaabbaabbbababbbabababbaabbb
bbaabbaababbaabaabaaaabaaaabbaaabbababaabbaabbbbbabbaaaabbbbbbababbbababaabbaaaabbaaabbb
aabbbbbaaabaaaaabbaaaaaabbabaabbabbaaaabbbbbbbbbbbababbbbbababab
bbbbaabbbbbaaabaaabbabbabbababbaababaabb
abaaaabbaaaaaaabbbbaababbabbbbaabbababaaaabbabbb
bbaabaabaabbabbababbabaaababaababaaaaaba
bababbaaabaabbabaaababaaaabbaababaabbaabaabbabbababaabbabaaabaaaaababbabaabaaabb
bbbaaabaaaababbbbbaabaabbaabaaba
abbabaaabbaaaaaabaaaabaa
bbaababbbbaaababbabaaaaaaabababbaaabbabb
abaabaaaaaabbbbbabbabaab
abbbaabbaabaabbbbaaaaabb
abbaaaabaaaaababbaaaabbbababaaabbaaabaaa
bbbaaabbbbaabaabbaaaaaab
aabbbbaababbbabaabaabbba
aababbbbbaabbbaaaaaabbabaababbba
abaaaaaaaabaaaabbabaababbbababbaaabbbaabaaaababbbbbabbba
bbaaaabaabbbbaabaaababaababbaabbbabbbaaa
aabaabbbbaaaabbbaabaabbbbbabaaabbbabbbbaaabbbababababbbbbaabbbbb
aabbaabaabbaaabaaabaaaab
bbbaabbbabaaaaabaaaababb
babaaaabbbbabbaaaaaabaabaaaababa
baaaabbbbbaababbbbaabbab
bbabbbbabbabbbbaabbbbabb
baaabbababbbaaaabaabbabbaabbabbb
abaabaaababbbababbabbaabaababaaaaaabbbbabbbaabaaaaabbaaaababbababbbbbbaa
aabbbaabbabbabaaabbaaaaaababbaaaababaabaabababbbaaabbbbaaabbabab
babbaaabbaaabbababaaaaab
abbaaabaabbaaaabaabaaabaababbaba
bbaabbbbaabbaabbbaaababa
aabaabbabababbbbabbabbabaaabaaab
ababaabbbbbbbaaaababbaba
abbbbaaaabaabbabbaabbabbaabbbabb
babbabaaaabaaaabbbbbabab
bbaaaaaaaabaaaaaababaaaa
aaabaabbaaabaabbbbaaababaaababbbaaabaaabbababaaa
bbbaaabababaaaabbabaabaabbaabaabbbbbaaaa
aaaabbbbabbaabbaaabbbbbb
abbabaaabaabbbababbbaabbbabaaaaaaaabaabbbbbbbabb
bababaabbaabbaaaaababbbbaaababababababbb
bbbbaabbabbbbaabbbaaaaaaaabbbbaababaabaaaabbaaaabbbaabbb
aaaaaaabbbbbaaabbabbaabaababbababababbbaaabbabbbbaaaabaa
babbababbabaabbabababaaa
aaaabaabbabbbbaabbbbbaabbaaaaaabbbaaaaabaabaabbbbabbbabbabbaabba
aaaabbabaabaaaaabaabaaaaaaabaababbbaabaa
aababbaaabababbaaabbaaaa
baabaaaabaaabbbaabbbaabbbbbaaaab
baabbbabaaaabaabbabbababbabaaaabaabbbabaaaabaaabbbbbbabb
ababababbbaaaaaaababbbbbbbbabbab
aabbbaabbaaabbaabaabbaaaabbabaaa
bbbaabaabaabbabbaabaaabbbbbaabaabaabaabbbbabbbaabaabbbbbbbbbbaaaabbbabab
bbaaababaaaaaaaabbbaabbb
abbbbbaababbbbabaabaaabaaabbbbaaaabbaabababbaaaaabaaaabbbbbababa
abbaaabababaababaabbbbaaaabbbbaaababaaaa
abaaaabaabbaaabababababa
abbaababbaaabbbaababababaabbaabbbbaaabaa
abbaabaabaaaabbbbbbbaabbaaababbbababaaba
abbbbbaabaabbbaaaaaaababbaaaaabb
abaabaabbbaabaabaabaabab
aabaabbbabbaaaabaabbbaba
bbbbbbababbbbbbbbbababab
babbbbabbabbbbabbbbbaabbbbbbaabbaaaaaabbbbbaabaa
bbabbabbbbbbbabaabbaaaababbabbababaabbaabbbbbbaaabbbbbbaabbabababbabaababbaaaaba
abaaaabaaaababbbaaabbbbbabbbabbaaaababaa
aababaaabaaabbababaaabba
abbaabbabbabbbaabaaaaabb
aababaabbbaaaaaabaabbbaabbaabbbbabbbbaaaaaaababb
abaabbbbaabbabaabaaabbaaababaabaabbabaab
bbbbaaabbabababbbbbabbaabbbbbabbababaaaa
bbbbaabbaaabaabbbbbbbabb
abbbaaabbaabababbabbbaaa
abbaabababbbabbbbbaaabbbaabaaaabbbbaaababbbbabba
bbbababbabbbbaabbabaabababbbaabaaabbaaab
bbabbbbbabaaabbabbbaaabbbbbaaababbbaababaabaabbaaabbaaabbabaaabbbbbbabbaaabbbaab
aabbbaababbababababbabaabaabababbbaaaaab
babbaaababababbabbbbbbba
baabbaaabaaaababbbabbaabbbabaababbbbbbaaaaaaabaa
abaabaabbbaabaabbaaabbabaabbbbaabaaabaab
babbbbababbbbbbbaaaaabaa
babaabbaabaaaaaabababbaa
aabbbbaabbabbbbaabbbbbaabbbbbababaabbbabbaabaaabbbaabaaa
abbabbaaaabaaaaaababbaaaabababbbaaaabababaaabbaa
baaaaaaabbaababbbbabaabaababbaabbbaaaabaaaababab
bababbabaabaaababbbaabab
aaababbbaaabbbaaaaabbbab
aabaaaaabbbabaababaaaaab
aabbbbababbbbaabababbbabbaabaabb
bbaababbbbababbaaaababaa
baaaaaaabbaabbbbbaabaaba
aabbabaabaabbaabbbbabbbbaaabbaababbabbbb
bbbbbaabbabaaaaaabbabbab
bbaaaabbaaaabbbbbbbbbbbbbabababa
baaabbaabbabbbbaaabaaababbbbaaabbabbabababaaabab
bbbababbabbbabbbabbaabbababaabbb
bbabaababbbbbabaabaaaabaabaababa
bbaababbbbbbbbbabbabaabbbbaababa
abbababbabaabaabbaaabbabbbbababababaabbb
aabaabaaaabaaaaaaaaabbabaaababbbaaabbbbaabaaaaab
babbbbabababbbaabbaabbbaabbaababaabbbbbaabbbbbbabbaaaaaaaaaaaabbabababbb
bbabaaabbaabbaaaaabbbaba
ababbbbbbbababbbbbaaabbbaaababba
bbbbbbabaaaaabababbbbabb
bbbababbabbbbaabaaaaaabbaaaabbbbaaaaaababbaabbab
bbabbaaaaaaabaaaaabbabbb
babbaaabbbbaaaabbaababbbabbbabaababbbbbabbbabbbaaabbbbaaaaabaaaabaaabbbb
bbabbaababbaabbabaabbaaaaabbbbbbbbbabababbabbbababababaa
abbaaababababaabaaaababa
aaabbbbbabaabbababaaaabb
ababaaababaaaaaaaabbabaabbabbbbbbabbbaaa
aabababbbbaaabbaaaababababababbb
baaabbbaabbbbaabbbabbbbababaaaaaabaabaabbbaaabbaabbabbba
aaaabbababbbaaabababbbabbbabbaba
bbaabababbbaabaaaaaaaaaababbaaabbbbababbabbaaaabababbaaababaaabbaababaababaaaabaaaaabbaa
aabbaababababbabbaabbaba
aabbaabbaabaaabababbbbabaababbabbbaababbbabbbbaa
ababababaabbbbbbbaaabaaabababababaaababbaaaaabbababaabaaaaaaaabaaababaab
baabaaaabaaaabbbabbbbbbbabbabbaabbbaabab
abababbabbaaaaaaaabbabbabaabaaba
baabbbaababaaaabbbaababa
abbbabbaababaaababbbbbab
bbaaaabaabbaaabbbbabaaabbabbbbbbbbbabbabbbaabaaa
baaabbaaabaaaaaababaabbb
baaaaaaaaaabaabbaabbbbba
bbbbbbabaaaabbbbbbbaabbabbaabbba
bbbbbababaaababbabbbbaabbabbabbbaabbaaaabaaaabaaaaaabababbbababa
aaaaaabbbabaaaabbabbaaba
aabaaaaaababbbbbaababaabbbbabaabaabbbbaa
bbaabaabababbbabbabbabbababbabababaaabab
baabbbaababaabaabaaabbaaaabaabbababaabaaabbbabaa
bbbababbbabbaaaabbabbaaaaabbbbaababababbaaaababb
ababaaabaabaabbbbaaaabba
abaaaaaabbbababbbaaaabba
aaaabbababbbaaaabbabbaabbbaaabbbabbbbababbbaaaab
babaabbaabbbbaabbbbaabbb
bbaaaabaabbbbaabbbababbbaaaabbabbbaabbab
aababaabaabbbaabaabbaaaa
aaaabbbbbabbabaabaaaaaab
aabbabbabbbbbbabaaaababa
bbabbbbaabbbabbbabaabaabababbbbbabbabbbb
bbbaabbababbabbaababbaabbbbbabaaabbaaabbbaabababaabbababbbbabbbabbaabaaabbbbaaba
bbbabbaabbaababbbabbbbabaaaabbbbababbaaababaabbb
aaaabbabbaaaababaabbbabb
abaabaaababaabbaabababbaababbaaabbbaabbaabaaaabbbaaaaabbbbbaabbbaaaabaaa
bbababbabbbbaaabbabbbbabbbabaaabaaababaa
baabbaaabbbbabaabaaaaabb
aababbaaababbbbbbaaabbbaabbbabab
aaaabaabbbabaaaabbabababbbaabbbbbabbaaaabaaaaaab
babaabbaaababbaaabaabbaa
abaaabaaabaaabaaaababaaaaaaaabbb
bbbbaabbaabbbbaabbbababa
bbaaaabbaabbaabbbbbabbba
aabaaaababbabaaaabaabbabaaaaaababbbbbbbb
aababaabaabaabaaababbaba
bbaaaabababbabaaaaaabbbb
babbbaabbbbababbbbabbaabaababaaa
baaabbbaaaabbbaabbaababa
abbbbbbbbbbabaabbabbabbb
aabaabbbaabbabbababbaaabaabbabbbabbabaab
aaabaabbaabbbbaaabaababa
abbababbaabbbbaaaaaaaaaaaaabbbbbbabaabbb
aababaaaaabbbaabaaaaabbb
aaaabbbbbbbbaabaaabababaababbbaa
aabaabaaabbaaaababbabbab
aaabbbaabbbbabaaabaabbabbabbaaaaababaaabaaabaabbbaaabaab
bbaaabaaaabbabbbaabaaaaababbababbbabbbaa
bbaabbbbaabaabbaabbbbbab
baabbbaabbaaaabaaaaababa
aaaaaaaaabaabaababbaabbb
abbbbababbaababbbaababbaabaabbabbabaabbabaaaaaab
aabbbaabbbabbbaaaababbba
ababbaaaaabaabbbabababbb
aabbbaababbbabbbaababaaaabbbbababbbaabaaabaaabbbbbaabbaa
bbbaaabbaababaabaabbabbb
baaaaaababababbaaaabbbbaaabaabaa
abbaaabababbbbabbbbaabbaabaabaabaaaaaaabbaaabbbb
baaaabaaaabaabaabaaabbaabaababbaabbbbbabaabaaababbbabaabaaaababb
ababbabbbbbabbbbaababaaa
bbbaaabbbababbabbaabaabb
babbbaabbbababbbbbbabaababbbbabbbbbbabbabbaaaabaabababbabbbbbbbbabbbaaabaaabbbab
bbbaabbababaababbaabbbba
aabaaabbbbabbbbbbaabbbaababbaaabbaaaaabaaaaaabbbbbaaaabbbaabababbabbaabaabbabaaaaabbabab
aababbbbabaababbbbaababaaabbaaab
bbaaabaaabababbaaabbababaaaabaaabaabbaaabbaababaaaaababababababbbaaaabbabaaabaaaabbaabba
abbaababbaaababbbbbaaaababbabbbbbbabbaba
aaaabbaaaababaaabaaabbbabbabaabababaabbb
baabaaaababbabaaaaabaabbbaaaaababbbbabba
babbababbbbabbbbbabaaaaabaabbbaaaaaabaabaabbbbabbbaaabaa
babbbbabbbbbbbabaabbbaaa
abbaaaababababbaababaaabbaabbbabbabbabaabbaabbabbbbbabbb
bababaababbbaaabbabbaaba
bbabbbbaababaaabbaabbbaabbbaabbbbbbaabaa
abbaaaababbaaaabbabaaabbbbabbaaabaaaabababbbbbbaaaababaaabaabbaabababaaa
babaaaabaaaabbbbbaabbbabababbaba
abbaaaaaaabaaaaaaaababaa
aabbbaabaaaaabababaaaaab
baaabbbababaaaabbbbbabab
babbbabaaababaaaaabbbabb
bbabbaabbbaaaabaabbaababbbbbabba
abaabaabbababaabbababbabbbabbaaabbaabaaa
aabbaabaaababaaabbababab
bbababbaaaabaaaaaaaabbbbbbbaaaabbbaabbab
bbaababbbabbaaaaababbbabbbaaabba
bbbaaaaabbbabaabaabaaaaaabaabbbaaabbaaab
babbbababababbabbabbbabb
bbbbbbabbbaaabbbbabaabaa
bbaabbbbbbabababbbbabbabbaabbbbaaabaabbbbbabbbbbbaabaaaabababbbaaabbbaaa
bbbbbababbbabbbbabbbbbaaabbaabaaabbbabaaabababbb
aaaabbabaabbaaabbaaabbbabbabbabaababaaabaaabbaabbabbbbbabaabbabb
babbaaabbbabbbbbbbbabbbbbbbaabbaaabaabbabbababaaaaaabbaa
abbaababababaaababaaabaaabaabaabbbababbabbbbbaaaaaaabbba
bbbbbbabbabababbabbbaabbaababababbaabbababbbbbab
baaaabababbaaabbbaabbaba
aababbbbbaabbbabbaaaaaba
abbbbaaaababbabbabbaabbaabaabbba
bbbbbaabbabaaabbbbaaabba
babaabbabaaabbabbbaaaaaaabaaaababaaabbba
bbbaabaaaabbbaaabbabaabb
aaaabbabaabaabbaabaaabba
bbbabbaaaabbaababaabaaba
bbabbbbaabbbbbbbabababbabbbababa
baabbbabbabbaaabbbbaabab
abbaaaaababababbabaaaabaaabbbabaababbbba
aaabbbaabababaaaaabaabab
baaabbaabababbabbaababbb
abbaabbaaabbbaabbaaabaaa
ababaaabbabaaaaaabbaaaaabbaabaaa
ababababbbaabaababbbbbba
bbbbaaabbbabbaaabbbaaababbabbaaabaabbabbabbbabab
ababbaaabababaabbaababbaaaabbbab
abbbabbbaabbabaaaaaabaabbbbbbabb
aababaabbbaaaabbabbaabaababbbbbbbbbbbbba
bbbaaabbbbbaaabbabbbabaa
abbabaaaababbaaabababbaa
babbbabababaaabbbbbaabab
bbbabbbbbabaabaaabbbabab
aababbabbabbaaabbbbabbaaaaabbbba
aabababbbbbbaaabababaababbbaabbabbbababbabbabaabaabbbabbabbbabbb
aababbabbbabbaaaaabbaaaa
bbbabbaabbbbbabbaababababaabaabb
baaaababaabbaabbabbaaabbbaababbabbbbabab
aababbaaabbbbababbaababa
bbabaaabababbbabbaabaaaababbaabb
abbaaaabaaaaaabbaabbabbb
babaaabbbaabbbabbbabaaaabaaaaaaaabbbabaa
aabaabbbbbaabaabbbabaaaabababbba
abbbbababaabaababbbbabbbabaaabbbaaaaabbbabababaabbbaabaa
baaabbbabbaabaabababbbba
bbbaaaaaabbbbbaabbabaaabbbbaaaab
aababaabaabbbbaaaabaaabb
aabaabbbbabaababbabbbabaabababbb
bbabaaabbaaaaaaaaaaabbbbaaaaabaa
baaabbabbbbaaababbaaabbbababbaba
abaabbabbabbabbababaababbabbbabbbbbbbbbb
bbaaaaaabbbbaabbbbbbaaaabbaabbaa
aabbabaabbababbabbbaaaab
baaabbbbaaababaabababbba
bbababbabbbabbaabbaabbbbabbaabbbbbaaabba
baababbaaabbaababaaaabbbbaabbbba
baababbababbbaababbbaabbbbabbaababaabababbbabbba
baaabbbababbaaabaabaaabb
babbabbaaabbbbbbabaaabbb
bbbaaababbbabaaaababbbbbbaababaa
aaababbbaaabbbbbbbabaaabaaabbbaabbaabbbaaaaabbbabbababaa
bbababbababababbbbaaabaa
babaabbaaaabbabbbaababbbaaababab
baaabbbaabbaaababbbbbabb
bbaaabbbbaabbbabbaababbababaaaabbaaaaababbaaabaabababbba
aabaaaababaaabaababaaaabaaaaaaba
aabbbbaaabaabaabababbabbbaabbbaabaaabaabaaaaaabaaababbba
bbabbaabbbbabaaabbbaaababbbaaaabbbababab
aababbbbababaaaababbabba
babbaaaabbbbabaabaababbbababaaabaaaabaabbababbaabbbabbaa
bbabbaaabbabbbbabaabbaaa
bbabbbaabbaaaabbaaaababb
bbabbbbaabbbbababaaaaaaaabbaabbb
abaabaaaaaaaaabbabbabbbb
aabaaabaabababababaaaababbbbbbababaabbbbbbbbbbbbbaaabbbb
babababbbbaabaaababbaaababaababa
abbbbaaaaabaabaabaababbaababbbbabbabbbab
aabaabaaabbbbaaaaaaaaaaaaababaaabbabaaabababaaaaababbaba
aaaabaabababbabbaaabaabbabaababbbaaaabbbbbbababa
bbbaabbabaaabbabbababbba
ababbbabaababbbbbabbbbba
abbaaaaabaabbaaababbbbabbbbbaaababaaaabbaabbabab
baabbaabbbabbaaaabababbb
abbbaaababbabaaababaabaabaaaabba
ababbaaaabbaaaaabbbbaabbaaabaaaabaaaaaabaaaabaaaaabbaaaababbbbaa
abaaaaaabbaaaababababbabababbabbbbaaabaabbbaabaa
aabbbbabababbbbbbaabbaba
bbbababbbaaabbaaabaabbba
abbaabaaababbaaaaaabaabbbbaaaabaabbbaabaaabaabab
bbabbbaabaaabbabaaabaabbbbabbbaabbbaababbaaaaabbaabbbbba`

  input2 := `0: 8 11
1: 41 41 | 48 48
2: 48 85 | 41 128
3: 48 118 | 41 25
4: 131 48 | 70 41
5: 41 110
6: 48 100 | 41 132
7: 41 41
8: 42 | 42 134
9: 89 41 | 100 48
10: 108 48 | 53 41
11: 42 31 | 42 241 31
12: 41 18 | 48 98
13: 43 41 | 69 48
14: 9 48 | 93 41
15: 41 20 | 48 63
16: 48 5 | 41 45
17: 62 48 | 26 41
18: 48 48
19: 48 72 | 41 18
20: 61 48 | 57 41
21: 48 83 | 41 51
22: 103 48 | 18 41
23: 84 41 | 27 48
24: 48 109 | 41 38
25: 48 81 | 41 7
26: 41 72 | 48 81
27: 100 41 | 7 48
28: 41 132 | 48 120
29: 116 48 | 3 41
30: 89 41 | 18 48
31: 48 133 | 41 127
32: 41 96 | 48 95
33: 120 41
34: 48 41 | 48 48
35: 123 48 | 91 41
36: 48 81 | 41 100
37: 41 7 | 48 120
38: 2 48 | 77 41
39: 41 67 | 48 41
40: 81 48 | 7 41
41: "a"
42: 41 68 | 48 105
43: 52 48 | 112 41
44: 101 48 | 114 41
45: 7 48
46: 48 29 | 41 82
47: 120 41 | 81 48
48: "b"
49: 81 41 | 7 48
50: 7 41 | 7 48
51: 41 1 | 48 81
52: 41 73 | 48 104
53: 54 48 | 119 41
54: 41 36 | 48 55
55: 48 1 | 41 18
56: 34 48 | 39 41
57: 41 92 | 48 80
58: 41 22 | 48 111
59: 41 47 | 48 125
60: 30 41 | 86 48
61: 41 37 | 48 6
62: 100 48 | 98 41
63: 41 59 | 48 17
64: 48 44 | 41 74
65: 16 48 | 58 41
66: 41 126 | 48 121
67: 48 | 41
68: 64 41 | 24 48
69: 106 41 | 32 48
70: 87 48 | 115 41
71: 67 34
72: 48 48 | 41 67
73: 1 48 | 103 41
74: 41 107 | 48 124
75: 89 41 | 39 48
76: 48 14 | 41 35
77: 40 48 | 47 41
78: 41 65 | 48 66
79: 67 1
80: 120 48 | 98 41
81: 48 48 | 67 41
82: 23 48 | 21 41
83: 67 89
84: 18 48
85: 120 41 | 89 48
86: 48 89 | 41 39
87: 103 48 | 72 41
88: 7 48 | 89 41
89: 41 41 | 41 48
90: 81 41 | 34 48
91: 34 48
92: 48 103 | 41 34
93: 110 41 | 34 48
94: 41 4 | 48 76
95: 34 41 | 103 48
96: 100 48 | 103 41
97: 41 120 | 48 34
98: 41 48
99: 97 48 | 92 41
100: 48 41 | 67 48
101: 129 48 | 93 41
102: 48 62 | 41 50
103: 67 67
104: 48 1 | 41 81
105: 46 41 | 78 48
106: 88 41
107: 79 41 | 33 48
108: 102 41 | 60 48
109: 48 113 | 41 99
110: 48 41 | 41 41
111: 103 67
112: 48 71 | 41 123
113: 49 41 | 56 48
114: 41 117 | 48 73
115: 67 132
116: 41 93 | 48 122
117: 48 120 | 41 39
118: 89 41 | 132 48
119: 48 12 | 41 19
120: 48 41
121: 48 49 | 41 47
122: 48 120 | 41 89
123: 110 48 | 98 41
124: 83 48 | 130 41
125: 41 120 | 48 132
126: 75 48 | 27 41
127: 41 94 | 48 10
128: 48 72 | 41 1
129: 48 110 | 41 120
130: 103 48 | 89 41
131: 90 41 | 28 48
132: 41 48 | 48 41
133: 48 15 | 41 13
134: 42 | 42 135
135: 42 | 42 136
136: 42 | 42 137
137: 42 | 42 138
138: 42 | 42 139
139: 42 | 42 140
140: 42 | 42 141
141: 42 | 42 142
142: 42 | 42 143
143: 42 | 42 144
144: 42 | 42 145
145: 42 | 42 146
146: 42 | 42 147
147: 42 | 42 148
148: 42 | 42 149
149: 42 | 42 150
150: 42 | 42 151
151: 42 | 42 152
152: 42 | 42 153
153: 42 | 42 154
154: 42 | 42 155
155: 42 | 42 156
156: 42 | 42 157
157: 42 | 42 158
158: 42 | 42 159
159: 42 | 42 160
160: 42 | 42 161
161: 42 | 42 162
162: 42 | 42 163
163: 42 | 42 164
164: 42 | 42 165
165: 42 | 42 166
166: 42 | 42 167
167: 42 | 42 168
168: 42 | 42 169
169: 42 | 42 170
170: 42 | 42 171
171: 42 | 42 172
172: 42 | 42 173
173: 42 | 42 174
174: 42 | 42 175
175: 42 | 42 176
176: 42 | 42 177
177: 42 | 42 178
178: 42 | 42 179
179: 42 | 42 180
180: 42 | 42 181
181: 42 | 42 182
182: 42 | 42 183
183: 42 | 42 184
184: 42 | 42 185
185: 42 | 42 186
186: 42 | 42 187
187: 42 | 42 188
188: 42 | 42 189
189: 42 | 42 190
190: 42 | 42 191
191: 42 | 42 192
192: 42 | 42 193
193: 42 | 42 194
194: 42 | 42 195
195: 42 | 42 196
196: 42 | 42 197
197: 42 | 42 198
198: 42 | 42 199
199: 42 | 42 200
200: 42 | 42 201
201: 42 | 42 202
202: 42 | 42 203
203: 42 | 42 204
204: 42 | 42 205
205: 42 | 42 206
206: 42 | 42 207
207: 42 | 42 208
208: 42 | 42 209
209: 42 | 42 210
210: 42 | 42 211
211: 42 | 42 212
212: 42 | 42 213
213: 42 | 42 214
214: 42 | 42 215
215: 42 | 42 216
216: 42 | 42 217
217: 42 | 42 218
218: 42 | 42 219
219: 42 | 42 220
220: 42 | 42 221
221: 42 | 42 222
222: 42 | 42 223
223: 42 | 42 224
224: 42 | 42 225
225: 42 | 42 226
226: 42 | 42 227
227: 42 | 42 228
228: 42 | 42 229
229: 42 | 42 230
230: 42 | 42 231
231: 42 | 42 232
232: 42 | 42 233
233: 42 | 42 234
234: 42 | 42 235
235: 42 | 42 236
236: 42 | 42 237
237: 42 | 42 238
238: 42 | 42 239
239: 42 | 42 240
240: 42 | 42
241: 42 31 | 42 242 31
242: 42 31 | 42 243 31
243: 42 31 | 42 244 31
244: 42 31 | 42 245 31
245: 42 31 | 42 246 31
246: 42 31 | 42 247 31
247: 42 31 | 42 248 31
248: 42 31 | 42 249 31
249: 42 31 | 42 250 31
250: 42 31 | 42 251 31
251: 42 31 | 42 252 31
252: 42 31 | 42 253 31
253: 42 31 | 42 254 31
254: 42 31 | 42 255 31
255: 42 31 | 42 256 31
256: 42 31 | 42 257 31
257: 42 31 | 42 258 31
258: 42 31 | 42 259 31
259: 42 31 | 42 260 31
260: 42 31 | 42 261 31
261: 42 31 | 42 262 31
262: 42 31 | 42 263 31
263: 42 31 | 42 264 31
264: 42 31 | 42 265 31
265: 42 31 | 42 266 31
266: 42 31 | 42 267 31
267: 42 31 | 42 268 31
268: 42 31 | 42 269 31
269: 42 31 | 42 270 31
270: 42 31 | 42 271 31
271: 42 31 | 42 272 31
272: 42 31 | 42 273 31
273: 42 31 | 42 274 31
274: 42 31 | 42 275 31
275: 42 31 | 42 276 31
276: 42 31 | 42 277 31
277: 42 31 | 42 278 31
278: 42 31 | 42 279 31
279: 42 31 | 42 280 31
280: 42 31 | 42 281 31
281: 42 31 | 42 282 31
282: 42 31 | 42 283 31
283: 42 31 | 42 284 31
284: 42 31 | 42 285 31
285: 42 31 | 42 286 31
286: 42 31 | 42 287 31
287: 42 31 | 42 288 31
288: 42 31 | 42 289 31
289: 42 31 | 42 290 31
290: 42 31 | 42 291 31
291: 42 31 | 42 292 31
292: 42 31 | 42 293 31
293: 42 31 | 42 294 31
294: 42 31 | 42 295 31
295: 42 31 | 42 296 31
296: 42 31 | 42 297 31
297: 42 31 | 42 298 31
298: 42 31 | 42 299 31
299: 42 31 | 42 300 31
300: 42 31 | 42 301 31
301: 42 31 | 42 302 31
302: 42 31 | 42 303 31
303: 42 31 | 42 304 31
304: 42 31 | 42 305 31
305: 42 31 | 42 306 31
306: 42 31 | 42 307 31
307: 42 31 | 42 308 31
308: 42 31 | 42 309 31
309: 42 31 | 42 310 31
310: 42 31 | 42 311 31
311: 42 31 | 42 312 31
312: 42 31 | 42 313 31
313: 42 31 | 42 314 31
314: 42 31 | 42 315 31
315: 42 31 | 42 316 31
316: 42 31 | 42 317 31
317: 42 31 | 42 318 31
318: 42 31 | 42 319 31
319: 42 31 | 42 320 31
320: 42 31 | 42 321 31
321: 42 31 | 42 322 31
322: 42 31 | 42 323 31
323: 42 31 | 42 324 31
324: 42 31 | 42 325 31
325: 42 31 | 42 326 31
326: 42 31 | 42 327 31
327: 42 31 | 42 328 31
328: 42 31 | 42 329 31
329: 42 31 | 42 330 31
330: 42 31 | 42 331 31
331: 42 31 | 42 332 31
332: 42 31 | 42 333 31
333: 42 31 | 42 334 31
334: 42 31 | 42 335 31
335: 42 31 | 42 336 31
336: 42 31 | 42 337 31
337: 42 31 | 42 338 31
338: 42 31 | 42 339 31
339: 42 31 | 42 340 31
340: 42 31 | 42 341 31
341: 42 31 | 42 342 31
342: 42 31

bbaaabbbbaabbaababaababbbabbbaaa
bbbbbaabbabaaaaabbaaaabbabaababbbbabbbbababaabaabaaaaaabbbaaaaab
baabbaabababbabbaaaaabba
ababbbbbbbabaaaaabaabbabbbbaaaabbabbaaba
abbbbabaabbbabbbabbaaaaaaaaaabba
bbaaaaaabaabbbababbaabbabbabaaaabbbaababaaaabbbaabababaa
bbbbbaabbabbbbabaabbbbaaaaaaabba
bbbaabbaaababbbbabbbaabb
abbbbaaabbbaabababaababa
aaaababbabaaabaaabbbbbaaaabaaabaaababbbbbaabaabababbbbbbbaabaaaa
babaababaabbbbabbabbbbabbababbaa
babbaaabababbaaabbbaabab
aaaabaabbbababbbbbbbbbbb
aabbbbbabbbbabbbbbbbbababbbabababaaaababbbaaaabaaaabaaaaabbaabbaaababaaaabaaabab
aaabbaabaaaabbbababbbbbbbabaaabaababaabb
abbaaaaabbaabaabbbabbbab
baabbbaabbabbabaabbabbbb
abababbbaaaaabbbbbbbbabaabaababbabbbbbbbbbabbabbaaabaaba
abbaaaaaaabbabbaaaababab
baabababbaaaaaaabbbbabab
babbabaaaaabaabbbbbaaaab
aabaaaababbaaabbbbbbaaaa
babaaaaabaaaababbbabbabb
abbabaaaabbababbaaabbaba
aaabaaaaaabaaaababaaaabaaaaabaabaababbabaabababa
ababbabbaaabaabababaabbaaabbabababaabbaa
baaaabbbbbabaaaabbbabaabbabbbaaa
abaabaaaababababbaabaabb
aabaabbbbabaaabbaababbaabbbbbbabbbbaabababbabbaa
abaaaaaabbaaaabaaabbaaab
aaaabbbbbabaabaaababaaabbbbbaaabbbbbabbbaaaabaaabababaaa
bbabbbbaaababbbbaaababbbbbbaabbaababbabbaaaabbbaaaaabbbaaabbabbbaaababab
babaaaaaaabaaabaabbbabab
aaaabbaaabaaabaabaabbbabaaaabbbbabbaaabbbbabbbbbbbaabbabbbbbababababaabb
abbbabbabbabbbbaaabaaaaababbaabbaabaaabb
abbaaabbbbaaababbbbaaabbabbbbbbbbbbbbaababbbbbbbabbbbabbabbabbbb
babbbaababbaabaaaabbbbbb
baaaabbbbbbbaabbbabbbababaaaabba
aaabbbaababbaaaabaaaabbbbaaabaaa
abbbbaaaaababbabbbabaaabbbaaabba
babbabaabaaaabbbbbbaaaab
abbababaabbbbaababaaabba
aababbbbbababaababbaababababbbbbbaaabbbbaaabaaab
baabbaabaaabaabaababbaaaabaaaabaaaababbbbbbabbbbbbbbabab
bbababbaaaabaaaaabbabbaa
bbababbbbababbabbbaababbbabaabbaabbaaabaabababaa
baababbabababbabaababbab
abbababbbaababababaaabaaababbbabbbbaaaab
abbbbbaaaabaaaaaabababbb
abbbbaaabaaabbbaababbbaa
bbababbbbabaabaabbbbabaaaaaaaabb
aaaabbaaaabbbbaaaaababba
abaaabaaaababbabbaabaaba
babbabababbaabaaabaaabbb
babababbbaaaabbbbabaababbbbaaaaaaabbaababbaababbaaabaaababaaaaabbbabbbab
baaabbabbaababbaaabbbaaa
bbbababbbabababbaaaabaabababaabaaabbabbb
bababbabbabbbbabaababbbbabaabbabaababaabbbaaaabaaaabbaabaaabbaaa
bbabbbaabaaaabbbbbabaaabababaababaaababa
aabbbbaabaaaaaaaaabbbaabbbbabaababbaaabbbabbabbbaaabababbaabbaba
bbbbaabbbabbabaaabbbbbaaabaaabab
baabbabababaabbbabbabbbaabaabaabaabbbbaabbabbabbabbbaaaabaaaaabaabbbaaba
ababbabbbbbaaabbbbbbaaaa
aaaaaabbababbbabbaababbb
ababbabbaaabaababaababaabaababaaababbababaaaabaa
baaaababaabaaaaabaabaaba
bbabbbbaabbbaabbbabbaaabbababaababbbbbabbbabbbbb
bbbababbbaabbbaaabbbbabb
bbbababbbbbbbaababbbaabbbababbbabbabbabb
abaaaaaaaaababbbaaaabaaa
baaaababbbababbaaaabaaab
babaaabbbbbabbaaaaababbbabbbabaaabbabbaa
bbaaaabbbaaabaaaabababbbbaabaaabbbbbbbbabaaaaabbababaaaaaaaabbabbbaaabaababbbaab
aabaaaabbbabbbaaaaaaabaa
ababaaabaabbbbabbbbbaabbaabbabbababbbbbbaabbbbbb
bababaabbbbabaabbababbba
bbbbbbabababbabbbbaaaaaabaabaaaaaababaaabbbbbbaabbbabbbaaaaaabaa
abbbbaaaababaaabaaabaabaabbbbbbbbabbbaaa
abaabaabbabbaabbbabaabaabbbabbbbabbbaaabaabbabaaabbabbaabbaaabba
abbbbbaaababbaaaabbaabaabbabbaaabbababbbaaaaabaaababaabb
abbaaababbbabbbbabbbbababaabababbbbbbbabababbababbbbabbb
aababaaaaabaaaababbbaaabbbabbaabababbabababbabbb
bbbababaaabababbabaabaaabbabbbab
bbabababbbaabbababbababaabbaaaaaaaabbaba
aaabaabbbbaaabbbabababbb
abababababbbaaaaabbababbaabaaaaabbaaababaabbbababbabbbabaaaaabaa
bbbbabaaabaabaabbbaaaabaaaababbbbbaababababababa
aaaabbaaabaababbbaaabbbb
baaabbabbabbababaaaabbabaabaaaabbbbaabbbabaaaabb
aabaaaaaaaababbbabbbaabaabbbbaabbbbbbbabaababababaababaa
babaabaaaaaaabababbaaabbbbbaabaabbababab
aababbbbbbbabaaaaaabaabbbbbabbaabbabbababbbbaaba
bbbabbaababbbaabaabbaaaa
abbbaabbbbbabaaaaaababab
babbababbbbabbbbaabaaabaabbaabbaabbabaab
abbaaaabbbabbbaaabbbbbab
aabbabbbbabababbaaabbbbaaababaabbabbbbbbbbaabbbaabbbabaababbbaabbbaabbaabbbaabbb
bbabaabaabaabaaaabbababaaabaabaa
bbbbbaabaaaaaabbaabbbaababaabaaabbabaaabbbabbaaaabaaaabb
aababbababaaaaaaabbbbbbbbaabaaab
aaabbbaabbbaabbababaabbabbbbabaabbaabbaaaabababbabbbbbba
bbbbbababbabbaaabbbabaabaabbbbbb
aababaabaaaaaaabaabbaabaaaabbaaaaababbbbabbbabbbbaaabbab
baabbaaabbaaaabbaaabbbbbabaabbabbbbabbba
bababaabbabbbbbabaabbbbabbbaaaab
ababbaaabbbaabbababbbbab
babaaaababbbbbbbaaaaaababaabbbba
aaaabbabbaaabbaabaabbaba
bbbbaabbbbbbbbababababbabababbaa
baaabbaaabbaaaaababababa
aaabaaaabbbbbaababbbbababbabbbaaaabbabbaababaabbbaaabaabababaababbabaabb
babaabbababaaabbbababbbb
baabababbbaaaabbbbabbaaaabbabbaabbabbbbb
babbbbababaaaabaabbbbbaaababbaabaabbbbbabbaabaaaaabababb
babbbaaabaabaaabaabaaabb
aaabbbbbbbaaabbbbaaaaaba
baaaaaabaabbbbbbabaababa
ababaaabbbbaaabbababaaaaababbaababbbbaaaaababaaabbabbbaababababbabaabaabbabbabbbbaabaabb
abaaaaaaabbaababbbaababbbbbaabbbbbaabbaa
ababbbabbabbbaabbbbbbbbb
abbaaaabbaabbbababaaaabb
bbababbbababbbababbabbbb
baabaaaababbabaaababaaabbbbaaabbbbbbbbbaaabbaaaaaaabbbba
baaabbabbaaaaaaababaaaabbaabbbba
abaaabbbbaabbaababababbaaabbbaaabaababab
bbbaabbaaabbabaababbababaabbabbbbabbaabaaaabbaabaabaabaabaaaabbb
aabbaabababaababaaababbbbabaabbb
abbaabbabaaabbabbaabaaab
bbbbabaababbbabaabaabbbb
baabaaaaababbbabbbbaaaaabbbabababaababaa
baaaababbbaaababbbabbaabaaabbaaa
babbabaabbaabbabbbabbabaabaaabab
aababaabbabaabbabbbabbba
aaaaaaaaabbaabaaaabaabbbabbabbab
babbbaabbbabbaaabbabbbbabaaabbbb
abbaaaaaabbbabbababbbbbb
babbabbabaababbaabbbabbbaaababaa
bbabbbaaaabaabbbaababaabaababaabaaaaabba
babbabaaaaabbbbbbbbabbba
baaaaaaabbbbbababbbbbbba
bbbbabaaaaaaabbbababbbba
aababbabbabaababbaaaabbbbaaabaaa
babaabbabbaaaabbbaabaaba
abbbaaabababababbbabbbbb
bbbbaabbababbabbababababbbbaaababbbaaaba
aabaaabaabbabababbaaaabbababaabbbbaabbba
bbbaabbabababbabbbbbbbbb
abbababbbbbbbaabbaaabbababbbbaabaabbaaaaabbbbbab
ababbabbbbaaabbbababbabbaaaaaaba
bbabaabababbbbabbababaabaabaabbaaabaabab
bbaaabbbbabaabbaaabbbaba
baaaabbbbaabbaaaabaaabaabaababaabbbaabaa
aababaababbaaaaaabaaaabb
bbbbabaababaababaaaaababbbbbabaababaababbbbaaaab
abababbbbababbbbbaaaabbbaabaaaaaababbabababbbbbbbbaabbba
bbaaaaaaabbbbaaabaaaaabb
bbabaaaaabbabaaababaaaabaaabbaab
bbbbbababaabbaaaaaaaaaaaabaaaabaabaabaabbaabaaab
baabbbabaaaaaaaabbbabaaaababbabbaaababab
aabbaabaabbabaaaaabbbbaabaaabbababbabbba
aaaabaababaaabaabbaaaaab
abaabbabababbbbbabbbabbabaabbaababababaa
bbbbaabbbbbbaabbbaabbbaaaabbbaabbaabbbababbbabababbaabbb
bbaabbaababbaabaabaaaabaaaabbaaabbababaabbaabbbbbabbaaaabbbbbbababbbababaabbaaaabbaaabbb
aabbbbbaaabaaaaabbaaaaaabbabaabbabbaaaabbbbbbbbbbbababbbbbababab
bbbbaabbbbbaaabaaabbabbabbababbaababaabb
abaaaabbaaaaaaabbbbaababbabbbbaabbababaaaabbabbb
bbaabaabaabbabbababbabaaababaababaaaaaba
bababbaaabaabbabaaababaaaabbaababaabbaabaabbabbababaabbabaaabaaaaababbabaabaaabb
bbbaaabaaaababbbbbaabaabbaabaaba
abbabaaabbaaaaaabaaaabaa
bbaababbbbaaababbabaaaaaaabababbaaabbabb
abaabaaaaaabbbbbabbabaab
abbbaabbaabaabbbbaaaaabb
abbaaaabaaaaababbaaaabbbababaaabbaaabaaa
bbbaaabbbbaabaabbaaaaaab
aabbbbaababbbabaabaabbba
aababbbbbaabbbaaaaaabbabaababbba
abaaaaaaaabaaaabbabaababbbababbaaabbbaabaaaababbbbbabbba
bbaaaabaabbbbaabaaababaababbaabbbabbbaaa
aabaabbbbaaaabbbaabaabbbbbabaaabbbabbbbaaabbbababababbbbbaabbbbb
aabbaabaabbaaabaaabaaaab
bbbaabbbabaaaaabaaaababb
babaaaabbbbabbaaaaaabaabaaaababa
baaaabbbbbaababbbbaabbab
bbabbbbabbabbbbaabbbbabb
baaabbababbbaaaabaabbabbaabbabbb
abaabaaababbbababbabbaabaababaaaaaabbbbabbbaabaaaaabbaaaababbababbbbbbaa
aabbbaabbabbabaaabbaaaaaababbaaaababaabaabababbbaaabbbbaaabbabab
babbaaabbaaabbababaaaaab
abbaaabaabbaaaabaabaaabaababbaba
bbaabbbbaabbaabbbaaababa
aabaabbabababbbbabbabbabaaabaaab
ababaabbbbbbbaaaababbaba
abbbbaaaabaabbabbaabbabbaabbbabb
babbabaaaabaaaabbbbbabab
bbaaaaaaaabaaaaaababaaaa
aaabaabbaaabaabbbbaaababaaababbbaaabaaabbababaaa
bbbaaabababaaaabbabaabaabbaabaabbbbbaaaa
aaaabbbbabbaabbaaabbbbbb
abbabaaabaabbbababbbaabbbabaaaaaaaabaabbbbbbbabb
bababaabbaabbaaaaababbbbaaababababababbb
bbbbaabbabbbbaabbbaaaaaaaabbbbaababaabaaaabbaaaabbbaabbb
aaaaaaabbbbbaaabbabbaabaababbababababbbaaabbabbbbaaaabaa
babbababbabaabbabababaaa
aaaabaabbabbbbaabbbbbaabbaaaaaabbbaaaaabaabaabbbbabbbabbabbaabba
aaaabbabaabaaaaabaabaaaaaaabaababbbaabaa
aababbaaabababbaaabbaaaa
baabaaaabaaabbbaabbbaabbbbbaaaab
baabbbabaaaabaabbabbababbabaaaabaabbbabaaaabaaabbbbbbabb
ababababbbaaaaaaababbbbbbbbabbab
aabbbaabbaaabbaabaabbaaaabbabaaa
bbbaabaabaabbabbaabaaabbbbbaabaabaabaabbbbabbbaabaabbbbbbbbbbaaaabbbabab
bbaaababaaaaaaaabbbaabbb
abbbbbaababbbbabaabaaabaaabbbbaaaabbaabababbaaaaabaaaabbbbbababa
abbaaabababaababaabbbbaaaabbbbaaababaaaa
abaaaabaabbaaabababababa
abbaababbaaabbbaababababaabbaabbbbaaabaa
abbaabaabaaaabbbbbbbaabbaaababbbababaaba
abbbbbaabaabbbaaaaaaababbaaaaabb
abaabaabbbaabaabaabaabab
aabaabbbabbaaaabaabbbaba
bbbbbbababbbbbbbbbababab
babbbbabbabbbbabbbbbaabbbbbbaabbaaaaaabbbbbaabaa
bbabbabbbbbbbabaabbaaaababbabbababaabbaabbbbbbaaabbbbbbaabbabababbabaababbaaaaba
abaaaabaaaababbbaaabbbbbabbbabbaaaababaa
aababaaabaaabbababaaabba
abbaabbabbabbbaabaaaaabb
aababaabbbaaaaaabaabbbaabbaabbbbabbbbaaaaaaababb
abaabbbbaabbabaabaaabbaaababaabaabbabaab
bbbbaaabbabababbbbbabbaabbbbbabbababaaaa
bbbbaabbaaabaabbbbbbbabb
abbbaaabbaabababbabbbaaa
abbaabababbbabbbbbaaabbbaabaaaabbbbaaababbbbabba
bbbababbabbbbaabbabaabababbbaabaaabbaaab
bbabbbbbabaaabbabbbaaabbbbbaaababbbaababaabaabbaaabbaaabbabaaabbbbbbabbaaabbbaab
aabbbaababbababababbabaabaabababbbaaaaab
babbaaababababbabbbbbbba
baabbaaabaaaababbbabbaabbbabaababbbbbbaaaaaaabaa
abaabaabbbaabaabbaaabbabaabbbbaabaaabaab
babbbbababbbbbbbaaaaabaa
babaabbaabaaaaaabababbaa
aabbbbaabbabbbbaabbbbbaabbbbbababaabbbabbaabaaabbbaabaaa
abbabbaaaabaaaaaababbaaaabababbbaaaabababaaabbaa
baaaaaaabbaababbbbabaabaababbaabbbaaaabaaaababab
bababbabaabaaababbbaabab
aaababbbaaabbbaaaaabbbab
aabaaaaabbbabaababaaaaab
aabbbbababbbbaabababbbabbaabaabb
bbaababbbbababbaaaababaa
baaaaaaabbaabbbbbaabaaba
aabbabaabaabbaabbbbabbbbaaabbaababbabbbb
bbbbbaabbabaaaaaabbabbab
bbaaaabbaaaabbbbbbbbbbbbbabababa
baaabbaabbabbbbaaabaaababbbbaaabbabbabababaaabab
bbbababbabbbabbbabbaabbababaabbb
bbabaababbbbbabaabaaaabaabaababa
bbaababbbbbbbbbabbabaabbbbaababa
abbababbabaabaabbaaabbabbbbababababaabbb
aabaabaaaabaaaaaaaaabbabaaababbbaaabbbbaabaaaaab
babbbbabababbbaabbaabbbaabbaababaabbbbbaabbbbbbabbaaaaaaaaaaaabbabababbb
bbabaaabbaabbaaaaabbbaba
ababbbbbbbababbbbbaaabbbaaababba
bbbbbbabaaaaabababbbbabb
bbbababbabbbbaabaaaaaabbaaaabbbbaaaaaababbaabbab
bbabbaaaaaaabaaaaabbabbb
babbaaabbbbaaaabbaababbbabbbabaababbbbbabbbabbbaaabbbbaaaaabaaaabaaabbbb
bbabbaababbaabbabaabbaaaaabbbbbbbbbabababbabbbababababaa
abbaaababababaabaaaababa
aaabbbbbabaabbababaaaabb
ababaaababaaaaaaaabbabaabbabbbbbbabbbaaa
aabababbbbaaabbaaaababababababbb
baaabbbaabbbbaabbbabbbbababaaaaaabaabaabbbaaabbaabbabbba
aaaabbababbbaaabababbbabbbabbaba
bbaabababbbaabaaaaaaaaaababbaaabbbbababbabbaaaabababbaaababaaabbaababaababaaaabaaaaabbaa
aabbaababababbabbaabbaba
aabbaabbaabaaabababbbbabaababbabbbaababbbabbbbaa
ababababaabbbbbbbaaabaaabababababaaababbaaaaabbababaabaaaaaaaabaaababaab
baabaaaabaaaabbbabbbbbbbabbabbaabbbaabab
abababbabbaaaaaaaabbabbabaabaaba
baabbbaababaaaabbbaababa
abbbabbaababaaababbbbbab
bbaaaabaabbaaabbbbabaaabbabbbbbbbbbabbabbbaabaaa
baaabbaaabaaaaaababaabbb
baaaaaaaaaabaabbaabbbbba
bbbbbbabaaaabbbbbbbaabbabbaabbba
bbbbbababaaababbabbbbaabbabbabbbaabbaaaabaaaabaaaaaabababbbababa
aaaaaabbbabaaaabbabbaaba
aabaaaaaababbbbbaababaabbbbabaabaabbbbaa
bbaabaabababbbabbabbabbababbabababaaabab
baabbbaababaabaabaaabbaaaabaabbababaabaaabbbabaa
bbbababbbabbaaaabbabbaaaaabbbbaababababbaaaababb
ababaaabaabaabbbbaaaabba
abaaaaaabbbababbbaaaabba
aaaabbababbbaaaabbabbaabbbaaabbbabbbbababbbaaaab
babaabbaabbbbaabbbbaabbb
bbaaaabaabbbbaabbbababbbaaaabbabbbaabbab
aababaabaabbbaabaabbaaaa
aaaabbbbbabbabaabaaaaaab
aabbabbabbbbbbabaaaababa
bbabbbbaabbbabbbabaabaabababbbbbabbabbbb
bbbaabbababbabbaababbaabbbbbabaaabbaaabbbaabababaabbababbbbabbbabbaabaaabbbbaaba
bbbabbaabbaababbbabbbbabaaaabbbbababbaaababaabbb
aaaabbabbaaaababaabbbabb
abaabaaababaabbaabababbaababbaaabbbaabbaabaaaabbbaaaaabbbbbaabbbaaaabaaa
bbababbabbbbaaabbabbbbabbbabaaabaaababaa
baabbaaabbbbabaabaaaaabb
aababbaaababbbbbbaaabbbaabbbabab
aaaabaabbbabaaaabbabababbbaabbbbbabbaaaabaaaaaab
babaabbaaababbaaabaabbaa
abaaabaaabaaabaaaababaaaaaaaabbb
bbbbaabbaabbbbaabbbababa
bbaaaabbaabbaabbbbbabbba
aabaaaababbabaaaabaabbabaaaaaababbbbbbbb
aababaabaabaabaaababbaba
bbaaaabababbabaaaaaabbbb
babbbaabbbbababbbbabbaabaababaaa
baaabbbaaaabbbaabbaababa
abbbbbbbbbbabaabbabbabbb
aabaabbbaabbabbababbaaabaabbabbbabbabaab
aaabaabbaabbbbaaabaababa
abbababbaabbbbaaaaaaaaaaaaabbbbbbabaabbb
aababaaaaabbbaabaaaaabbb
aaaabbbbbbbbaabaaabababaababbbaa
aabaabaaabbaaaababbabbab
aaabbbaabbbbabaaabaabbabbabbaaaaababaaabaaabaabbbaaabaab
bbaaabaaaabbabbbaabaaaaababbababbbabbbaa
bbaabbbbaabaabbaabbbbbab
baabbbaabbaaaabaaaaababa
aaaaaaaaabaabaababbaabbb
abbbbababbaababbbaababbaabaabbabbabaabbabaaaaaab
aabbbaabbbabbbaaaababbba
ababbaaaaabaabbbabababbb
aabbbaababbbabbbaababaaaabbbbababbbaabaaabaaabbbbbaabbaa
bbbaaabbaababaabaabbabbb
baaaaaababababbaaaabbbbaaabaabaa
abbaaabababbbbabbbbaabbaabaabaabaaaaaaabbaaabbbb
baaaabaaaabaabaabaaabbaabaababbaabbbbbabaabaaababbbabaabaaaababb
ababbabbbbbabbbbaababaaa
bbbaaabbbababbabbaabaabb
babbbaabbbababbbbbbabaababbbbabbbbbbabbabbaaaabaabababbabbbbbbbbabbbaaabaaabbbab
bbbaabbababaababbaabbbba
aabaaabbbbabbbbbbaabbbaababbaaabbaaaaabaaaaaabbbbbaaaabbbaabababbabbaabaabbabaaaaabbabab
aababbbbabaababbbbaababaaabbaaab
bbaaabaaabababbaaabbababaaaabaaabaabbaaabbaababaaaaababababababbbaaaabbabaaabaaaabbaabba
abbaababbaaababbbbbaaaababbabbbbbbabbaba
aaaabbaaaababaaabaaabbbabbabaabababaabbb
baabaaaababbabaaaaabaabbbaaaaababbbbabba
babbababbbbabbbbbabaaaaabaabbbaaaaaabaabaabbbbabbbaaabaa
babbbbabbbbbbbabaabbbaaa
abbaaaababababbaababaaabbaabbbabbabbabaabbaabbabbbbbabbb
bababaababbbaaabbabbaaba
bbabbbbaababaaabbaabbbaabbbaabbbbbbaabaa
abbaaaababbaaaabbabaaabbbbabbaaabaaaabababbbbbbaaaababaaabaabbaabababaaa
babaaaabaaaabbbbbaabbbabababbaba
abbaaaaaaabaaaaaaaababaa
aabbbaabaaaaabababaaaaab
baaabbbababaaaabbbbbabab
babbbabaaababaaaaabbbabb
bbabbaabbbaaaabaabbaababbbbbabba
abaabaabbababaabbababbabbbabbaaabbaabaaa
aabbaabaaababaaabbababab
bbababbaaaabaaaaaaaabbbbbbbaaaabbbaabbab
bbaababbbabbaaaaababbbabbbaaabba
bbbaaaaabbbabaabaabaaaaaabaabbbaaabbaaab
babbbababababbabbabbbabb
bbbbbbabbbaaabbbbabaabaa
bbaabbbbbbabababbbbabbabbaabbbbaaabaabbbbbabbbbbbaabaaaabababbbaaabbbaaa
bbbbbababbbabbbbabbbbbaaabbaabaaabbbabaaabababbb
aaaabbabaabbaaabbaaabbbabbabbabaababaaabaaabbaabbabbbbbabaabbabb
babbaaabbbabbbbbbbbabbbbbbbaabbaaabaabbabbababaaaaaabbaa
abbaababababaaababaaabaaabaabaabbbababbabbbbbaaaaaaabbba
bbbbbbabbabababbabbbaabbaababababbaabbababbbbbab
baaaabababbaaabbbaabbaba
aababbbbbaabbbabbaaaaaba
abbbbaaaababbabbabbaabbaabaabbba
bbbbbaabbabaaabbbbaaabba
babaabbabaaabbabbbaaaaaaabaaaababaaabbba
bbbaabaaaabbbaaabbabaabb
aaaabbabaabaabbaabaaabba
bbbabbaaaabbaababaabaaba
bbabbbbaabbbbbbbabababbabbbababa
baabbbabbabbaaabbbbaabab
abbaaaaababababbabaaaabaaabbbabaababbbba
aaabbbaabababaaaaabaabab
baaabbaabababbabbaababbb
abbaabbaaabbbaabbaaabaaa
ababaaabbabaaaaaabbaaaaabbaabaaa
ababababbbaabaababbbbbba
bbbbaaabbbabbaaabbbaaababbabbaaabaabbabbabbbabab
ababbaaabababaabbaababbaaaabbbab
abbbabbbaabbabaaaaaabaabbbbbbabb
aababaabbbaaaabbabbaabaababbbbbbbbbbbbba
bbbaaabbbbbaaabbabbbabaa
abbabaaaababbaaabababbaa
babbbabababaaabbbbbaabab
bbbabbbbbabaabaaabbbabab
aababbabbabbaaabbbbabbaaaaabbbba
aabababbbbbbaaabababaababbbaabbabbbababbabbabaabaabbbabbabbbabbb
aababbabbbabbaaaaabbaaaa
bbbabbaabbbbbabbaababababaabaabb
baaaababaabbaabbabbaaabbbaababbabbbbabab
aababbaaabbbbababbaababa
bbabaaabababbbabbaabaaaababbaabb
abbaaaabaaaaaabbaabbabbb
babaaabbbaabbbabbbabaaaabaaaaaaaabbbabaa
aabaabbbbbaabaabbbabaaaabababbba
abbbbababaabaababbbbabbbabaaabbbaaaaabbbabababaabbbaabaa
baaabbbabbaabaabababbbba
bbbaaaaaabbbbbaabbabaaabbbbaaaab
aababaabaabbbbaaaabaaabb
aabaabbbbabaababbabbbabaabababbb
bbabaaabbaaaaaaaaaaabbbbaaaaabaa
baaabbabbbbaaababbaaabbbababbaba
abaabbabbabbabbababaababbabbbabbbbbbbbbb
bbaaaaaabbbbaabbbbbbaaaabbaabbaa
aabbabaabbababbabbbaaaab
baaabbbbaaababaabababbba
bbababbabbbabbaabbaabbbbabbaabbbbbaaabba
baababbaaabbaababaaaabbbbaabbbba
baababbababbbaababbbaabbbbabbaababaabababbbabbba
baaabbbababbaaabaabaaabb
babbabbaaabbbbbbabaaabbb
bbbaaababbbabaaaababbbbbbaababaa
aaababbbaaabbbbbbbabaaabaaabbbaabbaabbbaaaaabbbabbababaa
bbababbababababbbbaaabaa
babaabbaaaabbabbbaababbbaaababab
baaabbbaabbaaababbbbbabb
bbaaabbbbaabbbabbaababbababaaaabbaaaaababbaaabaabababbba
aabaaaababaaabaababaaaabaaaaaaba
aabbbbaaabaabaabababbabbbaabbbaabaaabaabaaaaaabaaababbba
bbabbaabbbbabaaabbbaaababbbaaaabbbababab
aababbbbababaaaababbabba
babbaaaabbbbabaabaababbbababaaabaaaabaabbababbaabbbabbaa
bbabbaaabbabbbbabaabbaaa
bbabbbaabbaaaabbaaaababb
bbabbbbaabbbbababaaaaaaaabbaabbb
abaabaaaaaaaaabbabbabbbb
aabaaabaabababababaaaababbbbbbababaabbbbbbbbbbbbbaaabbbb
babababbbbaabaaababbaaababaababa
abbbbaaaaabaabaabaababbaababbbbabbabbbab
aabaabaaabbbbaaaaaaaaaaaaababaaabbabaaabababaaaaababbaba
aaaabaabababbabbaaabaabbabaababbbaaaabbbbbbababa
bbbaabbabaaabbabbababbba
ababbbabaababbbbbabbbbba
abbaaaaabaabbaaababbbbabbbbbaaababaaaabbaabbabab
baabbaabbbabbaaaabababbb
abbbaaababbabaaababaabaabaaaabba
ababbaaaabbaaaaabbbbaabbaaabaaaabaaaaaabaaaabaaaaabbaaaababbbbaa
abaaaaaabbaaaababababbabababbabbbbaaabaabbbaabaa
aabbbbabababbbbbbaabbaba
bbbababbbaaabbaaabaabbba
abbaabaaababbaaaaaabaabbbbaaaabaabbbaabaaabaabab
bbabbbaabaaabbabaaabaabbbbabbbaabbbaababbaaaaabbaabbbbba`

  s1 := strings.Split(input1, "\n\n")
  rules1    := strings.Split(s1[0], "\n")
  messages1 := strings.Split(s1[1], "\n")

  s2 := strings.Split(input2, "\n\n")
  rules2    := strings.Split(s2[0], "\n")
  messages2 := strings.Split(s2[1], "\n")

  p1 := P1(rules1, messages1)
  p2 := P1(rules2, messages2)

  fmt.Printf("P1: %d, P2: %d\n", p1, p2)
}

func P1 (rulesIn []string, messages []string) int {
  ret := 0

  ruleMatcher := regexp.MustCompile("[0-9]+: (.*)")

  rules := make(map[int]string)

  for idx, rule := range rulesIn {
    ruleMatch := ruleMatcher.FindStringSubmatch(rule)
    rules[idx] = ruleMatch[1]
  }

  regex := "^" + GetRegex(rules, 0, make(map[int]string)) + "$"

  messageMatcher := regexp.MustCompile(regex)

  for _, msg := range messages {
    if messageMatcher.MatchString(msg) {
      ret++
    }
  }

  return ret
}

func P2 (input string) int {
  ret := 0

  return ret
}

func GetRegex(input map[int]string, idx int, memo map[int]string) string {
  rule       := input[idx]
  ret, found := memo[idx]


  if (!found) {
    if strings.Contains(rule, "\"") {
      // This is a base rule
      ret = string(rule[1])
    } else {
      components := strings.Split(rule, " ")

      ret += "("

      for _, component := range components {
        if component == "|" {
          ret += "|"
        } else {
          v, _ := strconv.Atoi(component)
          ret += GetRegex(input, v, memo)
        }
      }

      ret += ")"
    }

    memo[idx] = ret
  }

  return ret
}
