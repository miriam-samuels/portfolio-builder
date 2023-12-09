package auth

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/miriam-samuels/portfolio-builder/internal/db"
	"github.com/miriam-samuels/portfolio-builder/internal/helper"
	"github.com/miriam-samuels/portfolio-builder/internal/models/auth"
	"github.com/miriam-samuels/portfolio-builder/internal/models/user"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("yay entered")
	cred := &auth.SignUpCredentials{}

	// decode request body and store in cred variable
	err := helper.ParseRequestBody(w, r, cred)
	if err != nil {
		helper.SendResponse(w, http.StatusBadRequest, false, "error parsing body:"+err.Error(), nil, err)
		return
	}
	// validate body to ensure all properties are available
	err = cred.ValidateSignUp()
	if err != nil {
		helper.SendResponse(w, http.StatusBadRequest, false, err.Error(), nil)
		return
	}

	var exists bool
	err = db.Portfolio.QueryRow("SELECT 1 FROM users WHERE email=$1", cred.Email).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered", nil, err)
		return
	}

	// send response that user exists
	if exists {
		helper.SendResponse(w, http.StatusBadRequest, false, "user exists", nil)
		return
	}

	// encrypt password
	encryptedPass, err := helper.Encrypt(cred.Password)
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	// generate user Id
	userId := helper.GenerateUUID().String()
	// prepare query statment
	stmt, err := db.Portfolio.Prepare("INSERT INTO users (id, username, password, email) VALUES ($1, $2, $3, $4)")
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, err.Error(), nil)
		return
	}

	defer stmt.Close()

	// execute query statement
	_, err = stmt.Exec(userId, cred.Username, encryptedPass, cred.Email)
	if err != nil {
		helper.SendResponse(w, http.StatusBadRequest, false, "Username or email already exist", nil)
		return
	}

	userToken, err := helper.SignJWT(userId)
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "Failed to register user", nil, err)
		return
	}

	res := map[string]interface{}{
		"token": userToken,
	}
	helper.SendResponse(w, http.StatusOK, true, "Signup successful", res)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	cred := &auth.LoginCredentials{}

	// decode request body and store in cred variable
	err := helper.ParseRequestBody(w, r, cred)
	if err != nil {
		helper.SendResponse(w, http.StatusBadRequest, false, "error parsing body:"+err.Error(), nil, err)
		return
	}

	// validate body to ensure all properties are available
	err = cred.ValidateLogin()
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
	err = helper.CompareHashAndString(storedCred.Password, cred.Password)
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
		"user":  storedCred,
	}
	helper.SendResponse(w, http.StatusOK, true, "Signin successful", res)
}
