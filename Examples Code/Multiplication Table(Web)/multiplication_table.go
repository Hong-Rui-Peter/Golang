package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/9x9", multiplicationTable)
	http.ListenAndServe(":9090", nil)
}

func multiplicationTable(w http.ResponseWriter, r *http.Request) {
	var row, col int
	var err1, err2 error
	inputRow := r.URL.Query().Get("row")
	inputCol := r.URL.Query().Get("col")

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

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			table[i-1][j-1] = strconv.Itoa(i) + "x" + strconv.Itoa(j) + "=" + strconv.Itoa(i*j)
		}
	}

	if tmpl, err := template.ParseFiles("multiplication_table.html"); err != nil {
		log.Fatal(err)
		return
	} else {
		if err = tmpl.Execute(w, struct{ Table [][]string }{table}); err != nil {
			log.Fatal(err)
			return
		}
	}
}
