package main

import (
	"REST-API/REST-API/api"
	"REST-API/REST-API/db"
)

func main() {
	db := db.DbConnection()
	api.RegisterHandler(db)
}
