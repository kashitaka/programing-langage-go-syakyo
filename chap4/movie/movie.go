package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	// フィールドタグ。json: で始まるものは"encoding/json"で使われる
	Year int `json:"released"`
	// omitemptyはゼロ値のときにjsonにkeyを出さない
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogard", "Ingrid Bergman"}},
	}
	// マーシャリング: 構造体の値をjson文字列に変換
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// 読みやすい形にindentをつけて[]byteにしてくれる
	data, err = json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// アンマーシャリング: json文字列から構造体の値を作る
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)

	var mvs []Movie
	if err := json.Unmarshal(data, &mvs); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%#v\n", mvs)

}
