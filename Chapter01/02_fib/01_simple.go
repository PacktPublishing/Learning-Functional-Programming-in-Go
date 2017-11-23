package fib

func FibSimple(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return FibSimple(n-2) + FibSimple(n-1)
	}
}
