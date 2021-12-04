package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func DepthReadings() ([]int, error) {
	r, err := readFile("inputs/depth_readings")
	if err != nil {
		return nil, err
	}
	rs := strings.Split(string(r), "\n")
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

type Command struct {
	Type string
	Unit int
}

func PlannedCourse() ([]Command, error) {
	r, err := readFile("inputs/planned_course")
	if err != nil {
		return nil, err
	}
	rawCourse := strings.Split(string(r), "\n")
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

func Diagnostics() ([]string, error) {
	r, err := readFile("inputs/diagnostics")
	if err != nil {
		return nil, err
	}
	rawDiagnostics := strings.Split(string(r), "\n")
	var diagnostics []string
	for _, d := range rawDiagnostics {
		diagnostics = append(diagnostics, d)
	}
	if len(diagnostics) == 0 {
		return nil, fmt.Errorf("oops no diagnostics")
	}
	return diagnostics, nil
}

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
	_, err := readFile("inputs/bingo")
	if err != nil {
		return Bingo{}, err
	}

	var bingo Bingo
	return bingo, nil
}

func readFile(path string) ([]byte, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	r, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	return r, nil
}