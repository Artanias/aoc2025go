package day6

import (
	tools "aoc2025/internal/tools"
	"aoc2025/puzzles"
	"path/filepath"
	"strconv"
	"strings"
)

var puzzlePath string = filepath.Join(puzzles.PuzzlePath, "day6")
var exampleFilePath string = filepath.Join(puzzlePath, "example.txt")
var dataFilePath string = filepath.Join(puzzlePath, "data.txt")

func calcRes(content string) (int64, error) {
	var res int64
	lines := strings.Split(content, "\n")
	ops := strings.Split(strings.ReplaceAll(lines[len(lines)-1], " ", ""), "")
	results := make([]int64, len(ops), len(ops))
	for _, line := range lines[:len(lines)-1] {
		numbers := strings.Fields(line)
		for i, num := range numbers {
			numInt, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				return 0, err
			}
			if ops[i] == "*" && results[i] == 0 {
				results[i] = numInt
			} else {
				switch ops[i] {
				case "*":
					results[i] *= numInt
				case "+":
					results[i] += numInt
				}
			}
		}
	}
	for _, result := range results {
		res += result
	}
	return res, nil
}

func calcRes2(content string) (int64, error) {
	var res int64
	lines := strings.Split(content, "\n")
	rows := len(lines)
	ops := strings.Split(strings.ReplaceAll(lines[rows-1], " ", ""), "")
	results := make([]int64, 0, len(ops))
	field := tools.MakeField(content)
	op := ""
	for j := 0; j < field.Columns; j++ {
		number := ""
		for i := 0; i < field.Rows-1; i++ {
			if field.F[i][j] != " " {
				number += field.F[i][j]
			}
		}
		if number == "" {
			continue
		}
		if field.F[field.Rows-1][j] != " " {
			op = field.F[field.Rows-1][j]
			switch op {
			case "*":
				results = append(results, 1)
			case "+":
				results = append(results, 0)
			}
		}
		numberInt, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			return 0, err
		}
		switch op {
		case "*":
			results[len(results)-1] *= numberInt
		case "+":
			results[len(results)-1] += numberInt
		}
	}
	for _, result := range results {
		res += result
	}
	return res, nil
}

func Run() {
	tools.Run(
		[]func(string) (int64, error){calcRes, calcRes2},
		[]string{exampleFilePath, dataFilePath},
		[]int64{4277556, 5524274308182, 3263827, 8843673199391},
	)
}
