package router

import (
	"encoding/json"
	"net/http"

	"github.com/arganaphangquestian/gobasic/model"
	"github.com/arganaphangquestian/gobasic/repository"
)

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User

	db, err := repository.OpenDB()
	defer db.Close()

	result, err := db.Query("SELECT * FROM users")
	for result.Next() {
		var user model.User
		_ = result.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Token, &user.RoleID)
		users = append(users, user)
	}
	if err != nil {
		json.NewEncoder(w).Encode(model.ResponseAPI{
			Status:  "Failed",
			Message: "Failed Get All User",
			Data:    nil,
		})
		return
	}
	json.NewEncoder(w).Encode(model.ResponseAPI{
		Status:  "Success",
		Message: "Get All User",
		Data:    users,
	})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenDB()
	defer db.Close()
	var user model.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(model.ResponseAPI{
			Status:  "Failed",
			Message: "Failed to decode data",
			Data:    r.Body,
		})
		return
	}
	stmt := "INSERT INTO users(username, email, password) VALUES($1, $2, $3) "
	user.Token = nil
	user.RoleID = "3"

	_, err = db.Exec(stmt, user.Email, user.Username, user.Password)
	if err != nil {
		json.NewEncoder(w).Encode(model.ResponseAPI{
			Status:  "Failed",
			Message: "Failed to insert data",
			Data:    err,
		})
		return
	}
	json.NewEncoder(w).Encode(model.ResponseAPI{
		Status:  "Success",
		Message: "Register Successfully",
		Data:    user,
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	db, err := repository.OpenDB()
	defer db.Close()

	var userData model.User
	err = json.NewDecoder(r.Body).Decode(&userData)

	if err != nil {
		json.NewEncoder(w).Encode(
			model.ResponseAPI{
				Status:  "Failed",
				Message: "Failed to Decode data you sent",
				Data:    err,
			},
		)
		return
	}
	stmt := "SELECT * FROM users WHERE username = $1 AND password = $2"

	result := db.QueryRow(stmt, userData.Username, userData.Password)
	err = result.Scan(&userData.ID, &userData.Email, &userData.Username, &userData.Password, &userData.Token, &userData.RoleID)

	if err != nil {
		json.NewEncoder(w).Encode(
			model.ResponseAPI{
				Status:  "Failed",
				Message: "Failed to Login",
				Data:    err,
			},
		)
		return
	}

	json.NewEncoder(w).Encode(model.ResponseAPI{
		Status:  "Success",
		Message: "Register Successfully",
		Data:    userData,
	},
	)
}
