package main

import (
	"bufio"
	"os"
)

var score_map = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var symbol_map = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

var outcome_score = map[int]int{
	1:  6,
	0:  3,
	-1: 0,
}

func counter(their string, strategy string) string {
	switch strategy {
	case "Y":
		switch their {
		case "rock":
			return "rock"
		case "paper":
			return "paper"
		case "scissors":
			return "scissors"
		}
	case "Z":
		switch their {
		case "rock":
			return "paper"
		case "paper":
			return "scissors"
		case "scissors":
			return "rock"
		}
	case "X":
		switch their {
		case "rock":
			return "scissors"
		case "paper":
			return "rock"
		case "scissors":
			return "paper"
		}
	}
	panic("Invalid strategy " + strategy)
}

func result(strategy string, their string) int {
	their_translated := symbol_map[their]
	if their_translated == "" {
		panic("Invalid symbol " + their)
	}
	mine_translated := counter(their_translated, strategy)
	symbol_score := score_map[mine_translated]
	score_from_game := 0
	switch mine_translated {
	case "rock":
		switch their_translated {
		case "rock":
			score_from_game = outcome_score[0]
		case "paper":
			score_from_game = outcome_score[-1]
		case "scissors":
			score_from_game = outcome_score[1]
		}
	case "paper":
		switch their_translated {
		case "rock":
			score_from_game = outcome_score[1]
		case "paper":
			score_from_game = outcome_score[0]
		case "scissors":
			score_from_game = outcome_score[-1]
		}
	case "scissors":
		switch their_translated {
		case "rock":
			score_from_game = outcome_score[-1]
		case "paper":
			score_from_game = outcome_score[1]
		case "scissors":
			score_from_game = outcome_score[0]
		}
	default:
		panic("Invalid symbol")
	}
	return score_from_game + symbol_score
}

func main() {
	file, err := os.Open("4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var my_score int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		first := line[0:1]
		second := line[2:3]
		my_score += result(second, first)
	}
	println(my_score)
}
