package main

import (
	"fmt"
	"github.com/MichaelTJones/walk"
	"os"
	"sync"
)

func main() {
	count := 0
	var mutex sync.Mutex

	walk.Walk("./", func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		mutex.Lock()
		count++
		mutex.Unlock()

		return nil
	})

	fmt.Println(count)
}