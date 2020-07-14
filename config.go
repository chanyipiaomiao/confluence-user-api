package main

import (
	"encoding/json"
	"io/ioutil"
)

type API struct {
	Prefix             string `json:"prefix"`
	LoginIndex         string `json:"login_index"`
	DoLogin            string `json:"do_login"`
	SecondAuthenticate string `json:"second_authenticate"`
	DisableConfirm     string `json:"disable_confirm"`
	Create             string `json:"create"`
	Delete             string `json:"delete"`
}

type AdminUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AppConfig struct {
	Listen    string     `json:"listen"`
	API       *API       `json:"api"`
	AdminUser *AdminUser `json:"admin_user"`
}

var (
	appConfig *AppConfig
)

func InitConfig(filename string) (*AppConfig, error) {
	var (
		err     error
		content []byte
	)

	if content, err = ioutil.ReadFile(filename); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(content, &appConfig); err != nil {
		return nil, err
	}

	appConfig.API.LoginIndex = appConfig.API.Prefix + appConfig.API.LoginIndex
	appConfig.API.DoLogin = appConfig.API.Prefix + appConfig.API.DoLogin
	appConfig.API.SecondAuthenticate = appConfig.API.Prefix + appConfig.API.SecondAuthenticate
	appConfig.API.DisableConfirm = appConfig.API.Prefix + appConfig.API.DisableConfirm
	appConfig.API.Create = appConfig.API.Prefix + appConfig.API.Create
	appConfig.API.Delete = appConfig.API.Prefix + appConfig.API.Delete

	return appConfig, nil
}
