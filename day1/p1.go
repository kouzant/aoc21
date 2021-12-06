package day1

import (
	"aoc21/inputs"
	"fmt"
	"strconv"
	"strings"
)

func DepthReadings() ([]int, error) {
	rs := strings.Split(inputs.DepthReadings, "\n")
	var depths []int
	for _, d := range rs {
		di, _ := strconv.Atoi(d)
		depths = append(depths, di)
	}
	if len(depths) == 0 {
		return nil, fmt.Errorf("oops no depths readings")
	}
	return depths, nil
}

func CreatePuzzle1() *Puzzle1 {
	return &Puzzle1{}
}

type Puzzle1 struct {}

func (p *Puzzle1) Solve() error {

	readings, err := DepthReadings()
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
