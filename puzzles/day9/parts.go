package day9

import (
	tools "aoc2025/internal/tools"
	"aoc2025/puzzles"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

var puzzlePath string = filepath.Join(puzzles.PuzzlePath, "day9")
var exampleFilePath string = filepath.Join(puzzlePath, "example.txt")
var dataFilePath string = filepath.Join(puzzlePath, "data.txt")

func calcRes(content string) (int64, error) {
	points := make([]tools.Point2D, 0)
	for _, line := range strings.Split(content, "\n") {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			return 0, fmt.Errorf("invalid line: %s", line)
		}
		column, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
		row, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
		points = append(points, tools.Point2D{X: row, Y: column})
	}
	var maxSquare int64 = 0
	for _, p1 := range points {
		for _, p2 := range points {
			if p1 == p2 {
				continue
			}
			square := int64(p1.Square(p2))
			if square > maxSquare {
				maxSquare = square
			}
		}
	}
	return int64(maxSquare), nil
}

func calcRes2(content string) (int64, error) {
	var res int64
	return res, nil
}

func Run() {
	tools.Run(
		[]func(string) (int64, error){calcRes, calcRes2},
		[]string{exampleFilePath, dataFilePath},
		[]int64{50, 4738108384, 0, 0},
	)
}
