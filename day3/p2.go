package day3

import (
	"fmt"
	"strconv"
)

func CreatePuzzle2() *Puzzle2 {
	return &Puzzle2{}
}

type Puzzle2 struct {}

func (p *Puzzle2) Solve() error {
	diagnostics, err := Diagnostics()
	if err != nil {
		return err
	}


	oxygenRating, err := calculateRating(diagnostics, mostCommonInPosition)
	if err != nil {
		return err
	}
	co2Scrubber, err := calculateRating(diagnostics, leastCommonInPosition)
	if err != nil {
		return err
	}
	oxygenRatingI, _ := strconv.ParseInt(oxygenRating, 2, 64)
	co2ScrubberI, _ := strconv.ParseInt(co2Scrubber, 2, 64)

	fmt.Printf("Oxygen rating: %d\n", oxygenRatingI)
	fmt.Printf("CO2 scrubber rating: %d\n", co2ScrubberI)

	fmt.Printf("Life support rating: %d\n", oxygenRatingI*co2ScrubberI)
	return nil
}

func calculateRating(initialDiagnostics []string, criteriaFunc func(zeros, aces [BINARY_SIZE]int, position int) rune) (string, error) {
	var examineBit int
	var filteredbinaries []string
	filteredbinaries = initialDiagnostics
	for {

		zeros, aces := calculateBitPopularity(filteredbinaries)
		mostCommonBit := criteriaFunc(zeros, aces, examineBit)
		var filteringBinaries []string

		for _, d := range filteredbinaries {
			if []rune(d)[examineBit] == mostCommonBit {
				filteringBinaries = append(filteringBinaries, d)
			}
		}

		if len(filteringBinaries) == 1 {
			return filteringBinaries[0], nil
		}
		filteredbinaries = filteringBinaries
		if examineBit == BINARY_SIZE {
			return "", fmt.Errorf("exhausted all available positions and could not find rating")
		}
		examineBit++
	}
}

func mostCommonInPosition(zeros, aces [BINARY_SIZE]int, position int) rune {
	if zeros[position] > aces[position] {
		return '0'
	}
	if zeros[position] < aces[position] {
		return '1'
	}
	return '1'
}

func leastCommonInPosition(zeros, aces [BINARY_SIZE]int, position int) rune {
	if zeros[position] > aces[position] {
		return '1'
	}
	if zeros[position] < aces[position] {
		return '0'
	}
	return '0'
}