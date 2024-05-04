// go with sql
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //為什麼要用底線_來占著位子？
)

var db *sql.DB

// 與DB連線,用 init() 初始化，比 main() 更早執行。
func init() {
	// db類型	 user:password@dbname
	if dbConnect, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/"); err != nil { // sql.Open 只建立物件、容器，並未進行連線
		log.Fatalln(err)
	} else {
		if err := dbConnect.Ping(); err != nil { // Ping() 這裡才開始建立連線。上面 sql.Open 只建立物件、容器，並未進行連線，無法連線並不造成err。
			log.Fatalln(err)
		} else {
			db = dbConnect // 用全域變數接
		}
	}

	db.SetMaxOpenConns(10) // 可設置最大DB連線數，設<=0則無上限（連線分成 in-Use正在執行任務 及 idle執行完成後的閒置 兩種）
	db.SetMaxIdleConns(10) // 設置最大idle閒置連線數。
}

func main() {
	// creatDB("`golang`")
	// deleteDB("`golang`")
	// createTable("`golang`", "`goTable`")
	// deleteTable("`golang`", "`goTable`")
	upadteTable("`golang`", "`goTable`")
	insertTable("peter")
}

// 建立資料庫
func creatDB(dbName string) {
	// 使用db.Exec執行sql語法,Exec只能用來(create、update、delete)
	// 加了IF NOT EXISTS之後，會先檢查要建立的物件存不存在，如果不存在再執行建立。
	if _, err := db.Exec("CREATE DATABASE IF NOT EXISTS" + dbName + ";"); err != nil {
		log.Fatalln(err)
	}
}

// 刪除資料庫
func deleteDB(dbName string) {
	if _, err := db.Exec("DROP DATABASE IF EXISTS" + dbName + ";"); err != nil {
		log.Fatalln(err)
	}
}

// 建立表格
func createTable(dbName, tableName string) {
	// 初始化 Table 沒給任何欄位時，會出現 `A table must have at least 1 column` 的錯誤
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS" + dbName + "." + tableName + "(`name` VARCHAR (10))"); err != nil {
		log.Fatalln(err)
	}

	// 或者用 USE 來指定該資料庫
	// db.Exec("USE `school`")
	// _, err := db.Exec("CREATE TABLE IF NOT EXISTS `student`(`name` VARCHAR (10))")
}

// 刪除表格
func deleteTable(dbName, tableName string) {
	if _, err := db.Exec("DROP TABLE IF EXISTS" + dbName + "." + tableName); err != nil {
		log.Fatalln(err)
	}
}

// 更新表格
func upadteTable(dbName, tableName string) {
	if _, err := db.Exec("ALTER TABLE" + dbName + "." + tableName + "ADD `id` INT AUTO_INCREMENT PRIMARY KEY;"); err != nil {
		log.Fatalln(err)
	}
}

// 插入
func insertTable(valueName string) {
	if rs, err := db.Exec("INSERT INTO `golang`.`goTable`(`name`) VALUES (?)", valueName); err != nil {
		log.Println(err)
	} else {
		rowCount, err := rs.RowsAffected()
		rowId, err := rs.LastInsertId() // 資料表中有Auto_Increment欄位才起作用，回傳剛剛新增的那筆資料ID

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("新增 %d 筆資料，id = %d \n", rowCount, rowId)
	}
}

// 查詢-待完成
// rows, err := db.Query("SELECT `name`, `age` FROM `school`.`teacher`")
// rows, err := db.Query("SELECT * FROM `school`.`teacher`") // 也可以使用`Select *`選取全部欄位。
