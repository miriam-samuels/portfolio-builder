package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/miriam-samuels/portfolio-builder/internal/db"
	"github.com/miriam-samuels/portfolio-builder/internal/helper"
	"github.com/miriam-samuels/portfolio-builder/internal/models/auth"
	"github.com/miriam-samuels/portfolio-builder/internal/models/user"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var cred auth.SignUpCredentials

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
	encryptedPass, err := helper.Encrypt(cred.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}

	// generate user Id
	userId := helper.GenerateUUID().String()
	// prepare query statment
	stmt, err := db.Portfolio.Prepare("INSERT INTO users (id, username, password, email) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Printf("%v Could not insert into db", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not insert into db")
		return
	}

	defer stmt.Close()

	// execute query statement
	_, err = stmt.Exec(userId, cred.Username, encryptedPass, cred.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Username or email already exist")
		log.Printf("%v", err)
		return
	}

	userToken, err := helper.SignJWT(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to register user")
		log.Printf("%v", err)
		return
	}

	res := map[string]interface{}{
		"token": userToken,
	}
	helper.SendResponse(w, http.StatusOK, true, "Signup successful", res)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var cred auth.LoginCredentials

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
	row := db.Portfolio.QueryRow("SELECT id,username,password FROM users WHERE username= $1", cred.Username)

	// variable to store data from db
	var storedCred user.UserInfo

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
	userToken, err := helper.SignJWT(storedCred.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to register user")
		log.Printf("%v", err)
		return
	}

	res := map[string]interface{}{
		"token": userToken,
	}
	helper.SendResponse(w, http.StatusOK, true, "Signup successful", res)
}
