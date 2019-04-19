package cron

import (
	"fmt"
	"os/exec"
	"bytes"
)

type CronJob struct {
	Command string
}

func GetJob(Command string) *CronJob {
	return &CronJob{
		Command: Command,
	}
}

func (self *CronJob) Run() {
	cmd := exec.Command("sh", "-c", self.Command)

	var outPut bytes.Buffer

	cmd.Stdout = &outPut
	cmd.Stderr = &outPut

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("%+v \n", string(outPut.Bytes()))
}
