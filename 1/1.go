package main

import (
	"bufio"
	"fmt"
	"os"
)

type elf_backpack struct {
	food_items []uint
}

type Day1 struct {
	elf_backpacks []elf_backpack
}

func read_input() Day1 {
	// read 1.txt file and return Day1 struct
	file, err := os.Open("1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	elfs := []elf_backpack{}
	current_elf := []uint{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elfs = append(elfs, elf_backpack{current_elf})
			current_elf = []uint{}
		} else {
			parsed_calories := 0
			fmt.Sscanf(line, "%d", &parsed_calories)
			current_elf = append(current_elf, uint(parsed_calories))
		}
	}
	day1 := Day1{elfs}
	return day1
}

func main() {
	day1 := read_input()
	var best uint = 0
	for _, elf := range day1.elf_backpacks {
		var sum uint = 0
		for _, food_item := range elf.food_items {
			sum += food_item
		}
		if sum > best {
			best = sum
		}
	}
	fmt.Println(best)
}
