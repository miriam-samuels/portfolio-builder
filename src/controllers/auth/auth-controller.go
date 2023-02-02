package authControllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/miriam-samuels/src/database"
	authModels "github.com/miriam-samuels/src/models/auth"
	"github.com/miriam-samuels/src/utils"
	"golang.org/x/crypto/bcrypt"
)

var db = database.PortfolioDb

func SignUp(w http.ResponseWriter, r *http.Request) {
	var cred authModels.SignUpCredentials
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

	// prepare query statment
	stmt, err := db.Prepare("INSERT INTO users SET username=?,password=?,email=?")
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// execute query statement
	result, err := stmt.Exec(cred.Username, encryptedPass, cred.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Username or email already exist")
		log.Printf("%v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Signup successful")
	log.Printf("%v", result)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var cred authModels.LoginCredentials
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
	row := db.QueryRow("SELECT username,password FROM users WHERE username=?", cred.Username)

	// variable to store data from db
	storedCred := authModels.LoginCredentials{}

	// copy columnn from matched row into variable pointed at
	err = row.Scan(&storedCred.Username, &storedCred.Password)
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

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Login Successful")
}
