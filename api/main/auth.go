package main

import (
	"circleTest/api/dbops"
	"circleTest/api/defs"
	"circleTest/api/session"
	"net/http"
)

var HEADER_FIELD_SESSION = "X-Session-Id"

//var HEADER_FIELD_UNAME = "X-User-Name"

//验证登录信息是否有效
func validateUserSession(r *http.Request) bool {
	//获取请求头中的session ID
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	//获取session 并判断是否过期
	_, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	// //session有效，将uname写入请求头 相当于分发访问权限
	// r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

// 用于限定为本人时，如发布帖子
//
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
