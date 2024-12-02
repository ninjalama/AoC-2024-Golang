package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

func allIncreasing(intList []int) bool {
	for idx := range intList {
		if idx == len(intList)-1 {
			// At last position, nothing to compare to so returning true and great success
			return true
		}

		first := intList[idx]
		second := intList[idx+1]

		if !(second > first && second-first <= 3) {
			return false
		}
	}
	return true
}

func allDecreasing(intList []int) bool {
	intListCopy := append([]int(nil), intList...)
	slices.Reverse(intListCopy)
	return allIncreasing(intListCopy)
}

func part1(numberList [][]int) int {
	safeReports := 0
	for idx := range numberList {
		if allIncreasing(numberList[idx]) || allDecreasing(numberList[idx]) {
			safeReports++
		}
	}

	return safeReports
}

func mutations(intList []int) [][]int {
	var combinations [][]int
	for i := 0; i < len(intList); i++ {
		var mutation []int
		mutation = append(mutation, intList[:i]...)
		mutation = append(mutation, intList[i+1:]...)
		combinations = append(combinations, mutation)
	}
	return combinations
}

func part2(numberList [][]int) int {
	safeReports := 0
	for idx := range numberList {
		mutationLists := mutations(numberList[idx])
		// Check if any of the mutations are safe
		reportIsSafe := false
		for _, mutatedList := range mutationLists {
			if allIncreasing(mutatedList) || allDecreasing(mutatedList) {
				reportIsSafe = true
				break
			}
		}
		if reportIsSafe {
			safeReports++
		}
	}

	return safeReports
}

func main() {
	inputStr, err := readInput("day02/input.txt")

	if err != nil {
		fmt.Println("Could not read input!", err)
		os.Exit(0)
	}

	inputLines := strings.Split(inputStr, "\r\n")
	nrOfLines := len(inputLines)

	numberList := make([][]int, nrOfLines)

	for rowIdx, row := range inputLines {
		columns := strings.Split(row, " ")
		numberList[rowIdx] = make([]int, len(columns))

		for colIdx, col := range columns {
			num, err := strconv.Atoi(col)
			if err == nil {
				numberList[rowIdx][colIdx] = num
			}
		}
	}

	safeReportsPart1 := part1(numberList)
	safereportsPart2 := part2(numberList)

	fmt.Println("Safe reports in part1:", safeReportsPart1)
	fmt.Println("Safe reports in part2:", safereportsPart2)
}
