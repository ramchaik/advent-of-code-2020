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

func partOneSolution(sample [][]string) (int) {
  var treeCount int = 0
  startIdx := 0
  step := 3

  for i := 0; i < len(sample) - 1; i++ {
    row := sample[i]
    startIdx = int(math.Mod(float64(startIdx + step), float64(len(row))));

    if sample[i + 1][startIdx] == "#" {
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
    treeCount := partOneSolution(sample)
    fmt.Printf("tree count: %d\n", treeCount)
}