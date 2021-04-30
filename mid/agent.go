package mid

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type MidAgent struct {
	repository JobRepository
}

func NewMidAgent(repository JobRepository) MidAgent {
	return MidAgent{repository}
}

func (m MidAgent) Run() {
	poolTime, _ := strconv.Atoi(os.Getenv("MID_POOL_TIME"))
	ticker := time.NewTicker(time.Duration(poolTime) * time.Second)
	defer ticker.Stop()
	log.Println("Mid-agent started. Wait for Job")
	for {
		select {
		case <-ticker.C:
			if job, ok, err := m.repository.GetJob(); ok {
				if err != nil {
					log.Println(err)
				} else {
					m.runJob(job)
				}
			}
		}
	}
}

func (m MidAgent) runJob(job Job) {
	stdout, exitCode := m.runCmd(job.Script)
	job.Result = stdout
	job.ExitCode = exitCode
	job.Log()
	err := m.repository.UpdateJob(job)
	if err != nil {
		log.Println("Error update job: " + err.Error())
	}
}

func (m MidAgent) runCmd(script string) (stdout string, exitCode int) {
	cmd := m.getCmd(script)
	cmdOut, err := cmd.Output()
	if err != nil {
		exitCode = 1
		stdout = err.Error()
	} else {
		stdout = string(cmdOut)
	}
	log.Println(stdout)
	return
}

func (m MidAgent) getCmd(script string) (cmd *exec.Cmd) {
	script = strings.ReplaceAll(script, "\n", ";")
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("powershell", "-Command", script)
	case "linux":
		cmd = exec.Command("bash", "-c", script)
	default:
		log.Fatalln("OS not implimented")
	}
	return
}
