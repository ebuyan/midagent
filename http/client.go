package http

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
)

type Client struct {
	conf Config
	auth Auth
}

func NewClient() Client {
	conf := Config{}
	return Client{conf, NewAuth(conf)}
}

func (c Client) Get() (jobId int, jobScript string, err error) {
	token, err := c.auth.GetToken()
	if err != nil {
		return
	}
	resp, err := c.sendRequest(http.MethodGet, token, nil)
	if err != nil {
		return
	}
	script, _ := base64.StdEncoding.DecodeString(resp.Data.Job.Script)
	jobScript = string(script)
	jobId = resp.Data.Job.JobId
	return
}

func (c Client) Put(data []byte) (err error) {
	token, err := c.auth.GetToken()
	if err != nil {
		return
	}
	_, err = c.sendRequest(http.MethodPut, token, data)
	return
}

func (c Client) sendRequest(method, token string, data []byte) (response ClientResponse, err error) {
	body, err := Request{}.Do(method, c.conf.GetJobUrl(), token, data)
	if err != nil {
		return
	}
	json.Unmarshal(body, &response)
	return
}

type ClientResponse struct {
	Data   ResponseData `json:"data"`
	Status string       `json:"status"`
}

type ResponseData struct {
	Job ResponseJob `json:"job"`
}

type ResponseJob struct {
	JobId  int    `json:"jobId"`
	Script string `json:"script"`
}
