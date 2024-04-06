// BMI網頁版
package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/BMI", BMI)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func BMI(w http.ResponseWriter, r *http.Request) {
	weight := r.URL.Query().Get("weight")
	height := r.URL.Query().Get("height")
	wg, _ := strconv.ParseFloat(weight, 64)
	hg, _ := strconv.ParseFloat(height, 64)
	bmi := wg / math.Pow(hg/100, 2)
	bmi = math.Round(bmi*100) / 100

	if weight == "" || height == "" {
		fmt.Fprintln(w, "請輸入身高體重!!")
		return
	}

	if tmpl, err := template.ParseFiles("BMI.html"); err != nil {
		log.Fatal(err)
	} else {
		tmpl.Execute(w, struct {
			Weight string
			Height string
			BMI    float64
		}{
			weight,
			height,
			bmi,
		})
	}
}
