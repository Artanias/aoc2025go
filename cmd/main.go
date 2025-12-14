package main

import (
	day1 "aoc2025/puzzles/day1"
	day2 "aoc2025/puzzles/day2"
	day3 "aoc2025/puzzles/day3"
	day4 "aoc2025/puzzles/day4"
	day5 "aoc2025/puzzles/day5"
	day6 "aoc2025/puzzles/day6"
	day7 "aoc2025/puzzles/day7"
	day8 "aoc2025/puzzles/day8"
	day9 "aoc2025/puzzles/day9"
	"fmt"
)

func main() {
	for i, f := range []func(){
		day1.Run,
		day2.Run,
		day3.Run,
		day4.Run,
		day5.Run,
		day6.Run,
		day7.Run,
		day8.Run,
		day9.Run,
	} {
		fmt.Println("=======", "DAY", i+1, "RUNNING", "=======")
		f()
	}
}
