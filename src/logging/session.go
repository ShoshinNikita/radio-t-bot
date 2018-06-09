package logging

import (
	"strconv"
	"sync"
	"time"
)

type sessionStruct struct {
	number       string
	creationTime time.Time
}

type sessionsMap struct {
	counter             string
	mutex               *sync.RWMutex
	sessions            map[string]sessionStruct
	timeBetweenClearing time.Duration
}

func (m *sessionsMap) init(timeBetweenClearing time.Duration, counter string) {
	m.counter = counter
	m.mutex = new(sync.RWMutex)
	m.sessions = make(map[string]sessionStruct)
	m.timeBetweenClearing = timeBetweenClearing
}

func (m *sessionsMap) add(session string) string {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.counter = incStr(m.counter)
	m.sessions[session] = sessionStruct{number: m.counter, creationTime: time.Now()}
	return m.counter
}

func (m *sessionsMap) delete(session string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.sessions, session)
}

// Should be runned when initialize map
func (m *sessionsMap) deleteOld() {
	ticker := time.NewTicker(m.timeBetweenClearing)

	// Delete every timeBetweenClearing
	for range ticker.C {
		m.mutex.Lock()
		defer m.mutex.Unlock()

		const maxHoursNumber = 2

		now := time.Now()
		for i, k := range m.sessions {
			if now.Sub(k.creationTime).Hours() > maxHoursNumber {
				delete(m.sessions, i)
			}
		}
	}
}

func (m *sessionsMap) getNumber(session string) (number string, ok bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if number, ok := m.sessions[session]; ok {
		return number.number, true
	}

	return "", false
}

func (m *sessionsMap) getCounter() string {
	return m.counter
}

// Increase string (s – int64, which was converted to string)
func incStr(s string) string {
	n, _ := strconv.ParseInt(s, 10, 64)
	n++
	return strconv.FormatInt(n, 10)
}
