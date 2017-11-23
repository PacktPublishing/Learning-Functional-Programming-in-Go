package fib

import "testing"

func TestChanneled(t *testing.T) {
	for _, ft := range fibTests {
		if v := FibChanneled(ft.a); v != ft.expected {
			t.Errorf("FibChanneled(%d) returned %d, expected %d", ft.a, v, ft.expected)
		}
	}
}

func BenchmarkFibChanneled(b *testing.B) {
	fn := FibChanneled
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}

func benchmarkFibChanneled(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibChanneled(i)
	}
}

func BenchmarkFibChanneled1(b *testing.B)  { benchmarkFibChanneled(1, b) }
func BenchmarkFibChanneled2(b *testing.B)  { benchmarkFibChanneled(2, b) }
func BenchmarkFibChanneled3(b *testing.B)  { benchmarkFibChanneled(3, b) }
func BenchmarkFibChanneled10(b *testing.B) { benchmarkFibChanneled(4, b) }
func BenchmarkFibChanneled20(b *testing.B) { benchmarkFibChanneled(20, b) }
func BenchmarkFibChanneled40(b *testing.B) { benchmarkFibChanneled(42, b) }