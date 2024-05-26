package api

import (
	"database/sql"
	"net/http"
)

func RegisterHandler(db *sql.DB) {
	http.HandleFunc("/printname", PrintName)
	http.HandleFunc("/user/create", CreateUser(db))
	http.HandleFunc("/user/update/{id}", UpdateUser(db))
	http.HandleFunc("/user/delete/{id}", DeleteUser(db))
	http.HandleFunc("/user/list", ListOfUser(db))

	http.ListenAndServe(":8086", nil)
}