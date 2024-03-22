// 判斷奇偶數
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int

	fmt.Print("請輸入一個數字:")
	fmt.Scan(&number)

	if number%2 == 0 { //除2餘數為0
		fmt.Println(strconv.Itoa(number) + "是偶數\n")
	} else {
		fmt.Println(strconv.Itoa(number) + "是奇數\n")
	}
}
