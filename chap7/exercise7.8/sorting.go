package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type trackKey int

const (
	title trackKey = iota
	artist
	album
	year
	leng
)

// 実体型でもいいがsortを考慮するとswapが多く行われるため
// ポインタ型の方が効率がいい
var tracks = []*Track{
	{"A", "X", "1", 2020, length("3m38s")},
	{"A", "X", "3", 2019, length("2m38s")},
	{"C", "X", "2", 2018, length("2m38s")},
}

type customSort struct {
	t        []*Track
	new, old trackKey
}

func (x customSort) Len() int      { return len(x.t) }
func (x customSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x customSort) Less(i, j int) bool {
	sortK1 := lessKey(x.t[i], x.t[j], x.new)
	if sortK1 {
		return sortK1
	}
	return lessKey(x.t[i], x.t[j], x.old)
}

func lessKey(x, y *Track, sortkey trackKey) bool {
	switch sortkey {
	case title:
		return x.Title < y.Title
	case artist:
		return x.Artist < y.Artist
	case album:
		return x.Album < y.Album
	case year:
		return x.Year < y.Year
	case leng:
		return x.Length < y.Length
	}
	return false
}

func sortByKey(c *customSort, key trackKey) {
	c.old = c.new
	c.new = key
	sort.Sort(c)
}

func main() {
	c := customSort{t: tracks}
	sortByKey(&c, leng)
	printTracks(c.t)

	fmt.Println("")
	sortByKey(&c, artist)
	printTracks(c.t)

	fmt.Println("")
	sortByKey(&c, leng)
	printTracks(c.t)
}

// ただのヘルパー関数
func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, v := range tracks {
		fmt.Fprintf(tw, format, v.Title, v.Artist, v.Album, v.Year, v.Length)
	}
	tw.Flush()
}
