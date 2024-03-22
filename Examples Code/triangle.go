// 打印三角形
package main

import "fmt"

var n int

func main() {
	fmt.Print("請輸三角形高度:")
	fmt.Scanln(&n)

	triangle1()
	triangle2()
	triangleTeacher()
}

//實心三角形
func triangle1() {
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

//空心三角形(教師版)
func triangleTeacher() {
	for r := 0; r < n; r++ {
		if r != n-1 { // 判斷是否為最後一行
			// 打印空白
			for c := n; c > r+1; c-- {
				fmt.Print(" ")
			}

			fmt.Print("* ") // 打印*(頭)

			//打印中間空白
			for c := 0; c < (r - 1); c++ {
				fmt.Print("  ")
			}

			// 打印*(尾)
			if r != 0 {
				fmt.Print("* ")
			}

			fmt.Println()
		} else {
			// 最後一行不需要打印空白
			for m := 0; m < n; m++ {
				fmt.Print("* ")
			}
		}
	}
}
