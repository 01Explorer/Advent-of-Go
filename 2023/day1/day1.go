package main

import (
	service "advent"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := service.ReadFile("day1/inputsPart1.txt")
	numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	sum := 0

	for _, line := range lines {
		first, last := 0, 0
		for index, char := range line {
			found := false
			for nIndex, number := range numbers {
				if strings.Contains(line[:index], number) {
					first = nIndex
					found = true
					break
				}
			}
			if found {
				break
			}
			charAsInt, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			first = charAsInt
			break
		}

		for i := len(line) - 1; i >= 0; i-- {
			found := false
			for nIndex, number := range numbers {
				if strings.Contains(line[i:], number) {
					last = nIndex
					found = true
					break
				}
			}
			if found {
				break
			}
			char := line[i]
			charAsInt, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			last = charAsInt
			break
		}
		sum += first*10 + last
	}

	fmt.Println(sum)
}
