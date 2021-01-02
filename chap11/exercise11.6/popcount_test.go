package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	var v uint64 = 1234567890
	for i := 0; i < b.N; i++ {
		PopCount(v)
	}
}

func BenchmarkPopCount24(b *testing.B) {
	var v uint64 = 1234567890
	for i := 0; i < b.N; i++ {
		PopCount24(v)
	}
}

func BenchmarkPopCount25(b *testing.B) {
	var v uint64 = 1234567890
	for i := 0; i < b.N; i++ {
		PopCount25(v)
	}
}
