package tools

import (
	"fmt"
	"strings"
)

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

func MakeField(content string) *Field {
	lines := strings.Split(content, "\n")
	rows := len(lines)
	columns := len(lines[0])
	field := make([][]string, rows, rows)
	for i := 0; i < rows; i++ {
		field[i] = make([]string, columns, columns)
		for j := 0; j < columns; j++ {
			field[i][j] = string(lines[i][j])
		}
	}
	return &Field{field, rows, columns}
}
