package charcount

import (
	"reflect"
	"strings"
	"testing"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input  string
		counts map[rune]int
		utflen []int
	}{
		{input: "aaaabbbcc\nccあ",
			counts: map[rune]int{'a': 4, 'b': 3, 'c': 4, '\n': 1, 'あ': 1},
			utflen: []int{0, 12, 0, 1, 0},
		},
	}
	for _, test := range tests {
		counts, utflen := CharCount(strings.NewReader(test.input))
		if !reflect.DeepEqual(test.counts, counts) {
			t.Errorf("CharCount(%s) got %v, want %v", test.input, counts, test.counts)
		}
		if !reflect.DeepEqual(test.utflen, utflen) {
			t.Errorf("CharCount(%s) got %v, want %v", test.input, utflen, test.utflen)
		}
	}
}
