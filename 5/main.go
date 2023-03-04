package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

type rucksack struct {
	compartment_1 []int
	compartment_2 []int
}

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

func parse(input string) rucksack {
	// check if the length of the input is even
	if len(input)%2 != 0 {
		panic("Invalid input (compartments must be of even length) " + input)
	}
	midpoint := len(input) / 2
	compartment_1 := make([]int, midpoint)
	compartment_2 := make([]int, midpoint)
	for i := 0; i < midpoint; i++ {
		compartment_1[i] = translate(string(input[i]))
		compartment_2[i] = translate(string(input[i+midpoint]))
	}
	return rucksack{compartment_1, compartment_2}
}

func invalidly_packed(bag rucksack) []int {
	invalid := []int{}
	for i := 0; i < len(bag.compartment_1); i++ {
		for j := 0; j < len(bag.compartment_2); j++ {
			if bag.compartment_1[i] == bag.compartment_2[j] {
				item := bag.compartment_1[i]
				if slices.Contains(invalid, item) {
					continue
				}
				invalid = append(invalid, item)
			}
		}
	}
	return invalid
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum_of_invalid := 0
	for scanner.Scan() {
		rawinput := scanner.Text()
		bag := parse(rawinput)
		invalid := invalidly_packed(bag)
		if len(invalid) > 1 {
			panic("Invalid input (more than one invalid item) " + rawinput)
		}
		for _, item := range invalid {
			sum_of_invalid += item
		}
	}
	fmt.Println(sum_of_invalid)
}
