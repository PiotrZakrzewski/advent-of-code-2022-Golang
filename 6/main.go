package main

import (
	"bufio"
	"fmt"
	"os"
)

func translate(item string) int {
	if len(item) != 1 {
		panic("Invalid item " + item)
	}
	const Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i, c := range Alphabet {
		if c == rune(item[0]) {
			return (i + 1)
		}
	}
	panic("Invalid item " + item)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum_of_badges := 0
	current := 0
	current_group := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		current_group = append(current_group, line)
		current += 1
		if current == 3 {
			counter := make(map[int]int)
			for _, group_line := range current_group {
				accounted_for := make(map[int]bool)
				for _, c := range group_line {
					translated := translate(string(c))
					if accounted_for[translate(string(c))] {
						continue
					}
					counter[translated] += 1
					accounted_for[translated] = true
				}
			}
			numer_of_3 := 0
			for item, count := range counter {
				if count == 3 {
					sum_of_badges += item
					numer_of_3 += 1
				}
			}
			if numer_of_3 != 1 {
				panic("Invalid input, only one badge expected")
			}
			current_group = []string{}
			current = 0
		}
	}
	fmt.Println(sum_of_badges)
}
