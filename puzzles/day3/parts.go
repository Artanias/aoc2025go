package day3

import (
	tools "aoc2025/internal/tools"
	"fmt"
	"strconv"
	"strings"
)

const exampleFilePath string = "puzzles/day3/example.txt"
const dataFilePath string = "puzzles/day3/data.txt"

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
	return res, nil
}

func Run() {
	for i, fun := range []func(string) (int64, error){calcRes, calcRes2} {
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
