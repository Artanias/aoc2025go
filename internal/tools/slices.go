package tools

import (
	"fmt"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Field struct {
	F       [][]string
	Rows    int
	Columns int
}

func (f *Field) PrintField() {
	for i := 0; i < f.Rows; i++ {
		for j := 0; j < f.Columns; j++ {
			fmt.Print(f.F[i][j])
		}
		fmt.Println()
	}
}

func (f *Field) GetNear(row int, column int) map[string]int {
	result := make(map[string]int, 3)
	if row > 0 {
		if column > 0 {
			result[f.F[row-1][column-1]] += 1
		}
		if column < f.Columns-1 {
			result[f.F[row-1][column+1]] += 1
		}
		result[f.F[row-1][column]] += 1
	}
	if row < f.Rows-1 {
		if column > 0 {
			result[f.F[row+1][column-1]] += 1
		}
		if column < f.Columns-1 {
			result[f.F[row+1][column+1]] += 1
		}
		result[f.F[row+1][column]] += 1
	}
	if column > 0 {
		result[f.F[row][column-1]] += 1
	}
	if column < f.Columns-1 {
		result[f.F[row][column+1]] += 1
	}
	return result
}

func (f *Field) FindPositions(value string) []Point {
	positions := make([]Point, 0)
	for row := 0; row < f.Rows; row++ {
		for column := 0; column < f.Columns; column++ {
			if f.F[row][column] == value {
				positions = append(positions, Point{X: row, Y: column})
			}
		}
	}
	return positions
}

func MakeField(content string) *Field {
	lines := strings.Split(content, "\n")
	rows := len(lines)
	columns := len(lines[0])
	for _, line := range lines {
		if len(line) > columns {
			columns = len(line)
		}
	}
	field := make([][]string, rows, rows)
	for i := 0; i < rows; i++ {
		field[i] = make([]string, columns, columns)
		for j := 0; j < columns; j++ {
			if j >= len(lines[i]) {
				field[i][j] = " "
			} else {
				field[i][j] = string(lines[i][j])
			}
		}
	}
	return &Field{field, rows, columns}
}
