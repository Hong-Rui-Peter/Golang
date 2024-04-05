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
	if r.Method == "GET" {
		if tmpl, err := template.ParseFiles("sign_in_sign_out.html"); err != nil {
			log.Fatal(err)
			return
		} else {
			err = tmpl.Execute(w, "")
		}
	} else {
		r.ParseForm()

		var username, password string

		for k, v := range r.Form {
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
