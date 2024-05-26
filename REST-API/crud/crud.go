package crud

import (
	"REST-API/REST-API/models"
	"database/sql"
	"fmt"
	"log"
)

func CreateUser(dbConn *sql.DB, user models.User) (bool, error) {

	query := "INSERT INTO users (name, email, phone,age,address) VALUES (?, ?, ?, ?, ?)"
	result, err := dbConn.Exec(query, user.Name, user.Email, user.Phone, user.Age, user.Address)

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	fmt.Printf("Inserted user with ID: %d\n", lastInsertID)
	return true, nil

}

func UpdateUser(dbConn *sql.DB, user models.User, updateId int) (bool, error) {
	query := "UPDATE users SET name = ?, email = ?, phone =? ,age=?,address =? WHERE id = ?"
	_, err := dbConn.Exec(query, user.Name, user.Email, user.Phone, user.Age, user.Address, updateId)

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	fmt.Println("Updated Successfully")
	return true, nil
}

func DeleteUser(dbConn *sql.DB, deletedId int) (bool, error) {
	query := "DELETE FROM users WHERE id = ?"
	result, err := dbConn.Exec(query, deletedId)
	if err != nil {
		fmt.Println("Error While deleting")
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Failed to delete user")
		return false, err
	}
	if rowsAffected == 0 {
		fmt.Println("User not found")

		return false, nil
	}

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	fmt.Println("Updated Successfully")
	return true, nil
}

func ListOfUser(dbConn *sql.DB) ([]models.User, error) {
	users := make([]models.User, 0)

	rows, err := dbConn.Query("SELECT name, email, age, address,phone FROM users")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Name, &user.Email, &user.Age, &user.Address, &user.Phone)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}
