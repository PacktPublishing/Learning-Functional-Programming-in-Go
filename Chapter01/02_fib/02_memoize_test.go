package fib

import "testing"

func TestMemoized(t *testing.T) {
	for _, ft := range fibTests {
		if v := FibMemoized(ft.a); v != ft.expected {
			t.Errorf("FibMemoized(%d) returned %d, expected %d", ft.a, v, ft.expected)
		}
	}
}

func BenchmarkFibMemoized(b *testing.B) {
	fn := FibMemoized
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}

func benchmarkFibMemoized(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibMemoized(i)
	}
}

func BenchmarkFibMemoized1(b *testing.B)  { benchmarkFibMemoized(1, b) }
func BenchmarkFibMemoized2(b *testing.B)  { benchmarkFibMemoized(2, b) }
func BenchmarkFibMemoized3(b *testing.B)  { benchmarkFibMemoized(3, b) }
func BenchmarkFibMemoized10(b *testing.B) { benchmarkFibMemoized(4, b) }
func BenchmarkFibMemoized20(b *testing.B) { benchmarkFibMemoized(20, b) }
func BenchmarkFibMemoized40(b *testing.B) { benchmarkFibMemoized(42, b) }