package userController

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	conn "github.com/miriam-samuels/src/database"
	userModels "github.com/miriam-samuels/src/models/user"
	authModels "github.com/miriam-samuels/src/models/auth"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	// get variable in request url
	vars := mux.Vars(r)
	username := vars["user"]

	// create var to store data
	var user userModels.UserInfo

	// query db for user info
	row := conn.Db.QueryRow("SELECT * FROM users WHERE username=$1", username)

	// variable to store column from db
	var skills string
	var projects string

	// copy column into var
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.Phone, &user.Github, &user.Medium, &user.Twitter, &user.LinkedIn, &user.Objective, &user.Tagline, &user.Theme, &skills, &projects)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// marshall data to be sent back
	json.NewEncoder(w).Encode(user)
}

func SetUserInfo(w http.ResponseWriter, r *http.Request) {
	var responseData authModels.Response

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
	stmt, err := conn.Db.Prepare("UPDATE users SET first_name=$1, last_name=$2, email=$3, phone=$4, github=$5, medium=$6, twitter=$7, linkedin=$8, tagline=$9, objective=$10, theme=$11, skills=$12, projects=$13 WHERE username=$14")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occured when preparing statement:: %v", err)
		return
	}

	defer stmt.Close()

	// converts struct back to json to be able to store in db
	skills, _ := json.Marshal(userInfo.Skills)
	projects, _ := json.Marshal(userInfo.Projects)

	// execute the statement
	res, err := stmt.Exec(userInfo.FirstName, userInfo.LastName, userInfo.Email, userInfo.Phone, userInfo.Github, userInfo.Medium, userInfo.Twitter, userInfo.LinkedIn, userInfo.Tagline, userInfo.Objective, userInfo.Theme, string(skills), string(projects), username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error occured when executing statement:: %v", err)
		return
	}

	// get the rows affected
	rows, _ := res.RowsAffected()
	fmt.Printf("record successfully set :: %d rows affected", rows)
	w.Header().Set("Content-Type", "application/json")
	if rows < 1 {
		w.WriteHeader(http.StatusBadRequest)
		responseData = authModels.Response{
			Status:  false,
			Data:    map[string]interface{}{},
			Message: "User does not exist",
		}
	} else {
		w.WriteHeader(http.StatusOK)
		responseData = authModels.Response{
			Status:  true,
			Data:    map[string]interface{}{},
			Message: "user updated successfully",
		}
	}
	
	json.NewEncoder(w).Encode(responseData)

}
