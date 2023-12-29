package main

import (
	service "advent"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part2()
}

type seedStruct struct {
	seedMin int
	seedMax int
}

func part2() {
	lines := service.ReadFile("day5/input.txt")
	mapPositions := make(map[string][]int)
	currentMapping := ""
	seeds := []seedStruct{}

	for lineIndex, line := range lines {
		if line != "" {
			if strings.Contains(line, "seeds") {
				seedNums := numbersStringToList(line)
				for j := 0; j < len(seedNums); j += 2 {
					seeds = append(seeds, seedStruct{seedMin: strToInt(seedNums[j]), seedMax: (strToInt(seedNums[j]) + strToInt(seedNums[j+1]) - 1)})
				}
				continue
			}
			if strings.Contains(line, ":") {
				str := strings.Split(line, " ")[0]
				currentMapping = str
				mapPositions[str] = []int{}
				continue
			}
			mapPositions[currentMapping] = append(mapPositions[currentMapping], lineIndex)
		}
	}
	fmt.Println(getLowestLocation(lines, seeds, mapPositions))
}

func part1() {
	lines := service.ReadFile("day5/input.txt")
	mapPositions := make(map[string][]int)
	currentMapping := ""
	seeds := []int{}

	for lineIndex, line := range lines {
		if line != "" {
			if strings.Contains(line, "seeds") {
				seedNums := numbersStringToList(line)
				for _, seed := range seedNums {
					seeds = append(seeds, strToInt(seed))
				}
				continue
			}
			if strings.Contains(line, ":") {
				str := strings.Split(line, " ")[0]
				currentMapping = str
				mapPositions[str] = []int{}
				continue
			}
			mapPositions[currentMapping] = append(mapPositions[currentMapping], lineIndex)
		}
	}
	fmt.Println(getLowestLocation(lines, seeds, mapPositions))

}
func getLowestLocation2(lines []string, seeds []seedStruct, relativePositions map[string][]int) int {
	locations := []int{}
	for _, seed := range seeds {
		needle := seed
		needle = getKey2(lines, needle, relativePositions["seed-to-soil"])
		needle = getKey2(lines, needle, relativePositions["soil-to-fertilizer"])
		needle = getKey2(lines, needle, relativePositions["fertilizer-to-water"])
		needle = getKey2(lines, needle, relativePositions["water-to-light"])
		needle = getKey2(lines, needle, relativePositions["light-to-temperature"])
		needle = getKey2(lines, needle, relativePositions["temperature-to-humidity"])
		needle = getKey2(lines, needle, relativePositions["humidity-to-location"])
		locations = append(locations, needle)
	}
	fmt.Println(locations)
	return slices.Min(locations)
}
func getKey2(lines []string, needle seedStruct, linesIndex []int) int {
	for _, lineIndex := range linesIndex {
		line := lines[lineIndex]
		numbersAsString := numbersStringToList(line)
		if needle >= strToInt(numbersAsString[1]) && needle <= (strToInt(numbersAsString[1])+strToInt(numbersAsString[2])) {
			key := strToInt(numbersAsString[0]) + (needle - strToInt(numbersAsString[1]))
			return key
		}
	}
	return needle

}

func getLowestLocation(lines []string, seeds []int, relativePositions map[string][]int) int {
	locations := []int{}
	for _, seed := range seeds {
		needle := seed
		needle = getKey(lines, needle, relativePositions["seed-to-soil"])
		needle = getKey(lines, needle, relativePositions["soil-to-fertilizer"])
		needle = getKey(lines, needle, relativePositions["fertilizer-to-water"])
		needle = getKey(lines, needle, relativePositions["water-to-light"])
		needle = getKey(lines, needle, relativePositions["light-to-temperature"])
		needle = getKey(lines, needle, relativePositions["temperature-to-humidity"])
		needle = getKey(lines, needle, relativePositions["humidity-to-location"])
		locations = append(locations, needle)
	}
	fmt.Println(locations)
	return slices.Min(locations)
}

func getKey(lines []string, needle int, linesIndex []int) int {
	for _, lineIndex := range linesIndex {
		line := lines[lineIndex]
		numbersAsString := numbersStringToList(line)
		if needle >= strToInt(numbersAsString[1]) && needle <= (strToInt(numbersAsString[1])+strToInt(numbersAsString[2])) {
			key := strToInt(numbersAsString[0]) + (needle - strToInt(numbersAsString[1]))
			return key
		}
	}
	return needle

}

func numbersStringToList(strToConvert string) []string {
	if !strings.Contains(strToConvert, ":") {
		return strings.Split(strings.Trim(strToConvert, " "), " ")
	}
	return strings.Split(strings.Trim(strings.SplitAfter(strToConvert, ":")[1], " "), " ")
}

func strToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
