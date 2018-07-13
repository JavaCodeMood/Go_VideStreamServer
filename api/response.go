package main

import(
	"io"
	"net/http"
	"encoding/json"
	"github.com/myproject/video_server/api/defs"
)

func sendErrorResponse(w http.ResponseWriter,errResp defs.ErrorResponse){
	w.WriteHeader(errResp.HttpSC)
	resStr , _ :=json.Marshal(&errResp.Error)  // obj -> json
	io.WriteString(w,string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string , sc int){
	w.WriteHeader(sc)
	io.WriteString(w,resp)
}