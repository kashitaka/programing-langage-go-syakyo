package popcount

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
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount24(x uint64) int {
	sum := 0
	// 0000.......0001 とのbit積を取る
	mask := uint64(1)
	for i := 0; i < 64; i++ {
		if (x>>i)&mask == 1 {
			sum += 1
		}
	}
	return sum
}

func PopCount25(x uint64) int {
	sum := 0
	// 「1が設定されている最下位ビットをクリア」
	// 例えば x = 1010101 なら
	// 1010100 -> 1010000 -> 1000000 -> 0000000
	// という風に1回ごとに1番後ろの1が0に反転する
	for x > 0 {
		x = x&x - 1
		sum += 1
	}
	return sum
}
