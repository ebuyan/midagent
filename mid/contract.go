package mid

type ClientIterface interface {
	Get() (jobId int, jobScript string, err error)
	Put(body []byte) error
}
