package exercise25

func PopCount(x uint64) int {
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
