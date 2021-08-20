package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
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
	return strings.Contains(passport, "byr:") && 
	strings.Contains(passport, "iyr:") && 
	strings.Contains(passport, "eyr:") && 
	strings.Contains(passport, "hgt:") && 
	strings.Contains(passport, "hcl:") && 
	strings.Contains(passport, "ecl:") && 
	strings.Contains(passport, "pid:") 
}

func validateFields(passport string) bool {
	valid := true
	fields := strings.Split(passport, " ")
	for _, field := range fields {
		keyAndValue := strings.Split(field, ":")
		if len(keyAndValue) == 2 {
			key := keyAndValue[0]
			value := keyAndValue[1]

			valid = valid && checkKeyAndValue(key, value)
			if !valid {
				break
			}
		}
	}
	return valid
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func checkKeyAndValue(key string, value string) bool {
	/*
	Validations:
	- byr (Birth Year) - four digits; at least 1920 and at most 2002.
	- iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	- eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	- hgt (Height) - a number followed by either cm or in:
		- If cm, the number must be at least 150 and at most 193.
		- If in, the number must be at least 59 and at most 76.
	- hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	- ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	- pid (Passport ID) - a nine-digit number, including leading zeroes.
	- cid (Country ID) - ignored, missing or not.
	*/
	eyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	sort.Strings(eyeColors)
	cmValue := []string{}
	inValue := []string{}

	switch key {

		case "byr":
			v,e := strconv.ParseInt(value, 10, 64) 
			if e != nil {
				fmt.Println(e)
			}
			return len(value) == 4 && v >= 1920 && v <= 2002

		case "iyr":
			v,e := strconv.ParseInt(value, 10, 64) 
			if e != nil {
				fmt.Println(e)
			}
			return len(value) == 4 && v >= 2010 && v<= 2020

		case "eyr":
			v,e := strconv.ParseInt(value, 10, 64) 
			if e != nil {
				fmt.Println(e)
			}
			return len(value) == 4 && v >= 2020 && v <= 2030

		case "hgt":
			if strings.Contains(value, "cm") {
				cmValue = strings.Split(value, "cm")
				if len(cmValue) == 2 {
					v,e := strconv.ParseInt(cmValue[0], 10, 64) 
					if e != nil {
						fmt.Println(e)
					} 
					return  v >= 150 && v <= 193
				}
				return false
			}
			if strings.Contains(value, "in") {
				inValue = strings.Split(value, "in")
				if len(inValue) == 2 {
					v,e := strconv.ParseInt(inValue[0], 10, 64) 
					if e != nil {
						fmt.Println(e)
					} 
					return  v >= 59 && v <= 76
				}
				return false
			}
			return false

		case "hcl":
			r, _ := regexp.Compile("^#(?:[0-9a-fA-F]{3}){2}$")
			return len(value) == 7 && value[0] == '#' && r.MatchString(value)

		case "ecl": 
			return contains(eyeColors, value)

		case "pid":
			return len(value) == 9

		case "cid":
			return true
	}
	return false
}


func getValidPassportCount(passports []string) int {
	validCnt := 0
	for _, passport := range passports {
		if checkFieldsExists(passport) && validateFields(passport) {
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