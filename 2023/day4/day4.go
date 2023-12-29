package main

import (
	service "advent"
	"fmt"
	"math"
	"slices"
	"strings"
)

func main() {
	sumPoints := part2()
	fmt.Println(sumPoints)
}

func part1() int {
	lines := service.ReadFile("day4/input.txt")
	sumPoints := 0
	for _, value := range lines {
		matches := 0
		match := []string{}
		firstTreat := strings.SplitAfter(value, ":")[1]
		firstTreat = strings.Trim(firstTreat, " ")

		winningNumbers := strings.Split(strings.Trim(strings.Split(firstTreat, "|")[0], " "), " ")
		numbers := strings.Split(strings.Trim(strings.Split(firstTreat, "|")[1], " "), " ")

		for _, number := range winningNumbers {
			if slices.Contains(numbers, number) {

				match = append(match, number)
				matches++
			}
		}

		fmt.Println(winningNumbers, "Matches ---> ", matches, "Matched numbers ---> ", match)
		if matches > 0 {
			sumPoints += int(math.Pow(2, float64(matches-1)))
		}
	}
	return sumPoints
}

func part2() int {
	lines := service.ReadFile("day4/input.txt")
	numberOfScratchCards := []int{}
	for range lines {
		numberOfScratchCards = append(numberOfScratchCards, 1)

	}
	sumPoints := 0
	for index, value := range lines {
		matches := 0
		match := []string{}
		firstTreat := strings.SplitAfter(value, ":")[1]
		firstTreat = strings.Trim(firstTreat, " ")

		winningNumbers := strings.Split(strings.Trim(strings.Split(firstTreat, "|")[0], " "), " ")
		numbers := strings.Split(strings.Trim(strings.Split(firstTreat, "|")[1], " "), " ")

		for _, number := range winningNumbers {
			if slices.Contains(numbers, number) && number != "" {

				match = append(match, number)
				matches++
			}
		}

		if matches > 0 {
			for i := 1; i <= matches; i++ {
				numberOfScratchCards[index+i] = numberOfScratchCards[index+i] + numberOfScratchCards[index]
			}
		}
	}
	for _, value := range numberOfScratchCards {
		sumPoints += value
	}
	return sumPoints

}
