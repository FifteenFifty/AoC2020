package main

import (
    "fmt"
    "strings"
)

func main() {
  input := `##.#####
#.##..#.
.##...##
###.#...
.#######
##....##
###.###.
.#.#.#..`

//input  = `.#.
//..#
//###`

  p1 := P1(input, 6)
  p2 := P2(input, 6)

  fmt.Printf("P1: %d, P2: %d\n", p1, p2)
}

type Cube3 struct {
  x int
  y int
  z int
}

type Cube4 struct {
  x int
  y int
  z int
  w int
}

func P1 (input string, rounds int) int {
  cubeMap := make(map[Cube3]bool)
  ret     := 0

  // Fill in our initialisation section
  for y, line := range strings.Split(input, "\n") {
    for x, char := range line {
      value := char == '#'
      cubeMap[Cube3 {x, y, 0} ] = value
    }
  }

  for i := 0; i < rounds; i++ {
    cubeMap = Cycle3(cubeMap)
  }

  for _, value := range cubeMap {
    if value {
      ret++
    }
  }

  return ret
}

func P2 (input string, rounds int) int {
  cubeMap := make(map[Cube4]bool)
  ret     := 0

  // Fill in our initialisation section
  for y, line := range strings.Split(input, "\n") {
    for x, char := range line {
      value := char == '#'
      cubeMap[Cube4 {x, y, 0, 0} ] = value
    }
  }

  for i := 0; i < rounds; i++ {
    cubeMap = Cycle4(cubeMap)
  }

  for _, value := range cubeMap {
    if value {
      ret++
    }
  }

  return ret
}

func Cycle3(input map[Cube3] bool) map[Cube3]bool {
  ret := make(map[Cube3]bool)
  for cube, _ := range input {
    for x := cube.x - 1; x <= cube.x + 1; x++ {
      for y := cube.y - 1; y <= cube.y + 1; y++ {
        for z := cube.z - 1; z <= cube.z + 1; z++ {

          activeNeighbours := ActiveNeighbours(input, x, y, z)

          active, exists := input[Cube3 { x, y, z}]
          if !exists || !active {
            if activeNeighbours == 3 {
              ret[Cube3 {x, y, z}] = true
            } else {
              ret[Cube3 {x, y, z}] = false
            }
          } else {
            if activeNeighbours == 2 || activeNeighbours == 3 {
              // Remains active
              ret[Cube3 {x, y, z}] = true
            } else {
              // Inactivated
              ret[Cube3 {x, y, z}] = false
            }
          }
        }
      }
    }
  }

  return ret
}

func Cycle4(input map[Cube4] bool) map[Cube4]bool {
  ret := make(map[Cube4]bool)
  for cube, _ := range input {
    for x := cube.x - 1; x <= cube.x + 1; x++ {
      for y := cube.y - 1; y <= cube.y + 1; y++ {
        for z := cube.z - 1; z <= cube.z + 1; z++ {
          for w := cube.w - 1; w <= cube.w + 1; w++ {

            activeNeighbours := ActiveNeighbours4(input, x, y, z, w)

            active, exists := input[Cube4 { x, y, z, w}]
            if !exists || !active {
              if activeNeighbours == 3 {
                ret[Cube4 {x, y, z, w}] = true
              } else {
                ret[Cube4 {x, y, z, w}] = false
              }
            } else {
              if activeNeighbours == 2 || activeNeighbours == 3 {
                // Remains active
                ret[Cube4 {x, y, z, w}] = true
              } else {
                // Inactivated
                ret[Cube4 {x, y, z, w}] = false
              }
            }
          }
        }
      }
    }
  }

  return ret
}

func ActiveNeighbours(input map[Cube3]bool, xIn int, yIn int, zIn int) int {
  activeNeighbours := 0

  for x := xIn - 1; x <= xIn + 1; x++ {
    for y := yIn - 1; y <= yIn + 1; y++ {
      for z := zIn - 1; z <= zIn + 1; z++ {
        if x == xIn && y == yIn && z == zIn {
          // Skip over the cube we're checking
          continue
        }

        adj, exists := input[Cube3 { x, y, z}]
        if exists && adj {
          activeNeighbours++
        }
      }
    }
  }

  return activeNeighbours
}

func ActiveNeighbours4(input map[Cube4]bool, xIn int, yIn int, zIn int, wIn int) int {
  activeNeighbours := 0

  for x := xIn - 1; x <= xIn + 1; x++ {
    for y := yIn - 1; y <= yIn + 1; y++ {
      for z := zIn - 1; z <= zIn + 1; z++ {
        for w := wIn - 1; w <= wIn + 1; w++ {
          if x == xIn && y == yIn && z == zIn && w == wIn {
            // Skip over the cube we're checking
            continue
          }

          adj, exists := input[Cube4 { x, y, z, w}]
          if exists && adj {
            activeNeighbours++
          }
        }
      }
    }
  }

  return activeNeighbours
}

func PrintMap3(input map[Cube3]bool) {
  for z := -100; z <= 100; z++ {
    level := ""

    for y := -100; y <= 100; y++ {

      yExists := false

      for x := -100; x <= 100; x++ {
        active, exists := input[Cube3 {x, y, z}]

        if exists {
          yExists = true
          if active {
            level += "#"
          } else {
            level += "."
          }
        }
      }

      if yExists {
        level += "\n"
      }
    }

    if level != "" {
      fmt.Printf("z=%d\n%s\n", z, level)
    }
  }
}
