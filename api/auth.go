package main

import (
	"net/http"
	"github.com/myproject/video_server/api/session"
	"github.com/myproject/video_server/api/defs"
)

var HEADER_FILED_SESSION = "X-Session_Id"
var HEADER_FIELD_UNAME = "X-User_Name"

func ValidateUserSession(r *http.Request) bool{
	sid := r.Header.Get(HEADER_FILED_SESSION)
	if len(sid) == 0{
		return false
	}

	uname , ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}

	r.Header.Add(HEADER_FIELD_UNAME,uname)
	return true ;
}

func ValidateUser(w http.ResponseWriter , r *http.Request) bool{
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w,defs.ErrorNotAuthUser)
		return false
	}
	return true
}