package basename1

func basename(s string) string {
	// 後ろから見てく
	for i := len(s) - 1; i >= 0; i-- {
		// "/"が見つかったらそれ以降を残す
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		// "."が見つかったらそれ以前を残す
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
