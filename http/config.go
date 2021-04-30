package http

import (
	"os"
)

type Config struct{}

func (c Config) GetUsername() string {
	return os.Getenv("MID_API_USERNAME")
}

func (c Config) GetPassword() string {
	return os.Getenv("MID_API_PASSWORD")
}

func (c Config) GetAuthUrl() string {
	return os.Getenv("MID_API_ENDPOINT") + "/v1/auth/login"
}

func (c Config) GetJobUrl() string {
	return os.Getenv("MID_API_ENDPOINT") + "/v1/mid-server/job/" + os.Getenv("MID_API_SERVERID")
}
