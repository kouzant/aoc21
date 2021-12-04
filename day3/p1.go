package day3

import (
	"aoc21/utils"
	"fmt"
	"strconv"
	"strings"
)

func CreatePuzzle1() *Puzzle1 {
	return &Puzzle1{}
}

type Puzzle1 struct {}

const BINARY_SIZE = 12

func (p *Puzzle1) Solve() error {
	diagnostics, err := utils.Diagnostics()
	if err != nil {
		return err
	}


	zeros, aces := calculateBitPopularity(diagnostics)

	var gama strings.Builder
	var epsilon strings.Builder
	for i := 0; i < BINARY_SIZE; i++ {
		if zeros[i] > aces[i] {
			gama.WriteRune('0')
			epsilon.WriteRune('1')
		} else {
			gama.WriteRune('1')
			epsilon.WriteRune('0')
		}
	}

	gamaI, _ := strconv.ParseInt(gama.String(), 2, 64)
	epsilonI, _ := strconv.ParseInt(epsilon.String(), 2, 64)

	fmt.Printf("Power consumption is: %d\n", gamaI*epsilonI)
	return nil
}

func calculateBitPopularity(diagnostics []string) ([BINARY_SIZE]int, [BINARY_SIZE]int) {
	var zeros [BINARY_SIZE]int
	var aces [BINARY_SIZE]int
	for _, d := range diagnostics {
		s := []rune(d)
		for idx, c := range s {
			if c == '0' {
				zeros[idx] = zeros[idx] + 1
			} else if c == '1' {
				aces[idx] = aces[idx] + 1
			}
		}
	}
	return zeros, aces
}
