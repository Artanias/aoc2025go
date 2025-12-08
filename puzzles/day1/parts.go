package day1

import (
	tools "aoc2025/internal/tools"
	"aoc2025/puzzles"
	"path/filepath"
	"strconv"
	"strings"
)

var puzzlePath string = filepath.Join(puzzles.PuzzlePath, "day1")
var exampleFilePath string = filepath.Join(puzzlePath, "example.txt")
var dataFilePath string = filepath.Join(puzzlePath, "data.txt")

func calcRes(content string) (int64, error) {
	startPos := 50
	var res int64
	for _, line := range strings.Split(content, "\n") {
		rotation := string(line[0])
		positions, err := strconv.Atoi(line[1:])
		if err != nil {
			return 0, err
		}
		if rotation == "L" {
			startPos -= positions
			startPos %= 100
		} else {
			startPos += positions
			startPos %= 100
		}
		if startPos == 0 {
			res += 1
		}
	}
	return res, nil
}

func calcRes2(content string) (int64, error) {
	var startPos int64 = 50
	var res int64
	for _, line := range strings.Split(content, "\n") {
		rotation := string(line[0])
		positions, err := strconv.ParseInt(line[1:], 10, 64)
		if err != nil {
			return 0, err
		}
		res += positions / 100
		positions %= 100
		if rotation == "L" {
			if startPos == 0 {
				startPos += 100
			}
			startPos -= positions
			if startPos < 0 {
				res += 1
				startPos += 100
			}
		} else {
			startPos += positions
			if startPos > 100 {
				res += 1
			}
			startPos %= 100
		}
		if startPos == 0 {
			res += 1
		}
	}
	return res, nil
}

func Run() {
	tools.Run(
		[]func(string) (int64, error){calcRes, calcRes2},
		[]string{exampleFilePath, dataFilePath},
		[]int64{3, 1097, 6, 7101},
	)
}
