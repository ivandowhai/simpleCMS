package core

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

// TODO: use real key
var store = sessions.NewCookieStore([]byte("key"))

func SessionGet(request *http.Request, name string) *sessions.Session {
	session, err := store.Get(request, name)
	if err != nil {
		fmt.Println(err.Error())
	}

	return session
}
