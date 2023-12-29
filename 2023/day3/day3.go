package main

import (
	service "advent"
	"fmt"
	"strconv"
)

func main() {
	part2()
}

func part1() {
	lines := service.ReadFile("day3/input.txt")
	sumParts := 0

	columns, numberOfLines := len(lines[0]), len(lines)

	for i := 0; i < numberOfLines; i++ {
		strNumber := ""
		for j := 0; j < columns; j++ {
			if isDigit(string(lines[i][j])) {
				strNumber += string(lines[i][j])
				if j != columns-1 {
					continue
				}
			}
			if strNumber != "" {
				if ok := isValidPart(lines, i, j, len(strNumber)+1, strNumber, make(map[string][]string)); ok {
					fmt.Println(strNumber)
					number, err := strconv.Atoi(strNumber)
					if err != nil {
						continue
					}
					sumParts += number
				}
				strNumber = ""
			}
		}
	}
	fmt.Println(sumParts)
}

func part2() {
	mapGears := make(map[string][]string)
	lines := service.ReadFile("day3/input.txt")
	sumParts := 0

	columns, numberOfLines := len(lines[0]), len(lines)

	for i := 0; i < numberOfLines; i++ {
		strNumber := ""
		for j := 0; j < columns; j++ {
			if isDigit(string(lines[i][j])) {
				strNumber += string(lines[i][j])
				if j != columns-1 {
					continue
				}
			}
			if strNumber != "" {
				if ok := isValidPart(lines, i, j, len(strNumber)+1, strNumber, mapGears); ok {
					fmt.Println(strNumber)
				}
				strNumber = ""
			}
		}
	}
	for _, value := range mapGears {
		if len(value) == 2 {
			gear1, err := strconv.Atoi(value[0])
			if err != nil {
				continue
			}
			gear2, err := strconv.Atoi(value[1])
			if err != nil {
				continue
			}
			sumParts += gear1 * gear2

		}
	}
	fmt.Println(sumParts)
}

func isDigit(char string) bool {
	_, err := strconv.Atoi(char)
	if err != nil {
		return false
	}
	return true
}

func isValidPart(lines []string, i, j, numberLenght int, number string, gears map[string][]string) bool {
	columns := len(lines[0]) - 1
	isStart, isEnd, isFirstLine, isLastLine := false, false, false, false
	if i == 0 {
		isFirstLine = true
	} else if i == len(lines)-1 {
		isLastLine = true
	}

	if j-numberLenght < 0 {
		isStart = true
	} else if j == columns {
		isEnd = true
	}

	if !isStart {
		if isSpecialChar(string(lines[i][j-numberLenght])) {
			addGear(gears, string(lines[i][j-numberLenght]), number, i, j-numberLenght)
			return true
		}
		if !isFirstLine && isSpecialChar(string(lines[i-1][j-numberLenght])) {
			addGear(gears, string(lines[i-1][j-numberLenght]), number, i-1, j-numberLenght)
			return true
		}
		if !isLastLine && isSpecialChar(string(lines[i+1][j-numberLenght])) {
			addGear(gears, string(lines[i+1][j-numberLenght]), number, i+1, j-numberLenght)
			return true
		}
	}
	if !isEnd {
		if isSpecialChar(string(lines[i][j])) {
			addGear(gears, string(lines[i][j]), number, i, j)
			return true
		}
		if !isFirstLine && isSpecialChar(string(lines[i-1][j])) {
			addGear(gears, string(lines[i-1][j]), number, i-1, j)
			return true
		}
		if !isLastLine && isSpecialChar(string(lines[i+1][j])) {
			addGear(gears, string(lines[i+1][j]), number, i+1, j)
			return true
		}
	}
	for k := 0; k < numberLenght; k++ {
		if !isFirstLine {
			if isSpecialChar(string(lines[i-1][j-k])) {
				addGear(gears, string(lines[i-1][j-k]), number, i-1, j-k)
				return true
			}
		}
		if !isLastLine {
			if isSpecialChar(string(lines[i+1][j-k])) {
				addGear(gears, string(lines[i+1][j-k]), number, i+1, j-k)
				return true
			}
		}
	}
	return false
}

func isSpecialChar(char string) bool {
	if char != "." && !isDigit(char) {
		return true
	}
	return false
}

func addGear(gears map[string][]string, char, value string, i, j int) {
	if char != "*" {
		return
	}
	key := fmt.Sprintf("%d,%d", i, j)
	if gears[key] == nil {
		gears[key] = []string{value}
	} else {
		gears[key] = append(gears[key], value)
	}
}
