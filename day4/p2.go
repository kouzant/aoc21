package day4

import "fmt"

func CreatePuzzle2() *Puzzle2 {
	return &Puzzle2{}
}

type Puzzle2 struct {}

func (p *Puzzle2) Solve() error {
	bingo, err := BingoBoard()
	if err != nil {
		return err
	}
	var lastWonBoard Board
	var numberDrawn int
	var boards []Board
	boards = bingo.Boards
	for _, d := range bingo.Sequence {
		var cBoards []Board
		for _, b := range boards {
			markBoard(d, b)
			if boardHasWon(b) {
				lastWonBoard = b
				numberDrawn = d
			} else {
				cBoards = append(cBoards, b)
			}
		}
		boards = cBoards
	}
	score := calculateScore(numberDrawn, lastWonBoard)
	fmt.Printf(">> Score of last winning board is: %d\n", score)
	return nil
}