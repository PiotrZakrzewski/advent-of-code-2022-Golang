package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type elf_backpack struct {
	food_items     []uint
	total_calories uint
}

type Day1 struct {
	elf_backpacks []elf_backpack
}

func read_input() Day1 {
	// read 1.txt file and return Day1 struct
	file, err := os.Open("2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	elfs := []elf_backpack{}
	current_elf := []uint{}
	current_total := uint(0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elfs = append(elfs, elf_backpack{current_elf, current_total})
			current_elf = []uint{}
			current_total = 0
		} else {
			parsed_calories := 0
			fmt.Sscanf(line, "%d", &parsed_calories)
			current_elf = append(current_elf, uint(parsed_calories))
			current_total += uint(parsed_calories)
		}
	}
	day1 := Day1{elfs}
	return day1
}

func main() {
	day1 := read_input()
	sort.Slice(day1.elf_backpacks, func(i, j int) bool {
		return day1.elf_backpacks[i].total_calories > day1.elf_backpacks[j].total_calories
	})
	bottom3 := day1.elf_backpacks[:3]
	bottom3sum := uint(0)
	for _, elf := range bottom3 {
		bottom3sum += elf.total_calories
	}
	fmt.Println(bottom3sum)
}
