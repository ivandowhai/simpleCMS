package core

import "github.com/gorilla/sessions"

// TODO: use real key
var Store = sessions.NewCookieStore([]byte("key"))
