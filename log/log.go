package log

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Mode int32

const (
	ModeDebug	Mode = 1
	ModeRelease	Mode = 2
)

var (
	curMode = ModeDebug
	mtx sync.Mutex
)

func SetMode(m Mode)  {
	curMode = m
}

func log(level string, format string, args ...interface{}) {
	mtx.Lock()
	defer mtx.Unlock()
	_, _ = fmt.Fprintf(os.Stdout, time.Now().Format("2006-01-02 15:04:05 "))
	_, _ = fmt.Fprintf(os.Stdout, "[%s] ", level)
	_, _ = fmt.Fprintf(os.Stdout, format, args...)
	_, _ = fmt.Fprintln(os.Stdout)
}

func Debug(format string, args ...interface{}) {
	if curMode == ModeRelease {
		return
	}
	log("DEBUG", format, args...)
}

func Info(format string, args ...interface{}) {
	log("INFO ", format, args...)
}

func Warn(format string, args ...interface{}) {
	log("WARN ", format, args...)
}

func Error(format string, args ...interface{}) {
	log("ERROR", format, args...)
}

func Proto(proto string, success bool, format string, args ...interface{}) {

	var status string

	if success {
		status = "*SUCCESS*"
	} else {
		status = "*FAIL*"
	}

	mtx.Lock()
	defer mtx.Unlock()
	_, _ = fmt.Fprintf(os.Stdout, time.Now().Format("2006-01-02 15:04:05 "))
	_, _ = fmt.Fprintf(os.Stdout, "[INFO ] %s %s ", status, proto)
	_, _ = fmt.Fprintf(os.Stdout, format, args...)
	_, _ = fmt.Fprintln(os.Stdout)
}