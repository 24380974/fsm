package v2

import (
	"fmt"
	"time"
)

type xnetworkConfigJob struct {
	done chan struct{}
	//connectController connector.ConnectController
}

func (job *xnetworkConfigJob) GetDoneCh() <-chan struct{} {
	return job.done
}

func (job *xnetworkConfigJob) Run() {
	defer close(job.done)
	fmt.Println("xnetworkConfigJob:", time.Now())
	//c := job.connectController
	//c.Refresh()
}

func (job *xnetworkConfigJob) JobName() string {
	return "fsm-xnetworkConfig-job"
}
