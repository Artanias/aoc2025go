package day8

import (
	tools "aoc2025/internal/tools"
	"aoc2025/puzzles"
	"fmt"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

var puzzlePath string = filepath.Join(puzzles.PuzzlePath, "day8")
var exampleFilePath string = filepath.Join(puzzlePath, "example.txt")
var dataFilePath string = filepath.Join(puzzlePath, "data.txt")

func findMinDistPoints(condition bool, points []tools.Point3D) chan tools.PointsPair {
	pointsPairs := make([]tools.PointsPair, 0)
	c := make(chan tools.PointsPair, 1)

	go func() {
		for condition {
			minDist := 0
			var minPoint1, minPoint2 tools.Point3D
			for _, p1 := range points {
				for _, p2 := range points {
					if p1 == p2 {
						continue
					}
					dist := p1.Distance(p2)
					if (minDist == 0 || minDist > dist) && !slices.Contains(
						pointsPairs, tools.PointsPair{P1: p1, P2: p2},
					) && !slices.Contains(
						pointsPairs, tools.PointsPair{P1: p2, P2: p1},
					) {
						minDist = dist
						minPoint1 = p1
						minPoint2 = p2
					}
				}
			}
			pointsPair := tools.PointsPair{P1: minPoint1, P2: minPoint2}
			c <- pointsPair
			pointsPairs = append(pointsPairs, pointsPair)
		}
		close(c)
	}()

	return c
}

func deepContains(point tools.Point3D, circuits [][]tools.Point3D) int {
	for i, circuit := range circuits {
		if slices.Contains(circuit, point) {
			return i
		}
	}
	return -1
}

func calcRes(content string, minPointsCnt int) (int64, error) {
	var res int64
	points := make([]tools.Point3D, 0)
	for line := range strings.SplitSeq(content, "\n") {
		numbers := strings.Split(line, ",")
		if len(numbers) != 3 {
			return 0, fmt.Errorf("invalid line: %s", line)
		}
		x, err := strconv.Atoi(numbers[0])
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(numbers[1])
		if err != nil {
			return 0, err
		}
		z, err := strconv.Atoi(numbers[2])
		if err != nil {
			return 0, err
		}
		points = append(points, tools.Point3D{X: x, Y: y, Z: z})
	}
	circuits := make([][]tools.Point3D, 0)
	condition := true
	k := 0
	for pointsPair := range findMinDistPoints(condition, points) {
		if k == minPointsCnt {
			condition = false
			break
		}
		i := deepContains(pointsPair.P1, circuits)
		j := deepContains(pointsPair.P2, circuits)
		k++
		if i != -1 && j != -1 && i == j {
			continue
		} else if i != -1 && j != -1 {
			circuits[i] = append(circuits[i], circuits[j]...)
			circuits = slices.Delete(circuits, j, j+1)
		} else if i != -1 {
			circuits[i] = append(circuits[i], pointsPair.P2)
		} else if j != -1 {
			circuits[j] = append(circuits[j], pointsPair.P1)
		} else {
			circuit := make([]tools.Point3D, 0)
			circuit = append(circuit, pointsPair.P1, pointsPair.P2)
			circuits = append(circuits, circuit)
		}
	}
	circuitsLens := make([]int64, 0)
	for _, circircuit := range circuits {
		circuitsLens = append(circuitsLens, int64(len(circircuit)))
	}
	slices.Sort(circuitsLens)
	slices.Reverse(circuitsLens)
	res = circuitsLens[0] * circuitsLens[1] * circuitsLens[2]
	return res, nil
}

func calcRes2(content string) (int64, error) {
	var res int64
	points := make([]tools.Point3D, 0)
	for line := range strings.SplitSeq(content, "\n") {
		numbers := strings.Split(line, ",")
		if len(numbers) != 3 {
			return 0, fmt.Errorf("invalid line: %s", line)
		}
		x, err := strconv.Atoi(numbers[0])
		if err != nil {
			return 0, err
		}
		y, err := strconv.Atoi(numbers[1])
		if err != nil {
			return 0, err
		}
		z, err := strconv.Atoi(numbers[2])
		if err != nil {
			return 0, err
		}
		points = append(points, tools.Point3D{X: x, Y: y, Z: z})
	}
	circuits := make([][]tools.Point3D, 0)
	condition := true
	for pointsPair := range findMinDistPoints(condition, points) {
		i := deepContains(pointsPair.P1, circuits)
		j := deepContains(pointsPair.P2, circuits)
		if i != -1 && j != -1 && i == j {
			continue
		} else if i != -1 && j != -1 {
			circuits[i] = append(circuits[i], circuits[j]...)
			circuits = slices.Delete(circuits, j, j+1)
		} else if i != -1 {
			circuits[i] = append(circuits[i], pointsPair.P2)
		} else if j != -1 {
			circuits[j] = append(circuits[j], pointsPair.P1)
		} else {
			circuit := make([]tools.Point3D, 0)
			circuit = append(circuit, pointsPair.P1, pointsPair.P2)
			circuits = append(circuits, circuit)
		}
		if len(circuits) == 1 && len(circuits[0]) == len(points) {
			res = int64(pointsPair.P1.X) * int64(pointsPair.P2.X)
			condition = false
			break
		}
	}
	return res, nil
}

func runner() func(content string) (int64, error) {
	i := 0
	minPointsCnts := [2]int{10, 1000}
	return func(content string) (int64, error) {
		res, err := calcRes(content, minPointsCnts[i])
		i++
		return res, err
	}
}

func Run() {
	tools.Run(
		[]func(string) (int64, error){runner(), calcRes2},
		[]string{exampleFilePath, dataFilePath},
		[]int64{40, 112230, 25272, 2573952864},
	)
}
