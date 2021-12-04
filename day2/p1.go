package day2

import (
	"aoc21/utils"
	"fmt"
)

func CreatePuzzle1() *Puzzle1 {
	return &Puzzle1{}
}

type Puzzle1 struct {}

func (p *Puzzle1) Solve() error {
	course, err := utils.PlannedCourse()
	if err != nil {
		return err
	}
	var horizontal int
	var depth int
	for _, c := range course {
		switch c.Type {
		case "FORWARD":
			horizontal += c.Unit
		case "DOWN":
			depth += c.Unit
		case "UP":
			depth -= c.Unit
		}
	}
	fmt.Printf("Result is %d\n", horizontal*depth)
	return nil
}
