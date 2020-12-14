package main

func main() {
	l := []int{1, 2, 3, 4, 5, 6}
	reverse(l)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
