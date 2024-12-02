package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Input struct {
	first  []int
	second []int
}

func (i *Input) ExtractInput(fileName string) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nums := (strings.Split(scanner.Text(), " "))

		num1, _ := strconv.Atoi(nums[0])
		num3, _ := strconv.Atoi(nums[3])

		i.first = append(i.first, num1)
		i.second = append(i.second, num3)
	}
}

func (i *Input) SortInput() {
	// Sort the input
	sort.Ints(i.first)
	sort.Ints(i.second)
}

func (i *Input) CalculateDistance() int {
	answer := 0
	for k := range len(i.first) {
		difference := i.first[k] - i.second[k]
		if difference < 0 {
			answer -= difference
		} else {
			answer += difference
		}
	}
	return answer
}

func main() {
	fmt.Println("Day 1 of AOC!")

	first := []int{}
	second := []int{}

	input := Input{
		first,
		second,
	}

	input.ExtractInput("input.txt")

	input.SortInput()

	answer := input.CalculateDistance()
	fmt.Println(answer)
}
