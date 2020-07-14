package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/chanyipiaomiao/hlog"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type UserMan struct {
	Client *http.Client
	Token  string
}

func NewUserMan() (*UserMan, error) {
	var (
		client                     = &http.Client{}
		jar                        *cookiejar.Jar
		err                        error
		resp                       *http.Response
		req                        *http.Request
		loginUrl                   *url.URL
		loginPostData              url.Values
		doc                        *goquery.Document
		token                      string
		secondAuthenticatePostData url.Values
	)

	// 1.访问登陆界面 获取cookie
	req, _ = http.NewRequest("GET", appConfig.API.LoginIndex, nil)
	if resp, err = client.Do(req); err != nil {
		return nil, err
	}
	resp.Body.Close()

	if jar, err = cookiejar.New(nil); err != nil {
		return nil, err
	}

	if loginUrl, err = url.Parse(appConfig.API.DoLogin); err != nil {
		return nil, err
	}
	jar.SetCookies(loginUrl, resp.Cookies())
	client.Jar = jar

	hlog.Info(nil, "get cookie success")

	// 2.登录
	loginPostData = url.Values{}
	loginPostData.Add("os_username", appConfig.AdminUser.Username)
	loginPostData.Add("os_password", appConfig.AdminUser.Password)
	loginPostData.Add("login", "登录")
	loginPostData.Add("os_destination", "")
	if _, err = client.PostForm(appConfig.API.DoLogin, loginPostData); err != nil {
		return nil, err
	}
	hlog.Info(nil, "login success")

	// 3.访问首页 获取token
	if resp, err = client.Get(appConfig.API.Prefix); err != nil {
		return nil, err
	}

	if doc, err = goquery.NewDocumentFromReader(resp.Body); err != nil {
		return nil, err
	}
	resp.Body.Close()

	token, _ = doc.Find("#atlassian-token").Attr("content")
	hlog.Info(hlog.D{"token": token}, "get token success")

	// 4.二次验证
	secondAuthenticatePostData = url.Values{}
	secondAuthenticatePostData.Add("password", appConfig.AdminUser.Password)
	secondAuthenticatePostData.Add("authenticate", "确认")
	secondAuthenticatePostData.Add("destination", "")
	if _, err = client.PostForm(appConfig.API.SecondAuthenticate, secondAuthenticatePostData); err != nil {
		return nil, err
	}

	hlog.Info(nil, "second authenticate success")

	return &UserMan{Client: client, Token: token}, nil
}

func (u *UserMan) Disable(username string) error {
	var (
		err                    error
		disableConfirmPostData url.Values
	)

	// 调用禁用确认接口
	disableConfirmPostData = url.Values{}
	disableConfirmPostData.Add("atl_token", u.Token)
	disableConfirmPostData.Add("username", username)
	disableConfirmPostData.Add("confirm", "确认")
	if _, err = u.Client.PostForm(appConfig.API.DisableConfirm, disableConfirmPostData); err != nil {
		return err
	}

	hlog.Info(hlog.D{"disable_username": username, "admin_user": appConfig.AdminUser.Username}, "disable user success")

	return nil
}

func (u *UserMan) Create(payload *CreateUserPayload) error {
	var (
		err            error
		createPostData url.Values
	)

	createPostData = url.Values{}
	createPostData.Add("atl_token", u.Token)
	createPostData.Add("username", payload.Username)
	createPostData.Add("fullName", payload.FullName)
	createPostData.Add("email", payload.Email)
	createPostData.Add("password", payload.Password)
	createPostData.Add("confirm", payload.Password)

	if _, err = u.Client.PostForm(appConfig.API.Create, createPostData); err != nil {
		return err
	}

	return nil
}

func (u *UserMan) Delete(username string) error {
	var (
		err            error
		deletePostData url.Values
	)

	deletePostData = url.Values{}
	deletePostData.Add("atl_token", u.Token)
	deletePostData.Add("username", username)
	deletePostData.Add("confirm", "确定")

	if _, err = u.Client.PostForm(appConfig.API.Delete, deletePostData); err != nil {
		return err
	}

	return nil
}
