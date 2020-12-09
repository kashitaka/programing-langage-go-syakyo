package exercise23

// 全て0の256要素のスライス
var pc [256]byte

// init関数はプログラム開始時に呼ばれる
// 他から呼べない
func init() {
	// [0 1 1 2 1 2 2 3 1 2 2 3 2 3 3 ..... 8]
	// のようにある整数がbitで何個1をもつかを表す
	// p[256] = 8 になる。255はバイナリで1が8個
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}
