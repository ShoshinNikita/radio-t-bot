// Package logging helps to log all events
//
// For logging errors the program uses errors.log
// For logging requests the program uses requests.db (boltDB)
// Structure of requests.db:
// requests
//      |-  0 (int64 -> string) – counter
// 		|-> requestID (== counter + "-" + session_id)
//          |- 0 - (int64 -> string) – counter (key always == 0). Shows the number of requests in a session
//			|- number - time + command
//          |- etc
// Counters have keys "0" for showing of them firstly
//
package logging

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// dd-mm-yyyy hh:mm:ss
const (
	formatTime          = "02-01-2006 15:04:05"
	errFileName         = "errors.log"
	reqFileName         = "requests.log"
	clearingHoursNumber = 6
)

var (
	logErrorsFile  *os.File
	logErrorsMutex *sync.Mutex
	logReqFile     *os.File
	logReqMutex    *sync.Mutex
)

// Init initializes files for logging
// path must consist path to the folder (not to the files)
// Name of the logErrorsFile – errors.log
//      of the logReqFile    - requests.log
func Init(path string) (err error) {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	// Init logErrors
	logErrorsMutex = new(sync.Mutex)
	logErrorsFile, err = os.OpenFile(path+errFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	// Init logReq
	logReqMutex = new(sync.Mutex)
	logReqFile, err = os.OpenFile(path+reqFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	return err
}

func LogError(err error) {
	if err == nil {
		return
	}
	logErrorsMutex.Lock()
	defer logErrorsMutex.Unlock()

	t := time.Now()
	logErrorsFile.WriteString(fmt.Sprintf("[ERR] %s Error: %s\n", t.Format(formatTime), err.Error()))
}

func LogRequest(command, sessionID string) {
	logReqMutex.Lock()
	defer logReqMutex.Unlock()

	t := time.Now()
	logReqFile.WriteString(fmt.Sprintf("[REQ] %s Session: %s Req: %s\n", t.Format(formatTime), sessionID, command))
}
