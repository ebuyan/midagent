package main

import (
	"log"
	"midagent/http"
	"midagent/mid"

	"github.com/joho/godotenv"
)

func main() {
	bootstrap()

	httpClient := http.NewClient()
	jobRepository := mid.NewJobRepository(httpClient)
	midAgent := mid.NewMidAgent(jobRepository)

	midAgent.Run()
}

func bootstrap() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
		return
	}
}
