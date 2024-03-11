// 費式數列(徒法煉鋼)
package main

func fibonacci2(n int) int {
	if n < 2 {
		return n
	}

	a, b := 0, 1

	for i := 0; i < n; i++ {
		next := a + b
		a, b = b, next
	}

	return a
}
