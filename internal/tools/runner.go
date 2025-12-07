package tools

import (
	"fmt"
	"path/filepath"
)

func Run(funcs []func(string) (int64, error), paths []string) {
	for i, fun := range funcs {
		fmt.Printf("=====Part %d=====\n", i+1)
		for _, path := range paths {
			res, err := fun(GetFileContent(path))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("'%s' result: %d\n", filepath.Base(path), res)
		}
		fmt.Println("================")
		fmt.Println()
	}
}
