package http

import (
	"encoding/json"
	"net/http"
)

type Auth struct {
	baseUrl string
	body    []byte
}

func NewAuth(conf Config) Auth {
	body := map[string]string{"username": conf.GetUsername(), "password": conf.GetPassword()}
	js, _ := json.Marshal(body)
	return Auth{conf.GetAuthUrl(), js}
}

func (a Auth) GetToken() (token string, err error) {
	body, err := Request{}.Do(http.MethodPost, a.baseUrl, "", a.body)
	if err != nil {
		return
	}
	response := AuthResponse{}
	json.Unmarshal(body, &response)
	token = response.Data.Token
	return
}

type AuthResponse struct {
	Data   AuthDataResponse `json:"data"`
	Status string           `json:"status"`
}

type AuthDataResponse struct {
	Token string `json:"auth_key"`
}
