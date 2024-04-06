// 登入網頁
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/signin", signIn)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func signIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { // 判斷http動作是否為GET
		if tmpl, err := template.ParseFiles("sign_in.html"); err != nil { // sign_in.html檔案內容
			log.Fatal(err)
			return
		} else {
			err = tmpl.Execute(w, "") // http執行
		}
	} else { // GET以外的行為
		r.ParseForm() // 解析POST

		var username, password string

		for k, v := range r.Form { // 取key,value值
			// if k == "username" {
			// 	username = v[0]
			// } else {
			// 	password = v[0]
			// }

			switch k {
			case "username":
				username = v[0]
			case "password":
				password = v[0]
			default:
			}
		}

		if username == "peter" && password == "123456" {
			fmt.Fprintln(w, "Login Success")
		} else {
			fmt.Fprintln(w, "Login Fail")
		}
	}
}
