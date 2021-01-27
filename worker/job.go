package worker

import (
	"runtime"
)

var (
	MaxWorker = runtime.NumCPU()
	MaxQueue = 10000
)

type Job interface {
	Exec(idx int) error
}

var JobQueue chan Job