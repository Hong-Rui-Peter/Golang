// 打印三角形
package main

import "fmt"

//實心三角形
func triangle1() {
	var n int

	fmt.Print("請輸三角形高度:")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		//打印空格
		for h := n; h >= i; h-- {
			fmt.Print("  ")
		}

		//	打印三角形
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

//空心三角形
func triangle2() {
	var n int

	fmt.Print("請輸三角形高度:")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		//打印空格
		for h := n; h >= i; h-- {
			fmt.Print("  ")
		}

		//	打印三角形
		for j := 1; j <= 2*i-1; j++ {
			if i == n {
				fmt.Print("* ")
				continue
			}

			if j > 1 && j < 2*i-1 {
				fmt.Print("  ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}
