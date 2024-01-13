package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miriam-samuels/portfolio-builder/internal/db"
	"github.com/miriam-samuels/portfolio-builder/internal/helper"
	"github.com/miriam-samuels/portfolio-builder/internal/models/user"
)

func GetUserRoutes(w http.ResponseWriter, r *http.Request) {
	var routes []user.UserRoute

	rows, err := db.Portfolio.Query("SELECT username,theme FROM users")
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "Internal Server Error", nil)
		return
	}

	for rows.Next() {
		var route user.UserRoute
		rows.Scan(&route.Username, &route.Theme)
		routes = append(routes, route)
	}

	helper.SendResponse(w, http.StatusOK, true, "Request Successful", routes)
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	// get variable in request url
	vars := mux.Vars(r)
	username := vars["user"]

	// create var to store data
	var user user.UserInfo

	// query db for user info
	row := db.Portfolio.QueryRow("SELECT * FROM users WHERE username=$1", username)

	// variable to store column from db
	var skills string
	var projects string
	var experience string

	// copy column into var
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.Phone, &user.Github, &user.Medium, &user.Twitter, &user.LinkedIn, &user.Objective, &user.Tagline, &user.Theme, &skills, &projects, &experience)
	if err != nil {
		// check if no rows is returned and handle it
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "user does not exist")
			log.Printf("%v", err)
			return
		}
		// check for other possible errors
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// what is coming from the db is json so it unmarshals it types golang understands
	json.Unmarshal([]byte(skills), &user.Skills)
	json.Unmarshal([]byte(projects), &user.Projects)
	json.Unmarshal([]byte(experience), &user.Experience)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// marshall data to be sent back
	json.NewEncoder(w).Encode(user)
}

func SetUserInfo(w http.ResponseWriter, r *http.Request) {
	// get variable in request url
	vars := mux.Vars(r)
	username := vars["user"]

	// get request body
	userInfo := &user.UserInfo{}

	// convert request body to types that golang understands
	err := helper.ParseRequestBody(w, r, userInfo)
	if err != nil {
		helper.SendResponse(w, http.StatusBadRequest, false, "error parsing body:"+err.Error(), nil, err)
		return
	}

	// prepare query statement
	stmt, err := db.Portfolio.Prepare("UPDATE users SET first_name=$1, last_name=$2, email=$3, phone=$4, github=$5, medium=$6, twitter=$7, linkedin=$8, tagline=$9, objective=$10, theme=$11, skills=$12, projects=$13, experience=$14 WHERE username=$15")
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered", nil, err)
		return
	}

	defer stmt.Close()

	// converts struct back to json to be able to store in db
	skills, _ := json.Marshal(userInfo.Skills)
	projects, _ := json.Marshal(userInfo.Projects)
	experience, _ := json.Marshal(userInfo.Experience)

	// execute the statement
	res, err := stmt.Exec(userInfo.FirstName, userInfo.LastName, userInfo.Email, userInfo.Phone, userInfo.Github, userInfo.Medium, userInfo.Twitter, userInfo.LinkedIn, userInfo.Tagline, userInfo.Objective, userInfo.Theme, string(skills), string(projects), username, string(experience))
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered", nil, err)
		return
	}

	// get the rows affected
	rows, _ := res.RowsAffected()

	if rows < 1 {
		helper.SendResponse(w, http.StatusBadRequest, false, "user" + username + "does not exist", nil)
	} else {
		helper.SendResponse(w, http.StatusOK, true, "user updated successfully", nil)
	}
}
