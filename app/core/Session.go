package core

import (
	"github.com/gorilla/sessions"
	"net/http"
)

// TODO: use real key
var store = sessions.NewCookieStore([]byte("key"))

func SessionGet(request *http.Request, name string) *sessions.Session {
	logger := Logger{}
	logger.Init()
	session, err := store.Get(request, name)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
	}

	return session
}
