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

func calcRes(content string) (int64, error) {
	var res int64
	parts := strings.Split(content, "\n\n")
	if len(parts) != 2 {
		return 0, fmt.Errorf("got invalid content with %d parts", len(parts))
	}
	ranges := make([]tools.Range, 0, len(parts[0]))
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
		ranges = append(ranges, tools.Range{Start: start, End: end})
	}
	for _, line := range strings.Split(parts[1], "\n") {
		value, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return 0, nil
		}
		for _, r := range ranges {
			if value >= r.Start && value <= r.End {
				res++
				break
			}
		}
	}
	return res, nil
}

func calcRes2(content string) (int64, error) {
	var res int64
	parts := strings.Split(content, "\n\n")
	if len(parts) != 2 {
		return 0, fmt.Errorf("got invalid content with %d parts", len(parts))
	}
	ranges := make([]tools.Range, 0, len(parts[0]))
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
		validateRanges := make([]tools.Range, 0, 2)
		validateRanges = append(validateRanges, tools.Range{Start: start, End: end})
		var vr tools.Range
		for len(validateRanges) != 0 {
			validateRanges, vr = validateRanges[:len(validateRanges)-1], validateRanges[len(validateRanges)-1]
			skip := false
			for _, r := range ranges {
				if vr.Start >= r.Start && vr.End <= r.End {
					skip = true
					break
				}
				if vr.Start < r.Start && vr.End > r.End {
					left := tools.Range{Start: vr.Start, End: r.Start - 1}
					right := tools.Range{Start: r.End + 1, End: vr.End}
					validateRanges = append(validateRanges, left, right)
					skip = true
					break
				}
				if vr.Start >= r.Start && vr.Start <= r.End && vr.End > r.End {
					vr.Start = r.End + 1
				}
				if vr.End <= r.End && vr.End >= r.Start && vr.Start < r.Start {
					vr.End = r.Start - 1
				}
			}
			if skip {
				continue
			}
			ranges = append(ranges, vr)
		}
	}
	for _, r := range ranges {
		res += r.Len()
	}
	return res, nil
}

func Run() {
	tools.Run(
		[]func(string) (int64, error){calcRes, calcRes2},
		[]string{exampleFilePath, dataFilePath},
		[]int64{3, 712, 14, 332998283036769},
	)
}
