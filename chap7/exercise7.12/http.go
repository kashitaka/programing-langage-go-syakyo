package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

const templ = `
<html><body>
<table border="1">
<tr>
<th>Item</th>
<th>Price</th>
</tr>
{{range $key, $value := .}}
<tr>
<th>{{$key}}</th>
<th>{{$value}}</th>
</tr>
{{end}}
</table>
</body></html>
`

var report = template.Must(template.New("db").
	Parse(templ))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := report.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item required")
		return
	}
	price, err := strconv.Atoi(req.URL.Query().Get("price"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid price: %q\n", err)
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "OK\n")
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
