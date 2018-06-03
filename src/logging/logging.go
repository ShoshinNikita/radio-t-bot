package logging

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const formatTime = "2006-01-02 15:04:05"

var logFile *os.File
var mutex *sync.Mutex

func Init(path string) (err error) {
	mutex = new(sync.Mutex)
	logFile, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	return err
}

func LogError(err error) {
	if err == nil {
		return
	}
	mutex.Lock()
	defer mutex.Unlock()

	t := time.Now()
	logFile.WriteString(fmt.Sprintf("[ERR] %s %s\n", t.Format(formatTime), err.Error()))
}

func LogRequest(req string) {
	mutex.Lock()
	defer mutex.Unlock()

	t := time.Now()
	logFile.WriteString(fmt.Sprintf("[REQ] %s %s\n", t.Format(formatTime), req))
}
