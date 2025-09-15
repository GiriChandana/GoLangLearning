package main

import (
	//"database/sql"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	id   int
	name string
}

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
func main() {
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DbUser, DbPassword, DBName)
	db, err := sql.Open("mysql", connectionString)
	checkError(err)

	result, err := db.Exec("insert into data values(4, 'xyz')")
	checkError(err)
	lastInsertId, err := result.LastInsertId()
	fmt.Println("lastInsertId: ", lastInsertId)
	checkError(err)
	rowsAffected, err := result.RowsAffected()
	fmt.Println("rowsAffected: ", rowsAffected)
	checkError(err)
	defer db.Close()

	rows, err := db.Query("SELECT * from data")
	checkError(err)

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		checkError(err)
		fmt.Println(data)
	}
}
