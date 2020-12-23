package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
	"text/template"
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
	T        []*Track
	new, old trackKey
}

func (x customSort) Len() int      { return len(x.T) }
func (x customSort) Swap(i, j int) { x.T[i], x.T[j] = x.T[j], x.T[i] }
func (x customSort) Less(i, j int) bool {
	sortK1 := lessKey(x.T[i], x.T[j], x.new)
	if sortK1 {
		return sortK1
	}
	return lessKey(x.T[i], x.T[j], x.old)
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

const templ = `
<html><body>
<table>
<tr>
<th><a href="./?s=title">Title</a></th>
<th><a href="./?s=artist">Artist</a></th>
<th><a href="./?s=album">Album</a></th>
<th><a href="./?s=year">Year</a></th>
<th><a href="./?s=length">Length</a></th>
</tr>
{{range .T}}
<tr>
<th>{{.Title}}</th>
<th>{{.Artist}}</th>
<th>{{.Album}}</th>
<th>{{.Year}}</th>
<th>{{.Length}}</th>
</tr>
{{end}}
</table>
</body></html>
`

var report = template.Must(template.New("trackList").
	Parse(templ))

var tracksForWeb = customSort{T: tracks}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.URL.Query() {
		if k == "s" {
			// ?s=foo&s=bar とした場合にv[1]にbarが入る
			switch v[0] {
			case "title":
				sortByKey(&tracksForWeb, title)
			case "artist":
				sortByKey(&tracksForWeb, artist)
			case "album":
				sortByKey(&tracksForWeb, album)
			case "year":
				sortByKey(&tracksForWeb, year)
			case "length":
				sortByKey(&tracksForWeb, leng)
			}
		}
	}
	if err := report.Execute(w, tracksForWeb); err != nil {
		log.Fatal(err)
	}
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
