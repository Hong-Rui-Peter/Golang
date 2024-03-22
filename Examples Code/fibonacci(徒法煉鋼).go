// 費式數列(徒法煉鋼)
package main

import "fmt"

func main() {
	var n int

	fmt.Print("請輸入費式數列數字:")
	fmt.Scanln(&n)

	if n < 2 {
		fmt.Println(n)
		return
	}

	a, b := 0, 1

	for i := 0; i < n; i++ {
		next := a + b
		a, b = b, next
	}

	fmt.Println(a)
}

// 0 1 1 2 3 5 8 13
