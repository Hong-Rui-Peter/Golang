// 正規表達式
package main

import (
	"fmt"
	"regexp"
)

func main() {
	checkUID()
	checkPhoneNumber()
	checkEmail()
}

// 檢查身分證
func checkUID() {
	var uid string

	for {
		fmt.Print("請輸入身分證:")
		fmt.Scanln(&uid)

		if regexp.MustCompile(`^[A-Za-z][1-2]\d{8}$`).MatchString(uid) { // MustCompile:檢查的規則,MatchString:被檢查的參數
			fmt.Println(uid)
			break
		}
	}
}

// 檢查電話號碼
func checkPhoneNumber() {
	var pn string

	for {
		fmt.Print("請輸入電話號碼:")
		fmt.Scanln(&pn)

		if regexp.MustCompile(`^09\d{8}`).MatchString(pn) { // MustCompile:檢查的規則,MatchString:被檢查的參數
			fmt.Println(pn)
			break
		}
	}
}

// 檢查Email
func checkEmail() {
	var email string

	for {
		fmt.Print("請輸入Email:")
		fmt.Scanln(&email)

		if regexp.MustCompile(`^\w+@\w+\.com$`).MatchString(email) { // MustCompile:檢查的規則,MatchString:被檢查的參數
			fmt.Println(email)
			break
		}
	}
}
