package mid

import (
	"encoding/json"
	"fmt"
	"log"
)

type JobRepository struct {
	client ClientIterface
}

func NewJobRepository(client ClientIterface) JobRepository {
	return JobRepository{client}
}

func (r JobRepository) GetJob() (job Job, ok bool, err error) {
	jobId, jobScript, err := r.client.Get()
	ok = jobId != 0
	if err != nil || !ok {
		return
	}
	job.Id = jobId
	job.Script = jobScript
	log.Println("Found Job: " + fmt.Sprint(job.Id) + ", " + job.Script)
	return
}

func (r JobRepository) UpdateJob(job Job) (err error) {
	body, _ := json.Marshal(job)
	err = r.client.Put(body)
	return
}
