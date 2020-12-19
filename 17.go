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

  p1 := P1(input, 6)
  p2 := P2()

  fmt.Printf("P1: %d, P2: %d\n", p1, p2)
}

type Cube struct {
  x      int
  y      int
  z      int
}

func P1 (input string, rounds int) int {
  cubeMap := make(map[Cube]bool)
  ret     := 0

  // Fill in our initialisation section
  for y, line := range strings.Split(input, "\n") {
    for x, char := range line {
      value := char == '#'
      cubeMap[Cube {x, y, 0} ] = value
    }
  }

  for i := 0; i < rounds; i++ {
    cubeMap = Cycle(cubeMap)
  }

  for _, value := range cubeMap {
    if value {
      ret++
    }
  }

  return ret
}

func P2 () int {
  return 1
}

func Cycle(input map[Cube] bool) map[Cube]bool {
  ret := make(map[Cube]bool)
  for cube, _ := range input {
    for x := cube.x - 1; x <= cube.x + 1; x++ {
      for y := cube.y - 1; y <= cube.y + 1; y++ {
        for z := cube.z - 1; z <= cube.z + 1; z++ {

          activeNeighbours := ActiveNeighbours(input, x, y, z)

          active, exists := input[Cube { x, y, z}]
          if !exists || !active {
            if activeNeighbours == 3 {
              ret[Cube {x, y, z}] = true
            } else {
              ret[Cube {x, y, z}] = false
            }
          } else {
            if activeNeighbours == 2 || activeNeighbours == 3 {
              // Remains active
              ret[Cube {x, y, z}] = true
            } else {
              // Inactivated
              ret[Cube {x, y, z}] = false
            }
          }
        }
      }
    }
  }

  return ret
}

func ActiveNeighbours(input map[Cube]bool, xIn int, yIn int, zIn int) int {
  activeNeighbours := 0

  for x := xIn - 1; x <= xIn + 1; x++ {
    for y := yIn - 1; y <= yIn + 1; y++ {
      for z := zIn - 1; z <= zIn + 1; z++ {
        if x == xIn && y == yIn && z == zIn {
          // Skip over the cube we're checking
          continue
        }

        adj, exists := input[Cube { x, y, z}]
        if exists && adj {
          activeNeighbours++
        }
      }
    }
  }

  return activeNeighbours
}

func PrintMap(input map[Cube]bool) {
  for z := -100; z <= 100; z++ {
    level := ""

    for y := -100; y <= 100; y++ {

      yExists := false

      for x := -100; x <= 100; x++ {
        active, exists := input[Cube {x, y, z}]

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
