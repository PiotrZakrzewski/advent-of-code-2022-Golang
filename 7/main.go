package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type task_range struct {
	min int
	max int
}

func parse(input string) task_range {
	sp := strings.Split(input, "-")
	if len(sp) != 2 {
		panic("Invalid input: " + input)
	}
	min := 0
	max := 0
	fmt.Sscanf(sp[0], "%d", &min)
	fmt.Sscanf(sp[1], "%d", &max)
	return task_range{min, max}
}

func overlapping(a task_range, b task_range) bool {
	if a.min >= b.min && a.max <= b.max {
		return true
	}
	if b.min >= a.min && b.max <= a.max {
		return true
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		sp := strings.Split(line, ",")
		if len(sp) != 2 {
			panic("Invalid input: " + line)
		}
		a := parse(sp[0])
		b := parse(sp[1])
		if overlapping(a, b) {
			result += 1
		}
	}
	fmt.Println(result)
}
