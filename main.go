package main

import (
	"aoc21/day1"
	"aoc21/day2"
	"aoc21/day3"
	"aoc21/day4"
	"fmt"
	"os"
)

type Puzzle interface {
	Solve() error
}

func main() {
	puzzle, err := getPuzzle(os.Args[1])
	if err != nil {
		panic(err)
	}
	if err = puzzle.Solve(); err != nil {
		panic(err)
	}
}

func getPuzzle(puzzle string) (Puzzle, error) {
	switch puzzle {
	case "d1p1":
		return day1.CreatePuzzle1(), nil
	case "d1p2":
		return day1.CreatePuzzle2(), nil
	case "d2p1":
		return day2.CreatePuzzle1(), nil
	case "d2p2":
		return day2.CreatePuzzle2(), nil
	case "d3p1":
		return day3.CreatePuzzle1(), nil
	case "d3p2":
		return day3.CreatePuzzle2(), nil
	case "d4p1":
		return day4.CreatePuzzle1(), nil
	}
	return nil, fmt.Errorf("no puzzle with name: %s", puzzle)
}