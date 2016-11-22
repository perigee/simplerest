package terrant

import (
	"os"
)


var (
	MaxWorker = os.Getenv("MAX_WORKER")
	MaxQueue  = os.Getenv("MAX_Queue")
)

type Job struct {
	Payload Payload
}

type Worker struct {
	WorkerPool  chan chan Job
	JobChannnel chan Job
	quit        chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit: make(chan bool)
	}
}

func (w Worker) Start() {

    go func ()  {
        for {
            w.WorkerPool <- w.JobChannnel

            select {
                case job := <- w.JobChannnel:
                    if err := job.Payload.Provision(); err != nil {
                        log.Errorf("%s", err.Error())
                    }

                case <- w.quit:
                    return

            }
        }
    }()

}

func (w Worker) Stop(){
    go func() {
        w.quit <- true
    }()

}