package tools

import (
	"fmt"
	"os"
	"path/filepath"
)

func Run(funcs []func(string) (int64, error), paths []string, results []int64) {
	for i, fun := range funcs {
		fmt.Printf("=====Part %d=====\n", i+1)
		for j, path := range paths {
			res, err := fun(GetFileContent(path))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("'%s' result: %d\n", filepath.Base(path), res)
			expected := results[i*len(funcs)+j]
			if res != expected {
				fmt.Printf("ERROR: Expected value: %d; Got: %d.\n", expected, res)
				os.Exit(1)
			}
		}
		fmt.Println("================")
		fmt.Println()
	}
}
