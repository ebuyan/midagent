package mid

import (
	"fmt"
	"log"
)

type Job struct {
	Id       int    `json:"jobId"`
	Script   string `json:"script"`
	Result   string `json:"response"`
	ExitCode int    `json:"exitCode"`
}

func (j Job) Log() {
	if j.ExitCode == 0 {
		log.Println("Job " + fmt.Sprint(j.Id) + " complete")
	} else {
		log.Println("Job " + fmt.Sprint(j.Id) + " failed")
	}
}
