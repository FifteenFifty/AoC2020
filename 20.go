package main

import (
  "fmt"
  "strings"
  "regexp"
  "strconv"
  "math"
)

func main() {
  input := `Tile 2411:
.#.##..#.#
.#.......#
.##...#.#.
......#...
#.#.......
##.....#.#
#..#...#..
##........
##....#...
##..#.##.#

Tile 1997:
#.#......#
.#......##
#.##..#..#
...#..#.#.
.#.#...#..
#.........
..#..#....
#.#..#.##.
#..##.....
#...#.#.#.

Tile 1427:
#...#..#..
..##.....#
..#...####
#..#.#...#
..#.#.#..#
..####....
#.#.###.#.
......#...
#.........
#.#.##...#

Tile 2161:
###.####.#
#....###.#
#.#......#
..#....#..
...#.....#
.#..#...##
#.#.....##
#..#......
.....#.#..
...###.#.#

Tile 1321:
###...#.#.
.#...#..##
.#.......#
..........
#.#.......
....#.....
######....
#.....#...
##.......#
#.###..###

Tile 1181:
.#..#..###
###...#.##
#.........
....###..#
#.#...#..#
#....###..
##..##...#
#....##.##
#.......#.
###.##...#

Tile 2749:
#####..##.
##........
......#.#.
#..#...#.#
.....##.#.
.....#...#
#.........
#........#
....#..#.#
.#..#...#.

Tile 3911:
##...#####
#...#..#.#
.##...#..#
#...#.....
#..#.#...#
..#......#
#.#.......
....##..#.
..#.#..##.
####.#..#.

Tile 3257:
#..#.##...
.#.....#.#
##.#.....#
##.#.#....
#..##..#..
........##
#.##...#.#
#.#..#...#
........##
.##.####.#

Tile 2237:
.###.#####
....#....#
.......###
#........#
..#.......
.#.#......
#...##...#
#.....#.#.
#...#.....
####.#####

Tile 2389:
#.#.###...
#.#.#...##
......#...
.####....#
..#..#..#.
#.......##
..###....#
#........#
#....##...
..#.#..##.

Tile 3209:
..##.#####
.........#
##..#.##..
...#.#...#
....#.#..#
#...#.....
#.#.#.....
#...#....#
#.#......#
...#..####

Tile 2579:
#...#.##..
....##.###
#......###
.......#..
##...#....
#.#......#
.##.#....#
.#.......#
#...#.....
##..#.##..

Tile 2087:
.###.##.#.
.##...##.#
..####.#.#
...#......
...#......
......#..#
#.##......
.#.#.#.#..
....##..##
...##.#..#

Tile 3517:
###.#...##
#.........
#.....#...
###..#.#..
##......#.
..#....#.#
#.....#...
..#...#.##
##...#.#..
##..##.###

Tile 3347:
##.#..#...
.........#
..........
.#.#.....#
..........
....#..#.#
#........#
##.#...#.#
....#..#..
##.#.#....

Tile 2711:
.##.#.#...
...#.#....
..........
#......###
..##......
#.........
#..#..##..
.##...##..
#......#..
#..#.##.#.

Tile 3877:
###.#..#.#
#...##...#
..#.#..#.#
##......##
.#........
##..#....#
..###..#.#
#.#.......
#.#......#
...###.###

Tile 1721:
###...##.#
......#..#
.##.##....
....##.#..
##..#..#..
###..#..#.
####..##.#
#.#.#....#
.......#.#
#..##....#

Tile 3457:
.#.#..##.#
..##.#.#..
##...#...#
...#......
..#....##.
#...#.#..#
..#.......
...#......
........#.
.#.#.##..#

Tile 1231:
.#.......#
#........#
#..#..#...
..#......#
.....#....
#.........
#...#..#.#
.....#..#.
.#.#....#.
###...##..

Tile 1901:
.###.#.#.#
...##.....
#......#.#
...##.....
....#.....
#.........
..#......#
##.....#..
#.........
#..##..###

Tile 3329:
#..#..#.#.
....#.....
###.....##
.........#
##..#.....
#.......##
#..#.....#
#..#.#..#.
...#......
...#...##.

Tile 1291:
..#....#.#
#.#...##.#
#..#....##
#..###...#
#...#.#...
####......
#.....##.#
#.#..#....
##....#..#
#.###.#...

Tile 1481:
#..#...##.
....#..#..
...##.....
#..#..##..
###.#..##.
##.#.##...
#..#.....#
#......#.#
#.#....#..
.#..#..#.#

Tile 2063:
.###..#..#
...#......
#..#.....#
#..#..#...
.....#.###
...#......
#.#.....#.
#........#
.#...#...#
##.#....##

Tile 2297:
.#..#..#.#
#..#......
#......#.#
#.#..###.#
..........
##..#.#.##
#......#..
#....#.#.#
#........#
.#..###...

Tile 2731:
..#.......
......#...
#..#...#..
..........
#..#...#..
#.#..#..##
##..#..##.
..........
#.##..#...
##....##..

Tile 1453:
....#.##.#
#..#..##.#
##.#.#....
#..##.#...
#......#..
#......###
...##...#.
#.##......
##........
#######.##

Tile 1117:
.###.#...#
.#.#.....#
..#.#...##
#.......#.
#.....#..#
#.###....#
##.#......
#.####..##
#..#..#.#.
##...#####

Tile 1823:
#.....#...
#......##.
.....#.#..
......###.
..#..###.#
....#.##.#
.##..#.#.#
#.#.......
##........
.#...##..#

Tile 3163:
..#####.##
...#......
#..###...#
..##..#...
.#.......#
..#....#.#
###.#.....
#..#..#...
..##.#....
.##.#.#..#

Tile 2083:
#..######.
....##.#.#
.....###.#
#...#..#..
#....#...#
...#...#.#
#..####...
#....#....
#..#....#.
#....#.#..

Tile 2659:
#.###.#...
...#......
#..#....##
..###.##..
##.......#
#.......#.
#....#...#
#..#.....#
..#.#.....
..#..##.#.

Tile 2113:
.##.######
.......###
#....#....
....#..#..
#.##......
####..##..
.##..####.
...#....##
.#.....#..
#.#.##..##

Tile 3943:
###..#....
.#.....#..
#.........
##....##..
#.#....#.#
#.#.#.....
#....#..#.
.##..##..#
....##....
.#.#.#####

Tile 3719:
#...#.#.#.
.........#
#.#.......
......#.#.
.#...#.#.#
.#....####
##........
.##..#.#.#
#.........
#.#.#....#

Tile 1409:
####.#####
#....#....
##.#....#.
#..#......
#....#....
#.#...#...
##.....#..
#.....#..#
......#...
..####...#

Tile 3547:
.#########
.#.##...#.
...#..#.#.
#.##...###
#......#..
##...#....
##.#.#.###
##........
.....#...#
###.##.##.

Tile 1123:
...##....#
.........#
#...#...##
..........
#.........
#......#.#
#..#.#.#.#
......#..#
#.#...#.##
####.#....

Tile 2203:
.####.#..#
##........
##.......#
#....#....
.#......##
.#..#..#.#
......#..#
##.###...#
#....#....
.##..##.##

Tile 3511:
#.#..#.#.#
#......###
#..#....#.
#.####....
##...#...#
#.##...##.
##.##..##.
.....#...#
#...##...#
.#.#..##.#

Tile 3499:
..#..#.###
#.........
..##..#...
...#.#.#.#
...#.#.#..
#.....##.#
...#......
..#...###.
##..###..#
#.#....###

Tile 1039:
#.#....##.
#..###.#.#
###.....#.
##..#....#
.#....##.#
#..#######
.##.#....#
#..#......
##.#...##.
#..#..##.#

Tile 3323:
#..#.#.##.
#......#..
.........#
.........#
.#....##.#
#.....#..#
#....##...
#...##...#
.###...#.#
#...###.##

Tile 1033:
###...##.#
#..#..#.##
.........#
.#.#.#..#.
#.........
#......##.
#...#.##.#
..........
###.......
.##.#.....

Tile 2269:
....#.##..
#........#
.#..#.#..#
#......#.#
#...#.###.
#..#.....#
....#....#
........##
#.#...#...
.#...#.#.#

Tile 1279:
...##.###.
....#...#.
#..#......
#......##.
.#.#....##
.#...##..#
..#..#...#
#......#..
#..#....#.
.#.#...#.#

Tile 2549:
..#..###..
..........
........##
....#....#
.........#
#...#..#.#
....##..#.
.#........
..##..#...
##.##.####

Tile 2377:
..#.######
..##..#..#
..#.....#.
#.#.#.#...
#....#....
..........
#..##....#
#...####..
.......#.#
#.#.#..##.

Tile 3169:
.######.##
#..#..####
#.....#..#
#.##..##.#
....#...##
..##......
#..#...###
.....##...
#......###
#.#....#..

Tile 1979:
###.#.##..
#..#...#..
#..#..#.#.
....##....
#.#.......
#..#.....#
....#.....
.#..#..#.#
#.#.###..#
#.####.#.#

Tile 2131:
..##..#.##
##...#....
#...#.#..#
.....#...#
.#.#.#..#.
##......#.
......#...
#..##.....
.###...###
#...######

Tile 2531:
.#.######.
....#..#.#
...###...#
.........#
#....#.#..
.....#...#
#...#....#
....#.....
..#.......
.....#.##.

Tile 3167:
####.###..
#..#.....#
#.......##
##..#..#.#
#.........
#.....#.#.
##.......#
...#.#....
.........#
######..##

Tile 3191:
#....#####
#...#.....
..##...#..
#.......#.
#..#.###.#
#####...##
#.#...###.
#.#..#.##.
#.........
#.##.##.#.

Tile 3767:
#..#..####
#......#..
#.#.##....
.....#..#.
.#...#.#..
..#.#..###
..##......
.###.##..#
.##....#.#
.##.##..##

Tile 2633:
##...##.##
#....#..##
...##.#...
#.#.#.....
.......###
.#...#....
#....#.###
.#.....#.#
##....#...
....#..##.

Tile 2677:
...#.##.#.
......#...
.....#....
#......#..
#...#.###.
..#....#.#
.........#
#..#...#.#
#.#....###
#.#..#....

Tile 3989:
.###.###.#
#.#..#..##
......#.#.
...#......
..####...#
....###.#.
#.###.#..#
....##....
#.........
.#.##.####

Tile 2897:
####...#.#
..........
#.#.......
#...#.....
......#...
##.###.#..
.##..##..#
...##.##..
...##...#.
.#.######.

Tile 3803:
#.#..#.##.
###....#..
#...#..#.#
.......###
#.....#..#
##.....#..
##..##..##
.#....##.#
.###......
#..#.#####

Tile 1187:
.#.#####..
.#.#......
#.#.####.#
#...#.....
..#.......
.#.....#..
#...#.#..#
#......###
#.##..#..#
#.##.##.##

Tile 1801:
#.##...#..
.##......#
#..####..#
#.##...#.#
#......#.#
.....#.#..
..###...#.
..#.#.....
#...#.#...
#....####.

Tile 2729:
.......#..
#.#...#..#
..#.#..#..
.....#.#..
#.........
#.#..#....
.#.......#
#.#....#.#
#.#...#.##
##.#.###.#

Tile 2969:
.......#.#
##..#....#
#.....#.##
##.....#.#
#..#......
.#.#.#....
#....#.#.#
##......##
#......##.
.##.##....

Tile 3079:
...##.####
.....#...#
..........
..#.#.....
.........#
.#...##.##
.###.....#
..#.#..##.
#...##...#
.##...##..

Tile 3407:
...#.#.#..
.#...##.##
##....#..#
.#....#.#.
.##.#....#
##......##
#...#...##
#..#.....#
#.....#...
##....##.#

Tile 3389:
####.#..#.
#..##.....
##...##.#.
#......#..
.....#..##
#...##..##
#...##...#
...##..#..
....##..##
#.#...#...

Tile 1663:
#...##.##.
..#..#.#.#
#.....#.#.
#.........
..##.##..#
.....#..#.
#....#....
#.##......
.##.#...##
.....#...#

Tile 1787:
#..##....#
......#.#.
#......##.
#.#......#
#....#...#
#.#.......
.....#...#
#...#..#..
.#......##
..###.#.##

Tile 2917:
##.#.#####
#..#..#.##
..##..#..#
#..#.##.##
#.#..#....
.#.#.##..#
.####.....
#..#...#.#
#...#....#
.#.#.##..#

Tile 3067:
##.#.#...#
.###..#..#
.##.#.#...
#.#.....##
..#.......
#...#...#.
##..#....#
...##..#.#
#.......#.
.###...#.#

Tile 2963:
.#..#.#...
#..#.##..#
##........
..#....###
###.#.#...
.#.#.#...#
..#.......
.#.#.#.#.#
#........#
#..#.#####

Tile 1753:
###.##.##.
#..#.#..#.
.....#...#
#.........
##.#..##.#
.#.#.#..#.
#......#.#
...#..#..#
#.#..#..#.
#.#..#.#.#

Tile 1609:
..#.####.#
#........#
..#.##.#.#
......#.#.
.....#....
#.......#.
..........
#.........
#..##.....
##...#.#.#

Tile 1237:
.#.##.##.#
...#.##...
..#......#
##...#..##
...#..#.##
.#.#####..
...#.##..#
#....#....
...#..#..#
#..#...##.

Tile 2267:
.#..#..###
...#.#...#
..#..#..##
##....###.
......#..#
..#.......
....#.#.#.
....##...#
###.....##
..##...#.#

Tile 1777:
...#...###
.......#..
##.#..#..#
.........#
#....##.#.
#.#..##.#.
.#..##.#.#
....#.#.##
.###..#.##
......##.#

Tile 1993:
..#.#..##.
#..#.#.#..
.#.#...##.
.......#.#
###.#.#...
#..#....#.
#...##...#
#...##...#
......##..
##.#.###.#

Tile 3571:
.###.#...#
...##.....
##..#.....
..#...#..#
......#.##
..#.##...#
##.#..#..#
#.###...##
..####....
..#.#..#..

Tile 3581:
#.#.#.#.#.
#.#.......
#.#....#..
#.###...#.
#.........
.........#
.#.......#
##........
#.##.....#
##.#..##.#

Tile 2383:
#..##..#..
..#...#..#
#...#.#.##
#.##.....#
....#..#..
...#.##..#
...#.....#
..###.#...
....##..#.
#...###..#

Tile 2591:
##...##.##
..#..#.#..
...#.....#
#.###.....
.....####.
...###.#..
..#....#..
#.........
....##...#
###.###.##

Tile 1153:
##.#...#.#
###......#
#....#.#..
....#...##
#...#...#.
..##....##
#..#...#..
#........#
.#..#.####
..##..#...

Tile 1489:
#.#.###..#
##.#.#..##
.#..##..#.
....#....#
#.#..#.#..
.#.......#
#......#..
##.....#.#
....##.##.
#....#####

Tile 1193:
##..###...
##.#......
......#..#
##......#.
#....#####
.#.#....##
...#..#..#
#..#......
##....##.#
########..

Tile 1163:
##.....###
.#.......#
#...#.##..
#..#...#..
...##.#...
.........#
....#...##
.####.#...
#.#.......
.#..##.###

Tile 2137:
##..#..#..
.#.#...#.#
#........#
#.........
..#.......
#...#....#
..##.....#
#....#.#..
..#......#
..##.##...

Tile 1559:
#..##..#.#
.........#
#..##.##..
#........#
#..#..#..#
....#...##
...#.....#
#......#..
.#........
###..####.

Tile 2521:
.###..#...
#..#......
....#.#.#.
#.....#.##
.#..#...#.
.........#
##...#...#
#...#.....
#..##.#.#.
##..######

Tile 1171:
..#####.##
..##.....#
.....#....
#.#.....##
#.#...#..#
#.##..#..#
#####....#
#..#.##...
##.......#
#.#.#.....

Tile 1361:
..##.###.#
.....##..#
###....#..
..........
#.......##
#..#.....#
#.........
.#.....#..
.......#.#
..##..#...

Tile 1259:
#.##..####
#...#..#..
..###....#
....#....#
...###....
...#..#.##
#.#.##....
##..##...#
##....#..#
#.###..##.

Tile 1987:
##.###..#.
.....#.#.#
#......#.#
#......#.#
..#.#.....
.#..#....#
.####.##..
..##.###.#
#........#
.....#....

Tile 2777:
#.#.##..#.
#...#.#..#
..#......#
...##.....
#.#......#
#........#
#....#....
.......##.
.....#..##
#..##.#..#

Tile 3673:
...###.###
###..#...#
..........
..#.#.##..
#.#..#....
#......#.#
..##......
.#...#.#..
#......#.#
##..#..##.

Tile 2819:
#.##.#####
#.....#.#.
####...#.#
.....#...#
#...#.#.#.
#..#...###
......#..#
....#.#...
.....#.#..
..#######.

Tile 1699:
###....#.#
......#..#
#.........
......#..#
#...#.##..
...#.....#
.........#
...##..#.#
#.........
.#.....##.

Tile 2351:
.#####.#.#
##....##..
###...#.#.
..#.#....#
.........#
.#.#.....#
..#...##..
#..#.#....
........##
###.#...##

Tile 1399:
.###...#..
#.#......#
.##....#.#
...#..#...
#.#.#..##.
#.......#.
#....#..##
....#...#.
#.....#..#
..#..#.#.#

Tile 3793:
##..####.#
.##...#..#
#.....#..#
#...#.....
.....#...#
..........
#....#...#
...##..#..
.#..##....
....#.#.##

Tile 3701:
.....##.#.
#..#.##...
........##
.........#
#..##...#.
#...#.#..#
..#..#.#..
#.....####
#...#..##.
...#..###.

Tile 2707:
#.###.#.#.
#......#..
####...#..
..........
#.#..##..#
#####..#.#
#..#..#.#.
###......#
#..#..#.#.
....#.#...

Tile 3559:
##.#.##.##
#....#...#
...##..#.#
..#.##....
.....#....
..##..#.##
....#.....
......##..
...###....
#.##...#..

Tile 2381:
..#...#.##
..#...#.#.
###...#...
##..##....
....#.##.#
...#......
.#..#....#
..........
##....#..#
###.#.##.#

Tile 1973:
.#...##...
#.#.#...##
#....#..##
#......###
......#..#
.........#
...##..###
.##..#..#.
#.....##.#
...##..###

Tile 3607:
#..#####..
.#.###.#.#
...##...#.
..#......#
.#.#......
##.#....##
###....#..
#...#...##
.#.....#..
#####..#..

Tile 2251:
##..#..#..
#........#
.........#
...#...#..
#.......##
.#.#..#..#
.##.##...#
###...##..
#........#
#.#..#####

Tile 3947:
..#.##..##
###.......
#.#..#....
#.#......#
#.#...#...
##.#.##..#
....##....
...##.#...
#.#.....#.
.##.#####.

Tile 3259:
.#.#####..
.....#.#..
##........
.....#....
#.#.......
.........#
......#...
......#..#
.##.......
#..#.##.##

Tile 3761:
.#....####
.##..#...#
........#.
###.#.....
#.....#..#
..##....##
........##
...#..#..#
...#.#....
###.######

Tile 2473:
#..###...#
##.....##.
#..###....
##....##..
.....#.#..
#.......#.
..#.....#.
......##..
....#....#
#...#.#...

Tile 2767:
##.##.#.#.
.#.####...
..##....##
#...#.....
..#..#...#
#.........
#........#
#...###...
#.#...##.#
.###..##..

Tile 2281:
#...#..#.#
...#......
.........#
#....#....
.......#..
.....##..#
#.#....#.#
##..##...#
........##
#....#...#

Tile 1693:
#.##.#####
......#...
#.#...##..
....#.#.#.
#..#...###
##.##.##.#
..#...####
#..##....#
#......#.#
#....###.#

Tile 2417:
.###.##.#.
##...#...#
#..#..##.#
#...#..#..
.........#
...#.....#
..##.#..#.
#.........
....##....
#.##.###.#

Tile 2879:
..#.#..#..
..##..#.#.
##......##
#..##...##
#.##..#...
#.......#.
.......#.#
##..#..#.#
#...#....#
#....#.##.

Tile 2851:
.......#.#
....#.....
##.###.##.
.......#..
.#..###..#
....###..#
#..#......
...#.#...#
.#.###...#
###.###.##

Tile 3931:
.....##.##
###....#..
#..#....##
#.#.....##
##...#....
#..##..#.#
.......#.#
...#..#.##
.#...#....
##....#.##

Tile 3181:
...##...#.
....#....#
#......#.#
#.###....#
##.......#
.#..#.#..#
###..#.#..
#....#.#..
#.......##
###..###..

Tile 2647:
#..#..#.##
#........#
#..#..##..
.....#...#
....##.#.#
#..###...#
..#......#
#.#..###..
#.......##
####...#..

Tile 2441:
..#..###..
##..###..#
##.##...#.
#.##.....#
.......###
...#...###
#...#.....
.#..####.#
.#.#.#..#.
#.#.##.#..

Tile 3733:
...#....##
..#..##..#
.....#..#.
##........
..#....#..
#....##.##
.#....#..#
.#.##.##..
#.##..#.#.
.#..####.#

Tile 1607:
#.#.#.#.##
#.........
.....#....
.#.#..#..#
..#......#
#.....#...
#.#..#....
......#.##
#.##...#..
..#.######

Tile 3301:
....##...#
#..##....#
....#.....
#.....###.
#....##.#.
#.#...#..#
......##.#
...##.....
.........#
..#######.

Tile 1109:
..#...###.
.#.......#
#.....#...
..........
##.#.#....
#...##....
#.....###.
.#....####
.#....#.##
###.#...#.

Tile 2789:
.#.#.####.
..........
#...#.#.#.
#..#....##
.....##...
.#..##.#.#
........##
#.#.#.#...
......#.##
#...#..#..

Tile 2593:
...#.#.#.#
#......###
..#.##....
#.......#.
...#.#...#
#..##....#
#..##.....
#.##.#...#
#..###..#.
.#....###.

Tile 2539:
...###..#.
..........
.........#
.....###..
.......#..
......#.#.
.#..#...##
#....#.#..
##...#....
.###.#.#.#

Tile 3307:
##.##.#...
.#..#.....
#..#..#..#
#.....##.#
..#.#.....
##....##..
.....#....
#..#.#....
######...#
####.#....

Tile 1613:
###..###.#
...#...#..
#.#..##...
......##.#
......#..#
..##......
#.....#..#
##.#..##.#
###...#..#
#.#..##.#.

Tile 3727:
..#...#.##
#.#......#
#.###...##
..####..#.
#.........
.....#...#
#..##....#
#.#.....##
#....###.#
..#..##.##

Tile 2017:
#...##..##
#..#..##.#
#........#
.....#..##
#..#.....#
....#....#
#....#....
.##.......
.......#..
#..##..#..

Tile 1933:
###.#.....
#.#.####..
..#..##...
##..###..#
###..#.###
..#.#....#
#.........
..#..#.###
#.#..###..
...#.#.#..

Tile 2239:
.#.#..#.##
...#.##..#
#...#....#
#..#......
#....#..#.
#..##..#.#
..#.#.#..#
#....#..#.
#...#.#...
.#.##.####

Tile 3449:
.##...###.
..#......#
.......#.#
..##......
#.###.##.#
....##.#..
........#.
##.#......
...#.#....
..###....#

Tile 2347:
..##....#.
#.........
..........
....#.#..#
.........#
##.##.#..#
......#..#
.#.##...#.
#.#####.##
.#.#....##

Tile 3023:
#...#.##..
.......#..
#.#..#.##.
#...#.#..#
..#.#...##
#.##...##.
......##..
.....#####
.#...##...
#..#.#....

Tile 3541:
###.######
..#.....#.
#....###..
......#...
..#..##..#
#........#
#..#......
###..#.##.
..#....#.#
#..###.#.#

Tile 3709:
##.#.##..#
#.........
.##.##...#
.#..#.#..#
#.......##
......#...
.#..#.....
.......#.#
#..##....#
.##.......

Tile 1657:
#.##.#...#
...#.#..#.
.....#..##
.....#.#..
######.###
##........
......#..#
#..#...#..
.......#..
...#....#.

Tile 1031:
####.#.##.
..........
.#...#..#.
...#.....#
.##.#....#
....#.##..
........##
#..#..#.##
#.##.....#
#.####...#

Tile 3691:
.#.#######
........##
#.#.......
#..##..#..
#.#.##.##.
....###.#.
.#..#..#..
..#.##.#..
.....#....
##.###...#`

/*
  input = `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`
*/
tiles := []Tile {}

  for _, tile := range strings.Split(input, "\n\n") {
    tiles = append(tiles, NewTile(tile))
  }

  p1 := P1(tiles)
  p2 := P2(tiles)

  fmt.Printf("P1: %d, P2: %d\n", p1, p2)
}

type Tile struct {
  id int
  t string
  r string
  b string
  l string
  data [][]string
}

func Reverse(s string) string {
  r := ""
  for _, l := range s {
    r = string(l) + r
  }
  return r
}

func NewTile(spec string) Tile {
  tile := Tile {}

  nameMatcher := regexp.MustCompile("Tile ([0-9]+):")
  nameMatch := nameMatcher.FindStringSubmatch(spec)

  v, _ := strconv.Atoi(nameMatch[1])
  tile.id = v

  lastString := ""

  allLines := strings.Split(spec, "\n")
  allLines = allLines[1:]

  for y, line := range allLines {
    if y == 0 {
      tile.t = line
    }

    if tile.data == nil {
      tile.data = make([][]string, len(line) - 2)
      for idx, _ := range tile.data {
        tile.data[idx] = make([]string, len(line) - 2)
      }
    }

    lastString = line

    for x, val := range line {
      if x == 0 {
        tile.l += string(val)
      } else if x == len(line) - 1 {
        tile.r += string(val)
      }

      if (y > 0 && x > 0 && y < len(allLines) - 1 && x < len(line) - 1) {
        tile.data[y - 1][x - 1] = string(val)
      }
    }
  }

  tile.b = lastString

  return tile
}

func P1 (tiles []Tile) int {
  ret := 1

  borderCounts := make(map[string]int)

  for _, tile := range tiles {
    borderCounts[tile.t]++
    borderCounts[tile.r]++
    borderCounts[tile.b]++
    borderCounts[tile.l]++

    borderCounts[Reverse(tile.t)]++
    borderCounts[Reverse(tile.r)]++
    borderCounts[Reverse(tile.b)]++
    borderCounts[Reverse(tile.l)]++
  }

  for _, tile := range tiles {
    borders := borderCounts[tile.t] +
               borderCounts[tile.r] +
               borderCounts[tile.b] +
               borderCounts[tile.l] +
               borderCounts[Reverse(tile.t)] +
               borderCounts[Reverse(tile.r)] +
               borderCounts[Reverse(tile.b)] +
               borderCounts[Reverse(tile.l)]
    if borders == 12 {
      ret *= tile.id
    }
  }

  return ret
}

func P2 (tiles []Tile) int {
  ret := 0

  borderCounts := make(map[string]int)

  for _, tile := range tiles {
    borderCounts[tile.t]++
    borderCounts[tile.r]++
    borderCounts[tile.b]++
    borderCounts[tile.l]++

    borderCounts[Reverse(tile.t)]++
    borderCounts[Reverse(tile.r)]++
    borderCounts[Reverse(tile.b)]++
    borderCounts[Reverse(tile.l)]++
  }

  var corner Tile
  cornerIdx := 0
  leftBorder := ""
  topBorder := ""

  for idx, tile := range tiles {
    borders := borderCounts[tile.t] +
               borderCounts[tile.r] +
               borderCounts[tile.b] +
               borderCounts[tile.l]
    if borders == 6 {
      if borderCounts[tile.t] == 1 {
        if leftBorder == "" {
          leftBorder = tile.t
        } else {
          topBorder = tile.t
        }
      }
      if borderCounts[tile.r] == 1 {
        if leftBorder == "" {
          leftBorder = tile.r
        } else {
          topBorder = tile.r
        }
      }
      if borderCounts[tile.b] == 1 {
        if leftBorder == "" {
          leftBorder = tile.b
        } else {
          topBorder = tile.b
        }
      }
      if borderCounts[tile.l] == 1 {
        if leftBorder == "" {
          leftBorder = tile.l
        } else {
          topBorder = tile.l
        }
      }

      corner = TryTile(tile, leftBorder, topBorder)
      if corner.id == 0 {
        leftBorder = ""
        topBorder = ""
        continue
      } else {
        cornerIdx = idx
        break
      }
    }
  }

  grid := ArrangeTiles(tiles, cornerIdx, leftBorder, topBorder)

  // Put the data into the big picture
  fullPicture := make([][]string, len(grid) * len(grid[0][0].data))
  for i, _ := range fullPicture {
    fullPicture[i] = make([]string, len(fullPicture))
  }

  for y := 0; y < len(grid); y++ {
    for x := 0; x < len(grid[y]); x++ {
      for dY := 0; dY < len(grid[y][x].data); dY++ {
        for dX := 0; dX < len(grid[y][x].data[dY]); dX++ {
          fpY := y * len(grid[0][0].data) + dY
          fpX := x * len(grid[0][0].data) + dX

          fullPicture[fpY][fpX] = grid[y][x].data[dY][dX]
        }
      }
    }
  }

  return MonsterCount(fullPicture)
}

func MonsterCount([][]string picture) int {
  head := regexp.MustCompile("..................#.")
  body := regexp.MustCompile("#....##....##....###")
  head := regexp.MustCompile(".#..#..#..#..#..#...")
  //TODO
  //
  return 1
}

func ArrangeTiles(tiles      []Tile,
                  cornerIdx  int,
                  leftBorder string,
                  topBorder  string) [][]Tile {
  gridSize := int(math.Sqrt(float64(len(tiles))))

  // Start with an empty grid
  grid := make([][]Tile, gridSize)
  for i := 0; i < gridSize; i++ {
    grid[i] = make([]Tile, gridSize)
  }

  // Rotate the corner tile so that it fits
  cornerTile := TryTile(tiles[cornerIdx], leftBorder, topBorder)

  // Put in the top left tile
  grid[0][0] = cornerTile

  // Remove the tile from the list
  tiles = Remove(tiles, cornerIdx)

  topBorder    = ""
  leftBorder   = ""

  // Put in all the other tiles
  for y := 0; y < gridSize; y++ {
    for x := 0; x < gridSize; x++ {
      if x > 0 {
        leftBorder = grid[y][x - 1].r
      }

      if y > 0 {
        topBorder = grid[y - 1][x].b
      }

      if x == 0 && y == 0 {
        // We've put in the first
        continue
      }

      for i, tile := range tiles {
        tile = TryTile(tile, topBorder, leftBorder)
        if tile.id != 0 {
          // The tile fits
          grid[y][x] = tile
          tiles = Remove(tiles, i)
          break
        }
      }

      leftBorder = ""
      topBorder = ""
    }
  }

  return grid
}

func Remove(tiles []Tile, idx int) []Tile {
  tiles[len(tiles) - 1], tiles[idx] = tiles[idx], tiles[len(tiles) -1]
  return tiles[:len(tiles) - 1]
}

func TryTile(tile Tile, top string, left string) Tile {
  if TileMatch(tile, top, left) {
    return tile
  }

  for flipV := 0; flipV <= 1; flipV++ {
    if flipV == 1 {
      tile = FlipV(tile)
    }
    for flipH := 0; flipH <= 1; flipH++ {
      if flipH == 1 {
        tile = FlipH(tile)
      }
      for rotate := 0; rotate < 4; rotate++ {
        tile = Rotate(tile)
        if TileMatch(tile, top, left) {
          return tile
        }
      }

      if flipH == 1 {
        // Put the tile's data back (slice is a reference type)
        tile = FlipH(tile)
      }
    }

    if flipV == 1 {
      // Put the tile's data back (slice is a reference type)
      tile = FlipV(tile)
    }
  }

  return Tile {}
}

func FlipV(tile Tile) Tile {
  tmp := tile.t
  tile.t = tile.b
  tile.b = tmp
  tile.l = Reverse(tile.l)
  tile.r = Reverse(tile.r)

  // Flip the data as well
  for i := 0; i < len(tile.data) / 2; i++ {
    opp := len(tile.data) - i - 1
    tile.data[i], tile.data[opp] = tile.data[opp], tile.data[i]
  }

  return tile
}

func FlipH(tile Tile) Tile {
  tmp := tile.l
  tile.l = tile.r
  tile.r = tmp
  tile.t = Reverse(tile.t)
  tile.b = Reverse(tile.b)

  // Flip the data as well
  for y := 0; y < len(tile.data); y++ {
    for x := 0; x < len(tile.data[y]) / 2; x++ {
      opp := len(tile.data[y]) - x - 1
      tile.data[y][x], tile.data[y][opp] = tile.data[y][opp], tile.data[y][x]
    }
  }

  return tile
}

func Rotate(tile Tile) Tile {
  tmp := tile.t
  tile.t = Reverse(tile.l)
  tile.l = tile.b
  tile.b = Reverse(tile.r)
  tile.r = tmp

  // Rotate the data as well
  for y := 0; y < len(tile.data); y++ {
    for x := y; x < len(tile.data[y]); x++ {
      tile.data[y][x], tile.data[x][y] = tile.data[x][y], tile.data[y][x]
    }
  }
  for y := 0; y < len(tile.data); y++ {
    for x := 0; x < len(tile.data[y]) / 2; x++ {
      opp := len(tile.data[y]) - x - 1
      tile.data[y][x], tile.data[y][opp] = tile.data[y][opp], tile.data[y][x]
    }
  }

  return tile
}

func TileMatch(tile Tile, top string, left string) bool {
  return (top == "" || tile.t == top) &&
         (left == "" || tile.l == left)
}

func PrintTile (tile Tile) {
  fmt.Println("Tile: ", tile.id)
  for _, line := range tile.data {
    fmt.Println(line)
  }
  fmt.Println("top: ", tile.t)
  fmt.Println("rig: ", tile.r)
  fmt.Println("bot: ", tile.b)
  fmt.Println("lef: ", tile.l)
}
