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

// 実体型でもいいがsortを考慮するとswapが多く行われるため
// ポインタ型の方が効率がいい
var tracks = []*Track{
	{"A", "Z", "1", 2020, length("3m38s")},
	{"A", "Y", "3", 2019, length("5m38s")},
	{"C", "X", "2", 2018, length("2m38s")},
}

// sort 対象の型 []*Track をもつ型を定義
// こいつに sort interfaceのメソッドを実装する
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	st := byArtist(tracks)
	printTracks(st)

	fmt.Println("")
	sort.Sort(st)
	printTracks(st)

	fmt.Println("")
	// Reverse は sort.Interface を返す
	sort.Sort(sort.Reverse(st))
	printTracks(st)

	fmt.Println("")
	cs := customSort{
		tracks,
		func(x, y *Track) bool {
			if x.Title != y.Title {
				return x.Title < y.Title
			}
			if x.Year != y.Year {
				return x.Year < y.Year
			}
			if x.Length != y.Length {
				return x.Length < y.Length
			}
			return false
		}}
	sort.Sort(cs)
	printTracks(cs.t)
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
