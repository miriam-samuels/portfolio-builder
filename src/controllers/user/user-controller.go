package userController

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miriam-samuels/src/database"
	userModels "github.com/miriam-samuels/src/models/user"
)

var db = database.PortfolioDb

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	// get variable in request url
	vars := mux.Vars(r)
	username := vars["user"]

	// create var to store data
	var user userModels.UserInfo

	// query db for user info
	row := db.QueryRow("SELECT * FROM users WHERE username=?", username)

	// variable to store column from db
	var skills string
	var projects string

	// copy column into var
	err := row.Scan(&user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Phone, &user.Github, &user.Medium, &user.Twitter, &user.LinkedIn, &user.Objective, &user.Tagline, &skills, &projects, &user.Theme)
	if err != nil {
		// check if no rows is returned and handle it
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "user does not exist")
			log.Printf("%v", err)
			return
		}
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// what is coming from the db is json so it unmarshals it types golang understands
	json.Unmarshal([]byte(skills), &user.Skills)
	json.Unmarshal([]byte(projects), &user.Projects)

	// marshall data to be sent back
	json.NewEncoder(w).Encode(user)
	w.WriteHeader(http.StatusOK)
}

func SetUserInfo(w http.ResponseWriter, r *http.Request) {
	// get variable in request url
	vars := mux.Vars(r)
	username := vars["user"]

	// get request body
	var userInfo userModels.UserInfo

	// convert request body to types that golang understands
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse body : %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// prepare query statement
	stmt, err := db.Prepare("UPDATE users SET first_name=?, last_name=?, email=?, phone=?, github=?, medium=?, twitter=?, linkedin=?, tagline=?, objective=?, skills=?, projects=?, theme=? WHERE username=?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occured when preparing statement:: %v", err)
		return
	}

	// converts struct back to json to be able to store in db
	skills, _ := json.Marshal(userInfo.Skills)
	projects, _ := json.Marshal(userInfo.Projects)

	// execute the statement
	res, err := stmt.Exec(userInfo.FirstName, userInfo.LastName, userInfo.Email, userInfo.Phone, userInfo.Github, userInfo.Medium, userInfo.Twitter, userInfo.LinkedIn, userInfo.Tagline, userInfo.Objective, skills, projects, userInfo.Theme, username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occured when executing statement:: %v", err)
		return
	}

	// get the rows affected
	rows, _ := res.RowsAffected()
	fmt.Fprintf(w, "record successfully set :: %d rows affected", rows)
}
