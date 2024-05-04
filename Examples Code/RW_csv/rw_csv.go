// 讀寫檔案
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var Pwd string
var FilePath string
var FileName = "rw_csv.csv"

func init() {
	Pwd, _ = os.Getwd()
	FilePath = filepath.Join(Pwd, FileName)
}

func main() {
	file, err := os.OpenFile(FilePath, os.O_RDONLY, 0777) // os.O_RDONLY 表示只讀、0777 表示(owner/group/other)權限

	if err != nil {
		log.Fatalln("找不到CSV檔案路徑:", FilePath, err)
	}

	// read
	r := csv.NewReader(file)
	r.Comma = ',' // 以何種字元作分隔，預設為`,`。所以這裡可拿掉這行

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		var a = record[0]
		var b = record[1]

		fmt.Println(a, b)
	}

	// Write()
}

func Write() {
	file, err := os.OpenFile(FilePath, os.O_WRONLY, 0777)

	if err != nil {
		log.Fatalln("找不到CSV檔案路徑:", FilePath, err)
	}

	// 寫入單行
	w := csv.NewWriter(file)
	x := []string{"999"}
	w.Write(x)
	w.Flush() // 把在buffer緩存中的所有資料輸出

	// 寫入多行
	ws := csv.NewWriter(file)
	xs := [][]string{{"999", "1"}, {"998", "997"}}
	ws.WriteAll(xs)
	ws.Flush() // 把在buffer緩存中的所有資料輸出
}
