// 乘法表網頁版
package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/9x9", multiplicationTable)

	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal(err)
	}
}

func multiplicationTable(w http.ResponseWriter, r *http.Request) {
	var row, col int
	var err1, err2 error
	// 讀取url帶入的參數(parameter)
	inputRow := r.URL.Query().Get("row") // 參數為row
	inputCol := r.URL.Query().Get("col") // 參數為col

	// 根據帶入的parameter產生幾欄幾列的乘法表，未帶parameter預設9x9
	if inputRow != "" && inputCol != "" {
		row, err1 = strconv.Atoi(inputRow)
		col, err2 = strconv.Atoi(inputCol)
	} else {
		row, col = 9, 9
	}

	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
		return
	}

	//生成表格
	table := make([][]string, row)

	for i := range table {
		table[i] = make([]string, col)
	}

	// table內容
	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			table[i-1][j-1] = strconv.Itoa(i) + "x" + strconv.Itoa(j) + "=" + strconv.Itoa(i*j)
		}
	}

	if tmpl, err := template.ParseFiles("multiplication_table.html"); err != nil { // 讀取multiplication_table.html檔案內容
		log.Fatal(err)
		return
	} else {
		if err = tmpl.Execute(w, struct{ Table [][]string }{table}); err != nil { // 執行
			log.Fatal(err)
			return
		}
	}
}
