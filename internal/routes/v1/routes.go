package routes

import (
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/portfolio-builder/internal/controller/auth"
	"github.com/miriam-samuels/portfolio-builder/internal/controller/theme"
	"github.com/miriam-samuels/portfolio-builder/internal/controller/user"
)

func Routes(r *mux.Router) {

	// Authentication Routes
	r.HandleFunc("/signup", auth.SignUp).Methods("POST")
	r.HandleFunc("/signin", auth.SignIn).Methods("POST")

	// User Routes
	r.HandleFunc("/userinfo/{user}", user.GetUserInfo).Methods("GET")
	r.HandleFunc("/userinfo/{user}", user.SetUserInfo).Methods("PUT")

	//  Theme Routes
	r.HandleFunc("/theme/all", theme.GetThemes).Methods("GET")

}
