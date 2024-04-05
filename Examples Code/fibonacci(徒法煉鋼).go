// 費式數列(徒法煉鋼)
package main

import "fmt"

func main() {
	var n int

	fmt.Print("請輸入費式數列數字:")
	fmt.Scanln(&n)

	if n == 0 {
		fmt.Println(n)
		return
	} else if n == 1 {
		fmt.Printf("%d , %d", 0, n)
		return
	}

	a, b := 0, 1
	fmt.Print(0, ",")

	for i := 0; i < n; i++ {
		next := a + b
		a, b = b, next
		fmt.Print(a, ",")
	}

}

// 0 1 1 2 3 5 8 13
