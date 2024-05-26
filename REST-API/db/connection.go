package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnection() *sql.DB {
	fmt.Print("heretyui")

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/chhayaGolang")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error While connecting DB")
		panic("Error While connecting DB")
	}
	fmt.Println("Connected successfully")
	fmt.Println(db)
	// defer db.Close()
	return db
}
