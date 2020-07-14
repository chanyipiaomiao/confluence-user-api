package main

import (
	"encoding/json"
	"fmt"
	"github.com/chanyipiaomiao/hlog"
	"io/ioutil"
	"net/http"
)

func JsonRespOK(w http.ResponseWriter, msg string, data interface{}) {
	r, _ := json.Marshal(&CommonResp{
		Code: 0,
		Msg:  msg,
		Data: data,
	})

	hlog.Info(hlog.D{"data": data}, msg)

	fmt.Fprintln(w, string(r))
}

func JsonRespError(w http.ResponseWriter, msg string, data interface{}) {
	r, _ := json.Marshal(&CommonResp{
		Code: 1,
		Msg:  msg,
		Data: data,
	})

	hlog.Error(hlog.D{"data": data}, msg)

	fmt.Fprintln(w, string(r))
}

func DisableHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		username string
		userMan  *UserMan
	)

	if err = r.ParseForm(); err != nil {
		JsonRespError(w, fmt.Sprintf("parse form error: %s", err), nil)
		return
	}

	username = r.Form.Get("username")
	if username == "" {
		JsonRespError(w, "expect username, got null", nil)
		return
	}

	if userMan, err = NewUserMan(); err != nil {
		JsonRespError(w, fmt.Sprintf("init userman error: %s", err), nil)
		return
	}

	if err = userMan.Disable(username); err != nil {
		JsonRespError(w, fmt.Sprintf("disable user: %s error: %s", username, err), nil)
		return
	}

	JsonRespOK(w, fmt.Sprintf("disable user: %s success", username), nil)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		payload CreateUserPayload
		body    []byte
		userMan *UserMan
	)

	if body, err = ioutil.ReadAll(r.Body); err != nil {
		JsonRespError(w, fmt.Sprintf("read request body error: %s", err), nil)
		return
	}

	if err = json.Unmarshal(body, &payload); err != nil {
		JsonRespError(w, fmt.Sprintf("json.Unmarshal request body error: %s", err), nil)
		return
	}

	if userMan, err = NewUserMan(); err != nil {
		JsonRespError(w, fmt.Sprintf("init userman error: %s", err), nil)
		return
	}

	if err = userMan.Create(&payload); err != nil {
		JsonRespError(w, fmt.Sprintf("create user: %s error: %s", payload.Username, err), nil)
		return
	}

	JsonRespOK(w, fmt.Sprintf("create user: %s success", payload.Username), nil)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		username string
		userMan  *UserMan
	)

	if err = r.ParseForm(); err != nil {
		JsonRespError(w, fmt.Sprintf("parse form error: %s", err), nil)
		return
	}

	username = r.Form.Get("username")
	if username == "" {
		JsonRespError(w, "expect username, got null", nil)
		return
	}

	if userMan, err = NewUserMan(); err != nil {
		JsonRespError(w, fmt.Sprintf("init userman error: %s", err), nil)
		return
	}

	if err = userMan.Delete(username); err != nil {
		JsonRespError(w, fmt.Sprintf("delete user: %s error: %s", username, err), nil)
		return
	}

	JsonRespOK(w, fmt.Sprintf("delete user: %s success", username), nil)
}
