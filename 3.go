package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func readInput(path string) ([][]string, error) {
  file, err := os.Open(path)

  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines [][]string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, strings.Split(scanner.Text(), ""))
  }

  return lines, scanner.Err()
}

func getTotalTreesTraversedBySlope(sample [][]string, step int, down int) (int) {
  var treeCount int = 0
  startIdx := 0

  for i := 0; i < len(sample) - down; i += down {
    row := sample[i]
    startIdx = int(math.Mod(float64(startIdx + step), float64(len(row))));

    if sample[i + down][startIdx] == "#" {
      treeCount++
    }
  }
  return treeCount
}

func main() {
    sample, e  := readInput("input-3.txt")
    if e != nil {
      fmt.Println("Reading input", e)
    }

    totalCountSlope1 := getTotalTreesTraversedBySlope(sample, 1, 1)
    totalCountSlope3 := getTotalTreesTraversedBySlope(sample, 3, 1)
    totalCountSlope5 := getTotalTreesTraversedBySlope(sample, 5, 1)
    totalCountSlope7 := getTotalTreesTraversedBySlope(sample, 7, 1)
    totalCountSlopeHalf := getTotalTreesTraversedBySlope(sample, 1, 2)

    fmt.Printf("slope 3: %d\n", totalCountSlope3)
    fmt.Printf("slope 1: %d\n", totalCountSlope1)
    fmt.Printf("slope 5: %d\n", totalCountSlope5)
    fmt.Printf("slope 7: %d\n", totalCountSlope7)
    fmt.Printf("slope half: %d\n", totalCountSlopeHalf)

    fmt.Printf("part 1 solution: %d\n", totalCountSlope3)
    fmt.Printf("part 2 solution: %d\n", (totalCountSlope1 * totalCountSlope3 * totalCountSlope5 * totalCountSlope7 * totalCountSlopeHalf))
}