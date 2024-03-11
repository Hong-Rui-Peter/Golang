// 程式進入點
package main

import "fmt"

func main() {
	fuck_golang()
	input_output()
	nine_to_nine()
	even_odd_num()
	sort()

	//費式數列
	var n int
	fmt.Print("請輸入費式數列數字:")
	fmt.Scanln(&n)
	fmt.Println(fibonacci1(n))
	fmt.Println(fibonacci2(n))

	triangle1()
	triangle2()
}
