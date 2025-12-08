package day3

import (
	tools "aoc2025/internal/tools"
	"aoc2025/puzzles"
	"path/filepath"
	"strconv"
	"strings"
)

var puzzlePath string = filepath.Join(puzzles.PuzzlePath, "day3")
var exampleFilePath string = filepath.Join(puzzlePath, "example.txt")
var dataFilePath string = filepath.Join(puzzlePath, "data.txt")

func calcRes(content string) (int64, error) {
	var res int64
	for _, line := range strings.Split(content, "\n") {
		var max int64
		for i := 0; i < len(line); i++ {
			first := string(line[i])
			for j := i + 1; j < len(line); j++ {
				second := string(line[j])
				number, err := strconv.ParseInt(first+second, 10, 64)
				if err != nil {
					return 0, err
				}
				if number > max {
					max = number
				}
			}
		}
		res += max
	}
	return res, nil
}

func calcRes2(content string) (int64, error) {
	var res int64
	const maxId int = 12
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		indexes := make([]int, 0, maxId)
		var i int
		for len(indexes) != maxId {
			var availableSpace int = len(line) - i - (maxId - len(indexes)) + 1
			var maxIndex int
			var maxNumber int
			for j := i; j < i+availableSpace; j++ {
				number, err := strconv.Atoi(string(line[j]))
				if err != nil {
					return 0, err
				}
				if number > maxNumber {
					maxIndex = j
					maxNumber = number
				}
			}
			indexes = append(indexes, maxIndex)
			i = maxIndex + 1
		}
		numberStr := ""
		for j := 0; j < maxId; j++ {
			numberStr = numberStr + string(line[indexes[j]])
		}
		number, err := strconv.ParseInt(numberStr, 10, 64)
		if err != nil {
			return 0, err
		}
		res += number
	}
	return res, nil
}

func Run() {
	tools.Run(
		[]func(string) (int64, error){calcRes, calcRes2},
		[]string{exampleFilePath, dataFilePath},
		[]int64{357, 17766, 3121910778619, 176582889354075},
	)
}
