// 輸入輸出
package main

import "fmt"

func fuck_golang() {
	fmt.Println("fuck golang")
}

func input_output() {
	var name string

	fmt.Print("請輸入名字:") // 輸出
	fmt.Scan(&name)     // 輸入
	fmt.Println(name + "好帥")
}
