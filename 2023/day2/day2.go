package main

import (
	service "advent"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	part2()
}

func part1() {

	maxRed, maxGreen, maxBlue := 12, 13, 14
	lines := service.ReadFile("day2/input.txt")
	sumValidGames := 0

	for _, line := range lines {
		firstTreat := strings.SplitAfter(line, ":")
		gameIndex, err := strconv.Atoi(strings.ReplaceAll(strings.Split(firstTreat[0], " ")[1], ":", ""))
		if err != nil {
			log.Fatal(err)
		}

		rounds := strings.Split(firstTreat[1], ";")
		for _, round := range rounds {
			if gameIndex == 0 {
				break
			}
			for _, play := range strings.Split(round, ",") {

				play = strings.Trim(play, " ")
				numberOfSquares, err := strconv.Atoi(strings.Split(play, " ")[0])
				if err != nil {
					log.Fatal(err)
				}
				color := strings.Split(play, " ")[1]

				if color == "red" {
					if numberOfSquares > maxRed {
						gameIndex = 0
						break
					}
				}
				if color == "blue" {
					if numberOfSquares > maxBlue {
						gameIndex = 0
						break
					}
				}
				if color == "green" {
					if numberOfSquares > maxGreen {
						gameIndex = 0
						break
					}
				}

			}
		}
		sumValidGames += gameIndex
	}
	fmt.Println(sumValidGames)
}

func part2() {

	lines := service.ReadFile("day2/input.txt")
	sumValidGames := 0

	for _, line := range lines {
		maxRed, maxGreen, maxBlue := 0, 0, 0
		firstTreat := strings.SplitAfter(line, ":")

		rounds := strings.Split(firstTreat[1], ";")
		for _, round := range rounds {
			for _, play := range strings.Split(round, ",") {

				play = strings.Trim(play, " ")
				numberOfSquares, err := strconv.Atoi(strings.Split(play, " ")[0])
				if err != nil {
					log.Fatal(err)
				}
				color := strings.Split(play, " ")[1]

				if color == "red" {
					if maxRed < numberOfSquares {
						maxRed = numberOfSquares
					}
				}
				if color == "blue" {
					if maxBlue < numberOfSquares {
						maxBlue = numberOfSquares
					}
				}
				if color == "green" {
					if maxGreen < numberOfSquares {
						maxGreen = numberOfSquares
					}
				}

			}
		}
		sumValidGames += maxRed * maxBlue * maxGreen
	}
	fmt.Println(sumValidGames)
}
