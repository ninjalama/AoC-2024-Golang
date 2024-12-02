package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func readInput(filePath string) (string, error) {
	data, err := os.ReadFile("../inputs/" + filePath)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return string(data), nil
}

type ListHolder struct {
	firstList  []int
	secondList []int
}

func getListsFromInput(input string) ListHolder {
	// Parse the data
	whitespacePattern := `[ \t]+`

	inputNoWhitespace := regexp.MustCompile(whitespacePattern).ReplaceAllString(input, ",")

	inputList := strings.Split(inputNoWhitespace, "\r\n")

	firstList := make([]int, len(inputList))
	secondList := make([]int, len(inputList))

	for index, value := range inputList {
		lineNumbers := strings.Split(value, ",")
		firstNum, err1 := strconv.Atoi(lineNumbers[0])
		secondNum, err2 := strconv.Atoi(lineNumbers[1])

		if err1 == nil && err2 == nil {
			firstList[index] = firstNum
			secondList[index] = secondNum
		}
	}

	return ListHolder{firstList: firstList, secondList: secondList}
}

func part1(listHolder ListHolder) int {
	// Sort both lists by ascending order
	sort.Ints(listHolder.firstList)
	sort.Ints(listHolder.secondList)

	// Calculate distance
	distanceSum := 0
	for i := 0; i < len(listHolder.firstList); i++ {
		distanceSum += absInt(listHolder.firstList[i] - listHolder.secondList[i])
	}

	return distanceSum
}

func part2(listHolder ListHolder) int {
	// Calculate similarity
	similaritySum := 0
	for _, firstListValue := range listHolder.firstList {
		similarityCount := 0
		for _, secondListvalue := range listHolder.secondList {
			if firstListValue == secondListvalue {
				similarityCount++
			}
		}
		similaritySum += firstListValue * similarityCount
	}

	return similaritySum
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	str, err := readInput("day01/input.txt")
	if err != nil {
		fmt.Println("Could not get input!")
		os.Exit(0)
	}

	listHolder := getListsFromInput(str)

	distanceSum := part1(listHolder)
	similaritySum := part2(listHolder)

	fmt.Println("Part 1 - distance sum:", distanceSum)
	fmt.Println("Part 2 - similarity sum:", similaritySum)
}
