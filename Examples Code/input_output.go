// 輸入輸出
package main

import "fmt"

func main() {
	fmt.Println("fuck golang")

	var name string

	fmt.Print("請輸入名字:") // 輸出
	fmt.Scan(&name)     // 輸入
	fmt.Println(name + "好帥")
}
