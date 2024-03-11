// 費式數列(遞迴)
package main

func fibonacci1(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci1(n-2) + fibonacci1(n-1)
}
