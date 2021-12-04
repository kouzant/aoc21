package day4

import "aoc21/utils"

type Point struct {
	Value int
	Marked bool
}

type Grid struct {
	X []Point
	Y []int
}
type Bingo struct {
	Sequence []int
	Boards []*Point
}

func BingoBoard() (Bingo, error) {
	_, err := utils.ReadFile("inputs/bingo")
	if err != nil {
		return Bingo{}, err
	}

	var bingo Bingo
	return bingo, nil
}

func CreatePuzzle1() *Puzzle1 {
	return &Puzzle1{}
}

type Puzzle1 struct {}

func (p *Puzzle1) Solve() error {
	return nil
}
