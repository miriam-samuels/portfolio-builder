package authControllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/miriam-samuels/src/config"
	conn "github.com/miriam-samuels/src/database"
	authModels "github.com/miriam-samuels/src/models/auth"
	userModels "github.com/miriam-samuels/src/models/user"
	"github.com/miriam-samuels/src/utils"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var cred authModels.SignUpCredentials
	var responseData authModels.Response

	// decode request body and store in cred variable
	_ = json.NewDecoder(r.Body).Decode(&cred)

	// validate body to ensure all properties are available
	err := cred.ValidateSignUp()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	// encrypt password
	encryptedPass, err := utils.Encrypt(cred.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	// generate user Id
	userId := config.GenerateId()
	// prepare query statment
	stmt, err := conn.Db.Prepare("INSERT INTO users SET id=?,username=?,password=?,email=?")
	if err != nil {
		log.Printf("%v where it was affected", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not insert into db")
		return
	}

	// execute query statement
	result, err := stmt.Exec(userId, cred.Username, encryptedPass, cred.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Username or email already exist")
		log.Printf("%v", err)
		return
	}

	userToken, err := authModels.GenerateToken(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to register user")
		log.Printf("%v", err)
		return
	}

	responseData = authModels.Response{
		Status: true,
		Data: map[string]interface{}{
			"token": userToken,
		},
		Message: "Signup successful"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
	log.Printf("%v", result)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var cred authModels.LoginCredentials
	var responseData authModels.Response

	// decode request body and store in cred variable
	_ = json.NewDecoder(r.Body).Decode(&cred)

	// validate body to ensure all properties are available
	err := cred.ValidateLogin()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	// prepare query statment for row matching username
	row := conn.Db.QueryRow("SELECT id,username,password FROM users WHERE username=?", cred.Username)

	// variable to store data from db
	var storedCred userModels.UserInfo

	// copy columnn from matched row into variable pointed at
	err = row.Scan(&storedCred.Id, &storedCred.Username, &storedCred.Password)
	if err != nil {
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

	// compare entered password with password in db
	err = bcrypt.CompareHashAndPassword([]byte(storedCred.Password), []byte(cred.Password))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%v", err)
		fmt.Fprintf(w, "Incorrect password")
	}

	// generate token
	userToken, err := authModels.GenerateToken(storedCred.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to register user")
		log.Printf("%v", err)
		return
	}

	responseData = authModels.Response{
		Status: true,
		Data: map[string]interface{}{
			"token": userToken,
		},
		Message: "Login successful"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}
