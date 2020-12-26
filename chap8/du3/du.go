package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done() // こういうdeferの使い方あり
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// direntsの並列数を制御する計数セマフォ
var sema = make(chan struct{}, 100)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // tokenを取得
	defer func() { <-sema }() // tokenを開放
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

var verbose = flag.Bool("v", false, "show verbose progress message")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var wg sync.WaitGroup
	fileSizes := make(chan int64)
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
