package charcount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func CharCount(rd io.Reader) (map[rune]int, [5]int) {
	counts := make(map[rune]int)
	// utflen: 何バイト長の文字がそれぞれ何個あるかを記録するmap
	// utf8.UTFMax = 4
	// utflen[0] は 使わず +1 して [1]から[4]を使う
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(rd)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			// 65533
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	return counts, utflen
}
