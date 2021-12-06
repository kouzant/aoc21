package day2

import (
	"aoc21/inputs"
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	Type string
	Unit int
}

func PlannedCourse() ([]Command, error) {
	rawCourse := strings.Split(inputs.PlannedCourse, "\n")
	var course []Command
	for _, c := range rawCourse {
		cp := strings.Split(c, " ")
		u, _ := strconv.Atoi(cp[1]);
		command := Command{strings.ToUpper(cp[0]), u}
		course = append(course, command)
	}
	if len(course) == 0 {
		return nil, fmt.Errorf("oops no planned course")
	}
	return course, nil
}

func CreatePuzzle1() *Puzzle1 {
	return &Puzzle1{}
}

type Puzzle1 struct {}

func (p *Puzzle1) Solve() error {
	course, err := PlannedCourse()
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
