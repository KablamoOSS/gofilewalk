package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"
)

func main() {
	var count int64
	filepath.WalkDir("./", func(root string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		atomic.AddInt64(&count, 1)
		return nil
	})

	fmt.Println(count)
}
