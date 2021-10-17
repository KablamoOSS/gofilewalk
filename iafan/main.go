package main

import (
	"fmt"
	"github.com/iafan/cwalk"
	"os"
	"sync/atomic"
)

func main() {
	var count int64

	cwalk.Walk("./", func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		atomic.AddInt64(&count, 1)
		return nil
	})
	fmt.Println(count)
}
