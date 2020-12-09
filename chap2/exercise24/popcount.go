package exercise24

func PopCount(x uint64) int {
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
