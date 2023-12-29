package service

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(filepath string) []string {
	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}
	path = fmt.Sprintf("%s/%s", path, filepath)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
