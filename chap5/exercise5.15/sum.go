package sum

func max(vals ...int) int {
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func max2(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}
