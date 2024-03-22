// 氣泡排序法
package main

import "fmt"

func main() {
	n := make([]int, 5) // slice(切片)初始化

	fmt.Print("請輸入5個數字:")

	// 一個一個輸入，包成array
	for i := 0; i < 5; i++ {
		fmt.Scan(&n[i])
	}

	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n)-1; j++ {
			if n[j] > n[j+1] { // 如果前者大於後者
				// 前後交換(三角交換)
				temp := n[j+1]
				n[j+1] = n[j]
				n[j] = temp
			}
		}
	}

	fmt.Println(n)
}
