package sessionConfig

import "github.com/gorilla/sessions"

var store *sessions.CookieStore

func NewStore() {
	store = sessions.NewCookieStore([]byte("my-secret-key"))
}
