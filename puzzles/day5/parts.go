package day5

import (
	tools "aoc2025/internal/tools"
	"aoc2025/puzzles"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
)

var puzzlePath string = filepath.Join(puzzles.PuzzlePath, "day5")
var exampleFilePath string = filepath.Join(puzzlePath, "example.txt")
var dataFilePath string = filepath.Join(puzzlePath, "data.txt")

type Range struct {
	start int64
	end   int64
}

func calcRes(content string) (int64, error) {
	var res int64
	parts := strings.Split(content, "\n\n")
	if len(parts) != 2 {
		return 0, fmt.Errorf("got invalid content with %d parts", len(parts))
	}
	ranges := make([]Range, 0, len(parts[0]))
	for _, line := range strings.Split(parts[0], "\n") {
		lineParts := strings.Split(line, "-")
		if len(lineParts) != 2 {
			return 0, fmt.Errorf("got invalid line '%s' content", line)
		}
		start, err := strconv.ParseInt(lineParts[0], 10, 64)
		if err != nil {
			return 0, nil
		}
		end, err := strconv.ParseInt(lineParts[1], 10, 64)
		if err != nil {
			return 0, nil
		}
		ranges = append(ranges, Range{start, end})
	}
	for _, line := range strings.Split(parts[1], "\n") {
		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return 0, nil
		}
		for _, r := range ranges {
			if value >= r.start && value <= r.end {
				res++
				break
			}
		}
	}
	return res, nil
}

func calcRes2(content string) (int64, error) {
	var res int64
	return res, nil
}

func Run() {
	tools.Run(
		[]func(string) (int64, error){calcRes, calcRes2},
		[]string{exampleFilePath, dataFilePath},
		[]int64{3, 712, 0, 0},
	)
}
