package routes

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/src/controllers/auth"
	userController "github.com/miriam-samuels/src/controllers/user"
)

func RoutesV1(r *mux.Router) {

	// Authentication Routes
	r.HandleFunc("/signup", authControllers.SignUp).Methods("POST")
	r.HandleFunc("/signin", authControllers.SignIn).Methods("POST")

	// User Routes
	r.HandleFunc("/userinfo/{user}", userController.GetUserInfo).Methods("GET")
	r.HandleFunc("/userinfo/{user}", userController.SetUserInfo).Methods("PUT")

}
