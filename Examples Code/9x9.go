// 九九成法表
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i, "*", j, "=", i*j, " ")
		}
		fmt.Println()
	}
}
