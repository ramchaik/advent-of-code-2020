package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func readInput(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func bsp(lower int,upper int, l string, u string) func(string, int, int) int {
	return  func(str string, sStart int, sEnd int) int {
		mid := 0

		for _,row := range strings.Split(str, "")[sStart: sEnd] {
			mid = (upper + lower) / 2

			if (row == l) {
				upper = mid
			}

			if (row == u) {
				lower = mid + 1
			}
		}
		return lower
	}
}

func getRow(boardingPass string) int {
	return bsp(0, 127, "F", "B")(boardingPass, 0, 7)
}

func getColumn(boardingPass string) int {
	return bsp(0, 7, "L", "R")(boardingPass, 7, 10)
}

func getSeatId(row int, column int) int {
	return row * 8 + column
}

func getMaxSeatId(seatIds []int) int {
	var maxSeatId int = 0
	for i, s := range seatIds {
    if i==0 || s > maxSeatId {
      maxSeatId = s
    }
	}
	return maxSeatId
}

func getMySeatId(seatIds []int) int {
	sort.Ints(seatIds)
	var mySeatId int
	for i, s := range seatIds {

		if (i + 1) < (len(seatIds) - 1) && s + 1 != seatIds[i + 1] {
			mySeatId = s + 1
		}
	}
	return mySeatId
}

func main() {
	sample, e  := readInput("input.txt")
	if e != nil {
		fmt.Println("Reading input", e)
	}

	var seatIds []int

	for _, boardingPass := range sample {
		row := getRow(boardingPass)
		column := getColumn(boardingPass)
		seatId := getSeatId(row, column)
		seatIds = append(seatIds, seatId)
	} 

	maxSeatId := getMaxSeatId(seatIds)
	mySeatId := getMySeatId(seatIds)

	fmt.Println("ans 1 :>> ", maxSeatId)
	fmt.Println("ans 2 :>> ", mySeatId)
}