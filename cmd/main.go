package main

import (
	day1 "aoc2025/puzzles/day1"
	day2 "aoc2025/puzzles/day2"
	day3 "aoc2025/puzzles/day3"
)

func main() {
	for _, f := range []func(){day1.Run, day2.Run, day3.Run} {
		f()
	}
}
