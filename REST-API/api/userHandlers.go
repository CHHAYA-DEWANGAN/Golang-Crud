package api

import (
	"REST-API/REST-API/common"
	"REST-API/REST-API/crud"
	"REST-API/REST-API/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
	excelize "github.com/xuri/excelize/v2"
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
func ExportListOfUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userdata, err := crud.ListOfUser(db)
		f := excelize.NewFile()

		var cell string

		for index, user := range userdata {
			typeOfField := reflect.TypeOf(user)
			valueOfField := reflect.ValueOf(user)

			for i := 0; i < typeOfField.NumField(); i++ {
				field := typeOfField.Field(i)
				value := valueOfField.Field(i).Interface()
				fmt.Printf("%s: %v\n", field.Name, value)

				cell = common.ToGetAlphaString(i+1) + fmt.Sprintf("%d", index)
				f.SetCellValue("Sheet1", cell, value)

			}
			// Set value of a cell
			cell = common.ToGetAlphaString(index) + fmt.Sprintf("%d", index)

		}

		// Save the Excel file
		randomString := common.GenerateRandomString(5)
		filepath := "example_" + randomString + ".xlsx"
		if err := f.SaveAs(filepath); err != nil {
			fmt.Println("Error saving Excel file:", err)
			return
		}
		currentDir := common.GetWorkingDirectory()
		fmt.Println("Excel file created successfully.")

		if err != nil {
			http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(currentDir + "/" + filepath)
	}
}
