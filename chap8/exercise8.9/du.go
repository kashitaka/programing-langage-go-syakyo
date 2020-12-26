package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

type rootSize struct {
	path string
	size int64
}

func walkDir(root, dir string, wg *sync.WaitGroup, fileSizes chan<- rootSize) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(root, subdir, wg, fileSizes)
		} else {
			fileSizes <- rootSize{root, entry.Size()}
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
	fileSizes := make(chan rootSize)
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, root, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	nfiles := map[string]int64{}
	nbytes := map[string]int64{}
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles[size.path]++
			nbytes[size.path] += size.size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes map[string]int64) {
	var buf bytes.Buffer
	keys := []string{}
	for k := range nfiles {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		buf.WriteString(
			fmt.Sprintf("%s: %d files %.1f GB\n", k, nfiles[k], float64(nbytes[k])/1e9))
	}
	buf.WriteString("\033[")
	buf.WriteString(fmt.Sprintf("%d", len(keys)))
	buf.WriteString("A")
	fmt.Fprintf(os.Stdout, buf.String())
}
