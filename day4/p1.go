package day4

import (
	"aoc21/inputs"
	"fmt"
	"strconv"
	"strings"
)

const X_SIZE = 5
const Y_SIZE = 5

type Point struct {
	Value int
	Marked bool
}

func (p Point) String() string {
	var s strings.Builder
	s.WriteString(strconv.Itoa(p.Value))
	if p.Marked {
		s.WriteString("<")
	}
	return s.String()
}

type Board struct {
	Points [Y_SIZE][X_SIZE]*Point
}

func (b Board) String() string {
	var s strings.Builder
	for i := 0; i < Y_SIZE; i++ {
		for j := 0; j < X_SIZE; j++ {
			s.WriteString(b.Points[i][j].String())
			if j < X_SIZE - 1 {
				s.WriteString(" ")
			}
		}
		s.WriteString("\n")
	}
	return s.String()
}

type Bingo struct {
	Sequence []int
	Boards []Board
}

func (b Bingo) String() string {
	var s strings.Builder
	s.WriteString("Sequence: ")
	for idx, n := range b.Sequence {
		s.WriteString(strconv.Itoa(n))
		if idx < len(b.Sequence) - 1 {
			s.WriteString(" ")
		}
	}
	s.WriteString("\n\n")
	s.WriteString("Boards:\n")
	for _, bs := range b.Boards {
		s.WriteString(bs.String())
		s.WriteString("\n\n")
	}
	return s.String()
}

func BingoBoard() (Bingo, error) {
	rawBingo := inputs.Bingo
	substrings := strings.Split(rawBingo, "\n")

	var bingo Bingo
	rawNumbers := strings.Split(substrings[0], ",")
	for _, i := range rawNumbers {
		n, _ := strconv.Atoi(i)
		bingo.Sequence = append(bingo.Sequence, n)
	}


	var r int
	for i := 2; i < len(substrings) - 4; i+=6 {
		var board Board
		for j := 0; j <= 4; j++ {
			r = i + j
			rawNums := strings.Split(substrings[r], " ")
			var nums []int
			for _, n := range rawNums {
				if n == "" {
					continue
				}
				nu, _ := strconv.Atoi(n)
				nums = append(nums, nu)
			}
			var d [X_SIZE]*Point
			for ix, n := range nums {
				d[ix] = &Point{Value: n}
			}
			board.Points[j] = d
		}
		bingo.Boards = append(bingo.Boards, board)
	}
	return bingo, nil
}

func CreatePuzzle1() *Puzzle1 {
	return &Puzzle1{}
}

type Puzzle1 struct {}

func (p *Puzzle1) Solve() error {
	bingo, err := BingoBoard()
	if err != nil {
		return err
	}
	for _, d := range bingo.Sequence {
		for _, b := range bingo.Boards {
			markBoard(d, b)
			if boardHasWon(b) {
				score := calculateScore(d, b)
				fmt.Printf(">> We have a winner with score: %d\n", score)
				return nil
			}
		}
	}

	return nil
}

func markBoard(number int, board Board) {
	for i := 0; i < Y_SIZE; i++ {
		for j := 0; j < X_SIZE; j++ {
			if board.Points[i][j].Value == number {
				board.Points[i][j].Marked = true
			}
		}
	}
}

func boardHasWon(board Board) bool {
	return checkHorizontal(board) || checkVertical(board)
}

func checkHorizontal(board Board) bool {
	for i := 0; i < Y_SIZE; i++ {
		complete := true
		for j := 0; j < X_SIZE; j++ {
			if !board.Points[i][j].Marked {
				complete = false
				break
			}
		}
		if complete {
			return true
		}
	}
	return false
}

func checkVertical(board Board) bool {
	for i := 0; i < Y_SIZE; i++ {
		complete := true
		for j := 0; j < X_SIZE; j++ {
			if !board.Points[j][i].Marked {
				complete = false
				break
			}
		}
		if complete {
			return true
		}
	}
	return false
}

func calculateScore(number int, board Board) int {
	var sum int
	for i := 0; i < Y_SIZE; i++ {
		for j := 0; j < X_SIZE; j++ {
			if !board.Points[i][j].Marked {
				sum += board.Points[i][j].Value
			}
		}
	}
	return sum * number
}