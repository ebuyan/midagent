package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Request struct{}

func (r Request) Do(method, url, token string, data []byte) (respBody []byte, err error) {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if len(token) > 0 {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respBody, err = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		respError := ErrorResponse{}
		json.Unmarshal(respBody, &respError)
		if len(respError.Errors) > 0 {
			err = errors.New(respError.Errors[0])
		} else {
			err = errors.New("Api client request error.")
		}
	}
	return
}

type ErrorResponse struct {
	Status string   `json:"status"`
	Errors []string `json:"errors"`
}
