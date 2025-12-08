package day4

import (
	tools "aoc2025/internal/tools"
	"aoc2025/puzzles"
	"path/filepath"
)

var puzzlePath string = filepath.Join(puzzles.PuzzlePath, "day4")
var exampleFilePath string = filepath.Join(puzzlePath, "example.txt")
var dataFilePath string = filepath.Join(puzzlePath, "data.txt")

func calcRes(content string) (int64, error) {
	var res int64
	field := tools.MakeField(content)
	for i := 0; i < field.Rows; i++ {
		for j := 0; j < field.Columns; j++ {
			if field.F[i][j] == string("@") {
				near := field.GetNear(i, j)
				if near["@"] < 4 {
					res += 1
				}
			}
		}
	}
	return res, nil
}

func calcRes2(content string) (int64, error) {
	var res, lastRes int64
	field := tools.MakeField(content)
	for {
		for i := 0; i < field.Rows; i++ {
			for j := 0; j < field.Columns; j++ {
				if field.F[i][j] == string("@") {
					near := field.GetNear(i, j)
					if near["@"] < 4 {
						lastRes += 1
						field.F[i][j] = "x"
					}
				}
			}
		}
		if lastRes == 0 {
			break
		}
		res += lastRes
		lastRes = 0
	}
	return res, nil
}

func Run() {
	tools.Run(
		[]func(string) (int64, error){calcRes, calcRes2},
		[]string{exampleFilePath, dataFilePath},
		[]int64{13, 1393, 43, 8643},
	)
}
