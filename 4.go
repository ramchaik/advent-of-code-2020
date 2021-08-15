package main

import (
	"bufio"
	"fmt"
	"os"
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
	input := ""
	for scanner.Scan() {
		input += (scanner.Text() + " ")
		if scanner.Text() == "" {
			lines = append(lines, strings.Replace(input, "\n", " ", -1))
			input = ""
		}
	}

	if input != "" {
		lines = append(lines, strings.Replace(input, "\n", " ", -1))
	}

	return lines, scanner.Err()
}

func checkFieldsExists(passport string) bool {
	/*
	Required fields:
	- byr (Birth Year)
	- iyr (Issue Year)
	- eyr (Expiration Year)
	- hgt (Height)
	- hcl (Hair Color)
	- ecl (Eye Color)
	- pid (Passport ID)

	Optional fields:
	- cid (Country ID)
	*/
	valid := (strings.Contains(passport, "byr:") && 
	strings.Contains(passport, "iyr:") && 
	strings.Contains(passport, "eyr:") && 
	strings.Contains(passport, "hgt:") && 
	strings.Contains(passport, "hcl:") && 
	strings.Contains(passport, "ecl:") && 
	strings.Contains(passport, "pid:")) 
	return valid
}

func getValidPassportCount(passports []string) int {
	validCnt := 0
	for _, passport := range passports {
		if checkFieldsExists(passport) {
			validCnt++
		}
	}
	return validCnt
}

func main() {
	sample, e  := readInput("input.txt")
	if e != nil {
		fmt.Println("Reading input", e)
	}

	fmt.Println("valid passport count: ", getValidPassportCount(sample))
}