package day7

import (
	tools "aoc2025/internal/tools"
	"aoc2025/puzzles"
	"fmt"
	"path/filepath"
)

var puzzlePath string = filepath.Join(puzzles.PuzzlePath, "day7")
var exampleFilePath string = filepath.Join(puzzlePath, "example.txt")
var dataFilePath string = filepath.Join(puzzlePath, "data.txt")

func calcRes(content string) (int64, error) {
	field := tools.MakeField(content)
	positions := field.FindPositions("S")
	uniqueSplits := make([]tools.Point2D, 0)
	if len(positions) != 1 {
		return 0, fmt.Errorf("invalid number of S positions in example file - %d", len(positions))
	}
	for len(positions) != 0 {
		fmt.Printf("%d     %v\r", len(positions), positions[0])
		point := positions[len(positions)-1]
		positions = positions[:len(positions)-1]
		if point.X+1 >= field.Rows {
			continue
		}
		if point.Y+1 >= field.Columns && point.Y-1 < 0 {
			continue
		}
		if field.F[point.X+1][point.Y] == "^" {
			uniq := true
			for _, uniqPoint := range uniqueSplits {
				if uniqPoint == point {
					uniq = false
					break
				}
			}
			if uniq {
				uniqueSplits = append(uniqueSplits, point)
				positions = append(positions, tools.Point2D{X: point.X + 1, Y: point.Y - 1})
				positions = append(positions, tools.Point2D{X: point.X + 1, Y: point.Y + 1})
			}
		} else {
			positions = append(positions, tools.Point2D{X: point.X + 1, Y: point.Y})
		}
	}
	return int64(len(uniqueSplits)), nil
}

func calcRes2(content string) (int64, error) {
	var res int64
	field := tools.MakeField(content)
	positions := field.FindPositions("S")
	if len(positions) != 1 {
		return 0, fmt.Errorf("invalid number of S positions in example file - %d", len(positions))
	}
	memory := make(map[tools.Point2D]int64, 0)
	for len(positions) != 0 {
		point := positions[0]
		positions = positions[1:]
		_, ok := memory[point]
		if ok {
			continue
		}
		if point.X-1 < 0 {
			memory[point] = 1
		}
		if point.X == field.Rows {
			continue
		}
		if point.X+1 < field.Rows && field.F[point.X+1][point.Y] == "^" {
			positions = append(positions, tools.Point2D{X: point.X + 1, Y: point.Y - 1})
			positions = append(positions, tools.Point2D{X: point.X + 1, Y: point.Y + 1})
		} else {
			positions = append(positions, tools.Point2D{X: point.X + 1, Y: point.Y})
		}
		if point.X-1 < 0 {
			continue
		}
		if point.Y-1 >= 0 && field.F[point.X][point.Y-1] == "^" {
			v, ok := memory[tools.Point2D{X: point.X - 1, Y: point.Y - 1}]
			if ok {
				memory[point] += v
			}
		}
		if point.Y+1 < field.Columns && field.F[point.X][point.Y+1] == "^" {
			v, ok := memory[tools.Point2D{X: point.X - 1, Y: point.Y + 1}]
			if ok {
				memory[point] += v
			}
		}
		v, ok := memory[tools.Point2D{X: point.X - 1, Y: point.Y}]
		if ok {
			memory[point] += v
		}
	}
	for j := 0; j < field.Columns; j++ {
		res += memory[tools.Point2D{X: field.Rows - 1, Y: j}]
	}
	return res, nil
}

func Run() {
	tools.Run(
		[]func(string) (int64, error){calcRes, calcRes2},
		[]string{exampleFilePath, dataFilePath},
		[]int64{21, 1649, 40, 16937871060075},
	)
}
