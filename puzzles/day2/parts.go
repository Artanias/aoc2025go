package day2

import (
	tools "aoc2025/internal/tools"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

const exampleFilePath string = "puzzles/day2/example.txt"
const dataFilePath string = "puzzles/day2/data.txt"

func cntDigits(number int64) (cnt int64) {
	for ; number != 0; number /= 10 {
		cnt += 1
	}
	return
}

func calcDivs(number int64) []int64 {
	res := make([]int64, 0)
	for i := int64(2); i < number; i++ {
		if number%i == 0 {
			res = append(res, i)
		}
	}
	res = append(res, number)
	return res
}

func calcRes(content string) (int64, error) {
	ranges := strings.Split(content, ",")
	var res int64
	for _, r := range ranges {
		rangeNums := strings.Split(r, "-")
		first, err := strconv.ParseInt(rangeNums[0], 10, 64)
		if err != nil {
			return 0, err
		}
		second, err := strconv.ParseInt(rangeNums[1], 10, 64)
		if err != nil {
			return 0, err
		}
		for i := first; i <= second; i++ {
			digits := cntDigits(i)
			if digits%2 != 0 {
				continue
			}
			degree := int64(math.Pow10(int(digits / 2)))
			left := i / degree
			right := i % degree
			if left == right {
				res += i
			}
		}
	}
	return res, nil
}

func calcRes2(content string) (int64, error) {
	ranges := strings.Split(content, ",")
	var res int64
	invalidIds := make([]int64, 0)
	for _, r := range ranges {
		rangeNums := strings.Split(r, "-")
		first, err := strconv.ParseInt(rangeNums[0], 10, 64)
		if err != nil {
			return 0, err
		}
		second, err := strconv.ParseInt(rangeNums[1], 10, 64)
		if err != nil {
			return 0, err
		}
		for i := first; i <= second; i++ {
			digits := cntDigits(i)
			divs := calcDivs(digits)
			if i < 10 {
				continue
			}
			for _, div := range divs {
				j := i
				degree := int64(math.Pow10(int(digits / div)))
				parts := make([]int64, 0)
				for j != 0 {
					parts = append(parts, j%degree)
					j /= degree
				}
				first := parts[0]
				equal := true
				for k := 1; k < len(parts); k++ {
					if first != parts[k] {
						equal = false
						break
					}
				}
				if equal {
					if !slices.Contains(invalidIds, i) {
						invalidIds = append(invalidIds, i)
						res += i
					}
					break
				}
			}
		}
	}
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
