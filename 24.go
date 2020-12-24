package main

import (
  "fmt"
  "regexp"
)

func main() {
  input := []string {
"neneeenenesweneswneesewenenwneew",
"eswswenweeseneseswwsenewsenw",
"seeweswwneswnwwwenwwwwwnenwwsesw",
"swsweseseseswswsenwswseswswsese",
"eseseseswsenweseseneseswnwneeewww",
"sweeeeeenweseenesweee",
"nwnwsewenwwnwnwsenwwnwseswnwnw",
"swswswwwswweswswswswswenenesesewwwe",
"swnwenenesenwsewswswseeeswnwwnewwsese",
"senwweewnwweswnenwwsw",
"ewneswswsweenwwnwsenwswnewsenwnwsw",
"nenenwwneenwnwnwwnesewsenenwnwnenwnw",
"nwnwwsenwwnwnwnwnwwnwnwewwnw",
"wnwnwnwnwseswwnwnwesewsenwnwnwnwnenew",
"seswewwneweswnwnwswwsw",
"neswesenenenewwswnenwnenenenwnenenenese",
"nwwswsenwnenwnesenwnwsenwnwnwwnwsenenwse",
"senwswwswwnweswswneswswwswnweswseswew",
"neeeeswnweswseeeeeneeeeenwwnene",
"nwnewnwnwnwswnewwsewnwwseenenwnww",
"eeseeeeeswneenwewee",
"neseseseeenwwseseeseneswnenwsewswse",
"wswswnwnwwswswneswwweeswswswwww",
"swnwnwneswswseswseswswswswswnwseswsesenesw",
"swneesenwnwnwwnwnwnenwnwnwnwnewnwsese",
"wswswswewnewnewsenewnwswseswswwse",
"swenwnwnwnenwnenwnwswnenenwswwnenwsenw",
"ewnenwnwweenwneesewnwwwwnwswse",
"eeewnweneneeesweesweeeneeee",
"nenwnwnwnwswwnwsewnwswswenwweswenw",
"senenenenenenwneneewnenenene",
"neneeneswnesewneneneeneneeenenenwe",
"nwswenwseswnwswenwnwnewnwneww",
"wnewwwweswwsww",
"seseswsewesesenesesewseewseseenenwe",
"neneseswnesenewwenenwneneneneneenew",
"swswswswwsweswseswsewwneswswswswneswswsw",
"sesenesesesenwsesesesenwsewnw",
"nenenenenenewnwswnwnenwsenesenwneswwne",
"swwnewweewwwwnwswnewwweswew",
"swnwwswesenewwwswwswwweseswswnesw",
"weeseseseneseseewwseseseseseseesese",
"nwenenenwwnenenenenwneneneseswnenwenesw",
"nwsewswswwswsweswnwwwnwswswsew",
"nwewnwwwewnwwnwnwwnwswwnwwsewnw",
"eesesesesesesewseseewneseesesesesenw",
"nwwenwweswnwswenwnw",
"sewseswswneneswnwenwneeswwnwwesenw",
"eseswwnwnwwenewneswnwsweenwseswne",
"wsewwswnwnwewnewswwwwswseswnew",
"eswnenwnwnenwswnenwnwnwnenwnwnwnenwsewne",
"neneewswnenenenenene",
"nwseeseseesesesesesesese",
"eenweneeeswswseeseeneesese",
"sewneeneneswseeewnwswsenewswenwwenw",
"nwwwweswnenwsenwwswnwwnwwwnwwnew",
"nwewwneswnwwwnwswswsesweswswwseww",
"neseswseweseeswsenwsesesesesw",
"nwswswseswswseeswswseswnewneseswswwsw",
"eweseeswseeeeseneenwenweeee",
"nesewwwwneeseweswwseneeseseswseswsw",
"nwswnwnewnewneeenwneeswwnwse",
"swwnwsesewneneeneenwseswseswnww",
"nwnwnwnwsenwnwnwnwnwnwnwnw",
"wwnewnwsesewwwwwwnwseewsenenew",
"swswswswenwswseswsweeswnwswswswswswswswnw",
"nwnwseeswnwnwswnenwwnwnewnwnwwnwnw",
"wnenenesesenwnwwnwsenenesewnewnenee",
"senwswnwnwwenenwnenwnenesewswenewnene",
"neseseeseneswswswswswwwswswnwswwsenesw",
"neswwsesenwnwnwneswseseseseswnesenwsese",
"eeweesenesweenweeneenwneeswse",
"wwnewswsenwnwsesenwswseneeswswseswnese",
"wnwsewnwnwnenwnwnwnwwnwnwwww",
"wwswswewwwwwwswwwwnewnwnenw",
"eeneseswswnewneneneenenenenenenenew",
"wswwnwswwswneneneswnwnwnwswseneneswsenene",
"nwswswewswsewswsewswsenweswswneswnwsw",
"wwwwesewnwnewwwnewwwwsese",
"nwnesenwnwnwnwswswnwwswwnwnwwwwneeenw",
"swseseswenwswswnwswwnenesenesesesesesese",
"nesewswswswsewswswnwswwswwswsw",
"wnwwwnwnwnwnwsewsenwsewwwwwww",
"wwwwwwwwswwswwwewwnw",
"swseseseseeswsesesenwseseswseswwsenwse",
"sweeeseeswnwenwneeenwneswenwneee",
"nwswseswswswswwsweswnesesweswswswseswse",
"nwnwenwnwnwnwnwswnwwnwnwwnenw",
"wseewnwnwnwenewnenenesenewnenwnenwne",
"nwsweswswswswswswswswsw",
"swswwswwswswswswnenewenwewweswnw",
"nwseeswnwnwseseseswneseswseseseenewsw",
"wseswseseesenweeeseeeewnwenene",
"newnenenwswnwnenenwneneseneneneenenesenw",
"wnwnwswseenwnwnwnwenwwwnwnenwnwswne",
"esesesesewneseswswsewswseswswneswsesesw",
"wsewwwnewswwnenenwnwwnwwnwnwwwsew",
"wsenwenwnwnwewwsesenwwwwnwnwnwnew",
"nwnwnenesenenenwnenewnenenwneneswswnene",
"eseneneewneneneeneeneneee",
"nenenenwnwseswswnwwnwnwwesenwsenewswne",
"neenwnwnenenwnwnwwnwnenwnwnwwnwswsenw",
"swwwswwswswwnwwwnenwseswswwseswsenw",
"nwsewneswnwesweeswnwwwnweswnwne",
"neswsewswseswseswswswneswseswneseseswsww",
"wsewwwnwnenwnwewwseenwwwwswnw",
"neseseswseswwswswswnenwsesewswswswswsw",
"swwsweswneseswswwswswswswweswwsw",
"nenwneneswneeneenene",
"eneweneseneeneeswswswenenesenenew",
"swseeseswswesenwsesenwswnwnwseeesenwswsw",
"nenenwnenwwwnenenweneneneseneswnenenene",
"neneenenwswswseswwwnweswnenesesewnewse",
"neneseneneseneenenenewswneeenenewnene",
"swsweesenesenenenwnwnwswsweneseenww",
"swnewwnwsewseswewsewsesenwneewnww",
"enenewnenenenwnenenenw",
"seseesesewwseseseseseswswseesesenese",
"eeenwneeeewneneeeeewswneneee",
"eswnwenweesenwseeswsweenwsenwnwesw",
"eesweeeneneeneeeeewnweeewe",
"seswswseseswswswsene",
"newswnewswwwneswwwse",
"neeswnenwewnwnesewsweneeswnwenesww",
"eswwwwwwswnesewwwnewwnwwww",
"ewsesenwseswsesenwsesesewesenwsenenese",
"nenwnewswneneneneeeneneseneneneenee",
"wwseseeeenwnwweseneewnweeeewsw",
"nwneswnwnwnwwwnwseenwnwnwenenwnw",
"swswswswswneswenwwweswswsw",
"swenwwwnwnewswnenwneswnwesweswwe",
"neswseenwseseeenwsesesew",
"seseeesesenwnwseeseseseswnesese",
"sweeeseseenwseeweesee",
"swneswsweswnewwswseswswsw",
"nwwwewsewwwwwwwwewswswnewsw",
"nenenewneswsenenenenesenenenenenenenwneenw",
"sewwwnesewswwnwwwwswewswwswew",
"swswseswnwnwswswswswswswnwseswenwseswswsw",
"nwnwnwnwnwnwnwnwnenwnwsenw",
"nenwwewnwwsewnewwnwsewwwewseswe",
"seseseswweewnwneneseseseesewsesesw",
"swswesenewswswswwewswsesweswsw",
"seseneneswwsewneeenwenenwseenwsee",
"nenwenwwneewnwneswnwnenwseswnwenwnwnw",
"sweswswswseswswswseseswseswwnesesesenw",
"nweeesweeeeneseeseeenwenwwe",
"sweewneneeeseeewswneeenwneenene",
"swsewseseenwsenesewseseseswswseswnese",
"wwwswseseswwnwwnwnenenwwswnwsenesene",
"seswswswseeseswseseseswnwsenwseswswsesene",
"nenwesenwswneneseeewsesesweeseeswse",
"nwseswswnwseeswnwswsenwnwnenwnenwnwnenwnw",
"neswnenwweesenwneswwneneswnwenenesw",
"nwneswnesesenwswenwsesewewsweewsw",
"neseeenenenewnenenweswew",
"eenweesweewneneseneeswee",
"weesenwsweenenenwswenenwneswenee",
"nwswswsweseswwsweswswswswnwswswswswsw",
"esweweseeeeseneeenweeeeee",
"swneswswswwswnenwewswenwswswwnew",
"eesweswnwseeeeeeeeneseeenwsee",
"senwwswnwnwnwnwnwnenwnwnwnwnwwnenenwnese",
"swwseeswswnwnenesewnenwsenweeeene",
"enenwnesenenenenwnenewneneneewswswwne",
"nwwneseseeseeseeseseeseseeseswnwsese",
"sewwswsewnwwswnwnwnewenweneswesw",
"eneeneeneeneewneeeeeeenwsesw",
"wwenwwnesenwnwnwwww",
"nwneneswnesenenenenwnwnenenenwnene",
"sweneseseswewnwswwneneneneseneesenwnw",
"nwewenwnwwwnwnenwsesenwnenesenwnwwwnw",
"nwenwnwswsenwnwnwnenwnwnwnwnwwesenenw",
"swswwnewweweswswnwswsenesewwwnww",
"nenewneswnenenweenenenwswsenene",
"swswswwswswswswswswswswswne",
"eeswswenwswneenwenweseenwnweeese",
"swnwnwnwwwsewnwnenwwwnww",
"wswwwenenwwwweswseswnwswweswsw",
"enewneneweseeswswenewnenwwnesene",
"swsenesweswswswswwswswswseseseswwsw",
"nwwseswnwsenwsenenwwnwnwneewwwwww",
"wwsesesenwnwnenesenwenenwswneneenesw",
"seseneweseeswseseneeseseesenwesese",
"nwnwwswwwweenwwnwwswsewwneww",
"eneneneeneswnenweneswenewneeseenenew",
"eseewneesenewesenwsewewwwsee",
"nwnwnwenenenwenwnwnwsenwnwnwswnewnenwne",
"neseeeesewneseseweeseseswseseesese",
"newwnwswwewneswseewsesweswneswwse",
"seseseseseeseenwseeenwsesesesese",
"eseswnesewseswswswsenwwswseswneswsee",
"eeseseseneeeeeseeeseeew",
"nwesweeneenwnwseneeeneeeneneneswse",
"swswswswseseswnwneswswseswsewse",
"senwnweswwnwwnwweseenwswsenwnwwse",
"wneneneenenesenewnwswswenwnene",
"eswseswneseseneseseeeeeewsenenwnew",
"sewseeeeenwseese",
"sesenwnwseseseswseseswswswseesenesesesenwse",
"nwwneewseswneeeneneneseenenwewneese",
"wwwswwnewwneswwwwwsewwwsw",
"seswwwwneswneswwneswswswswswseww",
"eewneenwsweseseweeeeneneneneesw",
"swswsewswswwswswnewwswsww",
"nwnwwswnenwwnwnenwswnwnwwsenwseeswwnw",
"swswsenwnewseeswseseneneswsewnwsesesenwsw",
"senwnwnwnwswneswnwenwnwwnesweswswnwnenw",
"neenenewswnwswnwsewwewswsenwnwwsw",
"senwnewnewnwwnwwnwseseswwnw",
"seeseseeneswneeswnenwwswswwswnwswwe",
"seswenweeneeneneneeneneneee",
"nwwnwwenwwwwwwwnw",
"swswswswwswwswwswsenenwsewnewneswswne",
"newswnwnwswnwenwwnenwnwnwwwwsewnwnw",
"eneewenweeeweeseenwseneeeswse",
"seseseswenwseseweseseenwnee",
"wnwwnenwswnwnwweeeswsw",
"senesesenwswseseseenwseswseswswseswsesw",
"seseseseseseseseseseseseenwsese",
"seenwnwnwswsewweeweseeneesewswene",
"wwwseswsesenwnwnenenwse",
"nwwewnwnwwneswnwwnwnwnwwnwewswne",
"neewneneseneneswnene",
"wswwswnwwwswwwewswswsww",
"wswwneswwswnwwnwnewseswsewswenwsw",
"ewnwwwwswwnwesweswswwsw",
"wswnwnwwnenwnewwsesewwseenewswenew",
"wswwnenwwswnenenwnenenenenweswseseswe",
"nesesenwswsweeswewwswneeneeeee",
"nwswwwnewwwwwwneewnwsewsewsew",
"nwwweswwswwwnwneswnwswenwnwneew",
"nwnwswnwswenenenwnwwnwnenwnenenweew",
"neswnwneseeneneenewswwnwnweseneene",
"nwswswwwwseeswnwswwsweswwnwseswww",
"neneseneswsesewenwseswneneswnenwwenew",
"neneneeneneneseneeneseneww",
"sweenenwenweeenesweeeenwesee",
"seseseseseseneesesesesesesesew",
"senesweewneeswenwseswseseeeesee",
"nesesenwwswnwnwwnewsenenwswnwnenwsee",
"neseewnwenenwwnewnenwnwsenwnwswswnwnw",
"eswnenenwswnewsweseswseewsesenwsee",
"enweeswenwewsweeeneenweeseesee",
"senwswswwswwswswwswswswswseswswswneswne",
"sweswwnwewseneeeenenewenweee",
"newnenwnwswwseseseesenwsese",
"enweeswseenwneeenenesweeneswenewsw",
"nwnwnwnwnwswenwnwnwnwnwnweswnwnwnwnwnw",
"nenenenweneswnesenwneswneneneneeneese",
"seseeenwswnweeeewswseeeswwnwsenw",
"esenwsewseseesewneseeseseseeeenese",
"seewnwsesesenwesesesesewneeeseee",
"eneseneewneneseweeneneeeneeee",
"nwwnwenwswsenwnwnwnwnwwnwswwnwnwnwnew",
"wnewwwwwwnwwwswnwww",
"wnwenwsenwswnwnwnwwnenwnwnw",
"eeneneswnwwneneswnwneeeneneneneewne",
"wsewnwwwwewwwwwwwsenwwnwwe",
"wwwseeseswsewsewewneswnwsewnenenw",
"seseneswsewneswswseswseswswswseneswsesesw",
"nwnwnwsenenweeneneeneswswneswneseeesw",
"eswseswwnwswseswswswswneenwswseswsese",
"eenewseneeeseneneeswwneeeenenee",
"nweeeeseeeseeseeswenwnenweswee",
"swswwneneneswenenwsenenenenenwnenwnwe",
"nwenwnwnwseswnwnwnwnwnwnwsenwnwnwnwnwnwnw",
"eswnwnweswsweseseswwswswnwseneswnenwswe",
"wnesewneswwneswwwwnwwwwsese",
"nwnenwnenenenenenenwneswnenwnenw",
"neswenenwswneneswnwee",
"wwnwwwwewwnwswnesewww",
"esenwnwswenenwnewswneswwenwswwnwswnww",
"swswwswnenwwswswnwseswswswwswswewsesw",
"nenwsenewnwswswneeswnwnenwnwnweenwnwne",
"neswwnwwwnwwwnenenwnwswwnwsenwnww",
"neneeneswenenenenenenee",
"swseweseneeeswwnenwswnwnwe",
"eewswweneweeneneeenenenenenenene",
"nwnwneneneneseneseswnenwnenwnwnenwnenene",
"enwnwenenwswswswnwsenwswswseswseenwe",
"eneneneneeeswneneenenenene",
"enwnenwwwsenwwwswneswnwwenwewww",
"neneneswnenewneneneeenwnenwnesewnene",
"seswsenwsesesenwewseswseseseseswsesenenwse",
"neneenewsesewneweeewseseenwneeene",
"nwwswwwsesenwwwsenewwnwnew",
"seseenesesesesesesesewesesee",
"nwwnwswswswseswwswseneswswswswswswswswe",
"wswwnwnwnwnwnwnwwnwnenwnwenwnwenwnw",
"seseswseneeswneswnewwswwwnwneesenene",
"nenwnwnwnenwnwnwnenwsenwnwnwnw",
"eswwwwwsewwnewswswwwswswneswww",
"esenwnwnweeewswneneswnwswnwswnenwnwnene",
"nenweseenwneenenwseeeneseene",
"newnwwsenwnwenwwwnwsewswnwnwenwww",
"newenenenenenenwenwswswswsene",
"nwsweswnwwnwewewwwenwwwwwnwwnw",
"sewnwsewwnwsenenwwww",
"swswnenwnwnewneswnewnweswsweswenwsenwse",
"wsweseeswnwseenewseneenee",
"eeeeeseeswnweeeenwseseewsese",
"sesesesenewsesesewesesesenwseseseene",
"sesewswseswsewseseseswweseeseseenese",
"esenwswwseseseseesesesenweseseseswe",
"nwnenenwnenwnwnenenesw",
"nwnenweneenwnenwnwnwswnenwneneswnenwnw",
"sewswswswswswnwswneeseswswswsw",
"swwswswswwwwneswneswwswswswnweswsw",
"neseswswesesewswwswswseswswnwseswnenese",
"nwswnwneewneweeseeseewewneneew",
"swseenenwwsenwnwnewneenwwewnwswe",
"swsweseswswswwwwswswswnewswsw",
"swnwwneneswswnwewwnwsenwenwnenwee",
"ewswseseneesewseneswsenwswseseswsese",
"eneeneeseswenwsewneeweenwnwese",
"swnwswswseswswswseseeswnwswswsweswnwsww",
"wwnwnwswnwsewwnwwnwwnwwnwsenwnee",
"eeseeeswneseneseeweneweseee",
"senesenwsesenwwsesesesesesesesesesesese",
"nwnewesewneswneeeeneeneeswnesenene",
"wesenwseswsweweenwnenwwnesese",
"nenwnwneeswneneeneneneswneneswneneenwnee",
"nwnwwsenwwnwnwwnwsenwnwenwnwnwwnwne",
"nwnwsenwnwsenwnwnwnenwnwwnesenwnwenwne",
"sewswweneesenenwneseswseswnweeeene",
"nwswnwnwnwnwwsenwnwwnwnwenwnwesenwnw",
"wwwwwwnenwwwwnesewwwse",
"swnenenenenenenwneneswnwneesenwnenewne",
"neneneswneneeneneneene",
"wswneseswswnewwnewswwwwwwwswsww",
"neeneseewnenwnwneseewneenenwwnenesww",
"enwseswseesesewnesweeeeneswsesenese",
"nwnwswnenenenwnwnwnenene",
"wswswswneswneswwswswsweseswswseswseswsw",
"sewseseesesweseneseseseseseswnwwsw",
"swseswswwnesenewnwwnweenweswwsenw",
"seseseswswseseseeseesesesesenwnwsesese",
"nenewnenwneswnwnese",
"seseneswseswswseseswswseswnwwswse",
"nwneeeeseseneseseeseeeeeeeeww",
"nwnwneenwswswnenwnenwnenwwneneenenwnwnw",
"nwnwenwnwnwsenwnwenwnwwnwnwnwswnwnwnw",
"wwnwwenesesewwwnewsesewnewwwe",
"wnwnwwwewnwseswwwwwnewnwswnwe",
"nwwwwswnesewwsewsenewnwnwsewnwnww",
"nweeeenwseeeseee",
"eneenewnwenewnwsesenewewswenenese",
"eswwsweswseswnwsenewnenwswseswwww",
"neswnesenenenwnenwenenwnwneneswneswnwne",
"enwseswnwnweesewnwseeseseeseseseseesw",
"wenwwwswnewswsesewwnwwnwnenewnw",
"nwwseseswswewwswwsewneswnenwnwesww",
"swnwnwnwswenenwnwnwnwsewenenenwnenesw",
"weneswwswswnwsweeesenenwsenwwwsw",
"eeseesenesesewnweseeenesweeeeew",
"eseweseseseesenwwsesenese",
"eneneneneneneneneeenenenenenwsw",
"esenwnweswenenwnwswnwwnewneenewnww",
"eeneneneeeeswneneenwsew",
"seweeseesenewswsenwewwneeeeee",
"newnwnwenwwnenesenewsenewswnweese",
"neswwwswswswneww",
"seswseswseneswseswswnewneswswswswsenenw",
"eneeneeeenwswneneeswnweenwseneswnene",
"neswseeswseseswswsewsese",
"wnewwswwneswswwenewswswwwswwnene",
"newswseswnwnwnwnenwnwnwnwnwnesenwne",
"neewswseswswnwswswseneswswswswswseswsesw",
"nwwneenwwwwsenwwseswenwnwswwnenw",
"wwwenwwnwnwwnwnwnwe",
"esesewsesesesesesenwsenwesenwswswswee",
"nwsesesenweseswseseseseswsesewseseenwse",
"newnwnwnenwswswwwnwwwnwwnwwnenwsw",
"swsewsenwseseseswwseseneesweswswsese",
"neeenweseeeeneswneeeeweee",
"nwenwnwnwnenwwnwnenenwnw",
"wwsewwenewwwwwwwsewwwnww",
"neneenesweneeneenwnenwneneseseeee",
"eswseneseseseeeeeeeseenwsewsese",
"nenenwswneeswnesenenewwnenenenwenwsese",
"nwwwwwewwswwswwnewwwwseeww",
"swswswswswswswseseswneeswswseneswwswwsww",
"nwneneswneeswenwseenweeneneewwwnee",
"swswswswneswwneseswwseswswswsw",
"swswseswneswwwswswwswsw",
"nwenweseesweeeeneeseeswseeesw",
"swswswsweswswnewwswwwswswwwnwswse",
"sewneseseswsesesesesewseseswnwseneswse",
"enenweeswsweswenwnwe",
"nenenwnwnenwnenwwnwswnwnenwsenenwne",
"neswseneswswnwwweneswnwseswswneewnese",
"nenweswsweeneeneswnweeneswneswee",
"wsenwseseseseeseseseseseseeseneswnese",
"swnenwwnwnwenwnwswnwnesenesesewnwnenenw",
"sweneeenenwenenenweeenenwsweeeswe",
"sewewnesesewnenwsesesesweseswswsesesw",
"neswswnenwneseseneneneneneneneswnenwnene",
"newnewweewwwnwseswwseswwewswww",
"senwseseseswswnesenweseswsewswsesesewsw",
"senenenenenenwnwnwsewnwnwenwnwnwwnene",
"nwnwenenwnwwneneswneneneneenwnwnwnese",
"seseseseseseswwneswseswesewswseswnwse",
"nenwnwnenwenenweswneneswweenwneneswnwnw",
"swnwseswewnewnesesesesesesesee",
"seswnenwwwewseneswswsenweneswnenenwe",
}

  p1 := P1(input)
  p2 := P2(input, 100)

  fmt.Printf("P1: %d, P2: %d\n", p1, p2)
}

func P1 (input []string) int {
  ret := 0

  floor   := make(map[string]bool)
  matcher := regexp.MustCompile("[ns]?[ew]")

  for _, instruction := range input {
    xOffset := 0
    x       := 0
    y       := 0
    for _, match := range matcher.FindAllString(instruction, -1) {
      switch match {
      case "e":
        x++

      case "w":
        x--

      case "ne":
        y--
        if xOffset == 0 {
          xOffset = 1
        } else {
          x++
          xOffset = 0
        }

      case "nw":
        y--
        if xOffset == 0 {
          x--
          xOffset = 1
        } else {
          xOffset = 0
        }

      case "se":
        y++
        if xOffset == 0 {
          xOffset = 1
        } else {
          x++
          xOffset = 0
        }

      case "sw":
        y++
        if xOffset == 0 {
          x--
          xOffset = 1
        } else {
          xOffset = 0
        }
      }
    }

    tile := fmt.Sprintf("%d,%d", x, y)
    floorVal, ok := floor[tile]
    if (!ok) {
      floor[tile] = true
    } else if floorVal == true {
      floor[tile] = false
    }
  }

  for _, v := range floor {
    if v {
      ret++
    }
  }

  return ret
}

func P2 (input []string, days int) int {
  ret := 0

  floor   := make([][]bool, 200)
  for idx, _ := range floor {
    floor[idx] = make([]bool, 200)
  }

  offset := 100

  matcher := regexp.MustCompile("[ns]?[ew]")

  for _, instruction := range input {
    xOffset := 0
    x       := 0
    y       := 0
    for _, match := range matcher.FindAllString(instruction, -1) {
      switch match {
      case "e":
        x++

      case "w":
        x--

      case "ne":
        y--
        if xOffset == 0 {
          xOffset = 1
        } else {
          x++
          xOffset = 0
        }

      case "nw":
        y--
        if xOffset == 0 {
          x--
          xOffset = 1
        } else {
          xOffset = 0
        }

      case "se":
        y++
        if xOffset == 0 {
          xOffset = 1
        } else {
          x++
          xOffset = 0
        }

      case "sw":
        y++
        if xOffset == 0 {
          x--
          xOffset = 1
        } else {
          xOffset = 0
        }
      }
    }

    floor[y + offset][x + offset] = !floor[y + offset][x + offset]
  }

  for d := 0; d < days; d++ {
    floor = DailyFlip(floor)
  }

  for _, row := range floor {
    for _, v := range row {
      if v {
        ret++
      }
    }
  }

  return ret
}

func DailyFlip(floor [][]bool) [][]bool {
  ret := make([][]bool, len(floor))

  for idx, _ := range ret {
    ret[idx] = make([]bool, len(floor))
  }

  xOffset := -1

  for y := 2; y < len(floor) - 2; y++ {
    for x := 2; x < len(floor) - 2; x++ {
      adjBlack := 0

      if floor[y][x - 1] {
        adjBlack++
      }
      if floor[y][x + 1] {
        adjBlack++
      }

      if floor[y + 1][x] {
        adjBlack++
      }
      if floor[y - 1][x] {
        adjBlack++
      }

      if floor[y + 1][x + xOffset] {
        adjBlack++
      }
      if floor[y - 1][x + xOffset] {
        adjBlack++
      }

      if floor[y][x] {
        if adjBlack == 0 || adjBlack > 2  {
          ret[y][x] = false
        } else {
          ret[y][x] = true
        }
      } else {
        if adjBlack == 2 {
          ret[y][x] = true
        } else {
          ret[y][x] = false
        }
      }
    }

    if xOffset == -1 {
      xOffset = 1
    } else {
      xOffset = -1
    }
  }

  return ret
}

func PrintFloor(floor [][]bool) {
  offset := ""

  for y, row := range floor {
    fmt.Printf(offset)
    for x, v := range row {
      if y == 20 && x == 20 {
        fmt.Printf("(")
      } else {
        fmt.Printf(" ")
      }
      if v {
        fmt.Printf("#")
      } else {
        fmt.Printf(".")
      }
    }

    if offset == " " {
      offset = ""
    } else {
      offset = " "
    }

    fmt.Println()
  }
}

