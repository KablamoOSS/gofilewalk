package main

import (
	"fmt"
	"os"
	"path"
	"sync/atomic"
)

func main() {
	fmt.Println(readDir("./") + 1) // account for the fact that ./ is not counted
}

func readDir(name string) int64 {
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

	count := int64(len(files))

	for _, f := range files {
		if f.IsDir() {
			atomic.AddInt64(&count, readDir(path.Join(name, f.Name())))
		}
	}

	return count
}
