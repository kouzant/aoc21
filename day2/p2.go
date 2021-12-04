package day2

import (
	"aoc21/utils"
	"fmt"
)

func CreatePuzzle2() *Puzzle2 {
	return &Puzzle2{}
}

type Puzzle2 struct {}

func (p *Puzzle2) Solve() error {
	course, err := utils.PlannedCourse()
	if err != nil {
		return err
	}

	var horizontal int
	var depth int
	var aim int
	for _, c := range course {
		switch c.Type {
		case "FORWARD":
			horizontal += c.Unit
			depth += aim * c.Unit
		case "DOWN":
			aim += c.Unit
		case "UP":
			aim -= c.Unit
		}
	}
	fmt.Printf("Result is: %d\n", horizontal*depth)
	return nil
}
