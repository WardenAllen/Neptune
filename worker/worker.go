package worker

import (
	"fmt"
)

type Worker struct {
	WorkerPool chan chan Job
	JobPool chan Job
	Quit chan bool
}

func NewWorker(workPool chan chan Job) Worker {
	return Worker {
		WorkerPool:workPool,
		JobPool:make(chan Job),
		Quit:make(chan bool),
	}
}

func (w Worker) Start(idx int) {
	go func(idx int) {
		for {
			// put worker's job queue into worker queue.
			// wait for next job.
			w.WorkerPool <- w.JobPool
			select {
			case job := <- w.JobPool:
				if err := job.Exec(idx); err != nil {
					fmt.Printf("Excute job failed with err: %v", err)
				}
			case <- w.Quit:
				return
			}
		}
	}(idx)
}
