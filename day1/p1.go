package day1

import (
	"aoc21/utils"
	"fmt"
)

func CreatePuzzle1() *Puzzle1 {
	return &Puzzle1{}
}

type Puzzle1 struct {}

func (p *Puzzle1) Solve() error {

	readings, err := utils.DepthReadings()
	if err != nil {
		return err
	}
	var increased int
	for i := 1; i < len(readings); i++ {
		previousReading := readings[i - 1]
		currentReading := readings[i]
		if currentReading > previousReading {
			increased++
		}
	}
	fmt.Printf("Depth increased %d times\n", increased)
	return nil
}
