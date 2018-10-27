package main

import (
	"circleTest/api/dbops"
	"circleTest/api/defs"
	"circleTest/api/session"
	"net/http"
)

var HEADER_FIELD_SESSION = "X-Session-Id"

//var HEADER_FIELD_UNAME = "X-User-Name"


func validateUserSession(r *http.Request) bool {
	
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}


	_, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	
	return true
}


func ValidateUser(w http.ResponseWriter, r *http.Request) (string, bool) {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return "", false
	}
	uname, ok := session.IsSessionExpired(sid)
	return uname, !ok
}

func IsAdmin(w http.ResponseWriter, r *http.Request) bool {
	name, _ := ValidateUser(w, r)
	role, err := dbops.GetUserRole(name)
	if err != nil {
		return false
	}

	if role != defs.ADMIN {
		return false
	}
	return true
}
