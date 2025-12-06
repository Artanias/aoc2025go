package day1

import (
	tools "aoc2025/internal/tools"
	"fmt"
	"strconv"
	"strings"
)

const exampleFilePath string = "puzzles/day1/example.txt"
const dataFilePath string = "puzzles/day1/data.txt"

func calcRes(content string) (int, error) {
	startPos := 50
	res := 0
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

func calcRes2(content string) (int, error) {
	startPos := 50
	res := 0
	for _, line := range strings.Split(content, "\n") {
		rotation := string(line[0])
		positions, err := strconv.Atoi(line[1:])
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
	for i, fun := range []func(string) (int, error){calcRes, calcRes2} {
		fmt.Printf("=====Part %d=====\n", i+1)
		res, err := fun(tools.GetFileContent(exampleFilePath))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Example result: %d\n", res)
		res, err = fun(tools.GetFileContent(dataFilePath))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Data result: %d\n", res)
		fmt.Println("================")
		fmt.Println()
	}
}
