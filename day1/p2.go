package day1

import (
	"aoc21/utils"
	"fmt"
	"math"
)

func CreatePuzzle2() *Puzzle2 {
	return &Puzzle2{}
}

type Puzzle2 struct {}

func (p *Puzzle2) Solve() error {
	readings, err := utils.DepthReadings()
	if err != nil {
		return err
	}
	var windowIndex int

	previousWindow := math.MaxInt
	var increased int
	for oi, _ := range readings {
		windowIndex = oi
		var slidingSum int
		for i := windowIndex; i < windowIndex+3; i++ {
			slidingSum += readings[i]
		}
		if slidingSum > previousWindow {
			increased++
		}
		previousWindow = slidingSum
		if oi + 3 >= len(readings) {
			break
		}
	}
	fmt.Printf("Sliding window increments: %d\n", increased)
	return nil
}
