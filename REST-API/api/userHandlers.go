package api

import (
	"REST-API/REST-API/crud"
	"REST-API/REST-API/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func PrintName(w http.ResponseWriter, r *http.Request) {
	// Extract the 'name' parameter from the URL query parameters
	name := r.URL.Query().Get("name")

	// If 'name' parameter is not provided, default to "chhaya"
	if name == "" {
		name = "chhaya"
	}

	crud.PrintName(name)

	// Write the response
	fmt.Fprintf(w, "Hello, %s!", name)
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Println("Invalid Request")
			fmt.Println(err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		fmt.Printf("Model Value %#v", user)
		fmt.Println(user.Address)
		result, err := crud.CreateUser(db, user)
		if err != nil {
			fmt.Println(err)
		}

		if result {
			fmt.Println("User Created Successfully")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.Result{Status: "success", Message: "User Created Successfully"})
	}
}

func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := models.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Println("Invalid Request")
			fmt.Println(err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			http.Error(w, "User ID not provided", http.StatusBadRequest)
			return
		}
		updatedId, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		fmt.Printf("Person struct: %#v\n", user)
		fmt.Println(user.Address)
		result, err := crud.UpdateUser(db, user, updatedId)
		if err != nil {
			fmt.Println(err)
		}

		if result {
			fmt.Println("User Update Successfully")
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.Result{Status: "success", Message: "User Updated Successfully"})

	}
}

func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		fmt.Println(idStr)
		if !ok {
			http.Error(w, "User ID not provided", http.StatusBadRequest)
			return
		}
		deleteID, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		result, err := crud.DeleteUser(db, deleteID)
		if err != nil {
			fmt.Println(err)
		}

		if result {
			fmt.Println("User Deleted Successfully")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.Result{Status: "success", Message: "User Deletd Successfully"})
	}
}
func ListOfUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userdata, err := crud.ListOfUser(db)

		if err != nil {
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userdata)
	}
}
