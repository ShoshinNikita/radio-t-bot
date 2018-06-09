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

	"github.com/boltdb/bolt"
)

// dd-mm-yyyy hh:mm:ss
const (
	formatTime          = "02-01-2006 15:04:05"
	errFileName         = "errors.log"
	reqFileName         = "requests.db"
	clearingHoursNumber = 6
)

var (
	logErrorsFile  *os.File
	logErrorsMutex *sync.Mutex
	logReqDB       *bolt.DB
	sessions       sessionsMap
)

// Init initializes files for logging
// path must consist path to the folder (not to the files)
// Name of the logErrorsFile – errors.log
//      of the logReqDB      – requests.db
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
	logReqDB, err = bolt.Open(path+reqFileName, 0600, nil)
	if err != nil {
		return err
	}
	var counter string
	err = logReqDB.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("requests"))
		// Add counter if it doesn't exist
		if v := b.Get([]byte("0")); v == nil {
			b.Put([]byte("0"), []byte("0"))
		}
		counter = string(b.Get([]byte("0")))

		return nil
	})

	// Init sessions
	sessions.init(time.Duration(clearingHoursNumber)*time.Hour, counter)
	go sessions.deleteOld()

	return err
}

func LogError(err error) {
	if err == nil {
		return
	}
	logErrorsMutex.Lock()
	defer logErrorsMutex.Unlock()

	t := time.Now()
	logErrorsFile.WriteString(fmt.Sprintf("[ERR] %s %s\n", t.Format(formatTime), err.Error()))
}

func LogRequest(command, sessionID string) {
	var (
		sessionNumber string
		ok            bool
	)

	if sessionNumber, ok = sessions.getNumber(sessionID); !ok {
		sessionNumber = sessions.add(sessionID)
	}
	time := time.Now().Format(formatTime)
	key := sessionNumber + "-" + sessionID

	logReqDB.Update(func(tx *bolt.Tx) error {
		requests := tx.Bucket([]byte("requests"))
		b := requests.Bucket([]byte(key))

		// Counter has key "0"
		counter := "1"
		if b == nil {
			// If bucket doesn't exits - create bucket
			b, _ = requests.CreateBucket([]byte(key))
			// Change counter of sessions. The counter was increase when we tried to getNumber()
			// and got error, then sessions.add() was called
			requests.Put([]byte("0"), []byte(sessions.getCounter()))
		} else {
			// If bucket exists - increase counter
			counter = string(b.Get([]byte("0")))
			counter = incStr(counter)
		}
		b.Put([]byte("0"), []byte(counter))

		text := time + " - " + command
		b.Put([]byte(counter), []byte(text))

		return nil
	})
}
