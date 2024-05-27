package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHandler(db *sql.DB) {
	http.HandleFunc("/printname", PrintName)
	http.HandleFunc("/user/create", CreateUser(db))
	// http.HandleFunc("/user/update/{id}", UpdateUser(db))
	// http.HandleFunc("/user/delete/{id}", DeleteUser(db))
	http.HandleFunc("/user/list", ListOfUser(db))

	http.HandleFunc("/callexternalapi", CallExternalApi)

	router := mux.NewRouter()

	router.HandleFunc("/user/update/{id}", UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/user/delete/{id}", DeleteUser(db)).Methods("DELETE")
	http.Handle("/", router)

	http.ListenAndServe(":8086", nil)
}
