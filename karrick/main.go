package main

import (
	"fmt"
	"github.com/karrick/godirwalk"
	"sync/atomic"
)

func main() {
	var count int64
	godirwalk.Walk("./", &godirwalk.Options{
		Unsorted: true,
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			atomic.AddInt64(&count, 1)
			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			return godirwalk.SkipNode
		},
	})
	fmt.Println(count)
}
