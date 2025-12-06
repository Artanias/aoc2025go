package main

import (
	day1 "aoc2025/puzzles/day1"
	day2 "aoc2025/puzzles/day2"
)

func main() {
	for _, f := range []func(){day1.Run, day2.Run} {
		f()
	}
}
