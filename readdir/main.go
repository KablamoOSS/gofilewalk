package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	fmt.Println(readDir("./") + 1)
}

func readDir(name string) int {
	d, err := os.Open(name)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	defer d.Close()

	files, err := d.ReadDir(-1)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	count := len(files) // account for this one

	for _, f := range files {
		if f.IsDir() {
			count += readDir(path.Join(name, f.Name()))
		}
	}

	return count
}

