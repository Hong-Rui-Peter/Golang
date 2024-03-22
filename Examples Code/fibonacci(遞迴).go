// 費式數列(遞迴)
package main

import "fmt"

func main() {
	var n int

	fmt.Print("請輸入費式數列數字:")
	fmt.Scanln(&n)
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

// 0 1 1 2 3 5 8 13
