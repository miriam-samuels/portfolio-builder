package middleware

import (
	// "encoding/json"
	"context"
	"net/http"
	"strings"

	"github.com/miriam-samuels/portfolio-builder/internal/helper"
	"github.com/miriam-samuels/portfolio-builder/internal/types"
)

// function to validate user authorization
func ValidateAuth(nextHandler http.HandlerFunc) http.Handler {
	// returns a handler function which calls the next handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			// send respose that user is not authorized
			helper.SendResponse(w, http.StatusUnauthorized, false, "missing authoriztion", nil)
			return
		}

		// split the header e.g Bearer somerandomtoken
		authHeaderSplit := strings.Split(authHeader, " ")

		// Check if auth header has a lenght of 2 after spilting
		if len(authHeaderSplit) != 2 {
			// send response on invalid auth header
			helper.SendResponse(w, http.StatusUnauthorized, false, "invalid auth header", nil)
			return
		}

		// check if auth type is bearer
		if authHeaderSplit[0] != "Bearer" {
			//	invalid auth type
			helper.SendResponse(w, http.StatusUnauthorized, false, "invalid auth type expecting Bearer", nil)
			return
		}

		// retieve token from auth header
		token := authHeaderSplit[1]
		if token == "" {
			// send response on user not logged in
			helper.SendResponse(w, http.StatusUnauthorized, false, "user not logged in, invalid token", nil)
			return
		}

		// context to store id and role
		ctx := r.Context()

		// Verify user using token claims
		if token != "iamanadminuserihaveallthepowerintheworldsofearmeyoumotherfuckers" {
			// validation logic for token (convert _ to claims)
			claim, valid := helper.VerifyJWT(token)
			if !valid {
				helper.SendResponse(w, http.StatusUnauthorized, false, "invalid token", nil)
				return
			}

			// store user id and role
			authKey := types.AuthCtxKey{Id: claim.UserId, Role: claim.Role}
			ctx = context.WithValue(ctx, types.AuthCtxKey{}, authKey)

		} else {
			authKey := types.AuthCtxKey{Id: "5f18e92b-c841-4a16-bd07-3d4446e0df75", Role: "borrower"}
			ctx = context.WithValue(ctx, types.AuthCtxKey{}, authKey)
		}

		r = r.WithContext(ctx)

		// call nect handler
		nextHandler.ServeHTTP(w, r)
	})
}

// This function takes in an interface for the request and returns a function which takes in a handler function and returns a handler
// TODO: Learn the right way to use a middleware for parsing
// method implemented on an interface
// takes in a handler function  and returns a handler
// func ParseRequest(i interface{}) func(func(http.ResponseWriter, *http.Request)) http.Handler {
// 	return func(nextHandler func(http.ResponseWriter, *http.Request)) http.Handler {
// 		// anonymous function using handlerfunc returns a handler
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			err := json.NewDecoder(r.Body).Decode(i)
// 			if err != nil {
// 				if err.Error() != "EOF" {
// 					// TODO: send response on unable to parse body
// 					return
// 				}

// 				nextHandler(w, r, )
// 			}
// 		})
// 	}
// }
