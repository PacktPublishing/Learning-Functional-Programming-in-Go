package fib

func Memoize(mf Memoized) Memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := mf(key)
		cache[key] = temp
		return temp
	}
}

type Memoized func(int) int
var fibMem Memoized
func FibMemoized(n int) int {
	n += 1
	fibMem = Memoize(func(n int) int {
		if n == 0 || n == 1 {
			return n
		}
		return fibMem(n - 2) + fibMem(n - 1)
	})
	return fibMem(n)
}
