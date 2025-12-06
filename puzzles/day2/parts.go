package day2

import (
	tools "aoc2025/internal/tools"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const exampleFilePath string = "puzzles/day2/example.txt"
const dataFilePath string = "puzzles/day2/data.txt"

func cntDigits(number int) (cnt int) {
	for ; number != 0; number /= 10 {
		cnt += 1
	}
	return
}

func calcRes(content string) (int, error) {
	ranges := strings.Split(content, ",")
	res := 0
	for _, r := range ranges {
		rangeNums := strings.Split(r, "-")
		first, err := strconv.Atoi(rangeNums[0])
		if err != nil {
			return 0, err
		}
		second, err := strconv.Atoi(rangeNums[1])
		if err != nil {
			return 0, err
		}
		for i := first; i <= second; i++ {
			digits := cntDigits(i)
			if digits%2 != 0 {
				continue
			}
			degree := int(math.Pow10(digits / 2))
			left := i / degree
			right := i % degree
			if left == right {
				res += i
			}
		}
	}
	return res, nil
}

func Run() {
	for i, fun := range []func(string) (int, error){calcRes} {
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
