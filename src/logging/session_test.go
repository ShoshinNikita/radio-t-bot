package logging

import (
	"testing"
	"time"
)

var s sessionsMap

func TestMain(m *testing.M) {
	s.init(6 * time.Hour, "0")

	m.Run()
}

func TestAdd(t *testing.T) {
	tests := []struct {
		sessionID string
		key       string
	}{
		{"hello", "1"},
		{"123", "2"},
		{"569", "3"},
	}

	c := "0"
	for _, tt := range tests {
		c = incStr(c)
		s.add(tt.sessionID, c)
	}

	if len(s.sessions) != len(tests) {
		t.Errorf("Bad size Want: %d Got: %d\n%v", len(tests), len(s.sessions), s.sessions)
		return
	}

	for i, tt := range tests {
		if res, ok := s.sessions[tt.sessionID]; !ok {
			t.Errorf("Test #%d There's no '%s' key", i, tt.key)
			continue
		} else if res.number != tt.key {
			t.Errorf("Test #%d There's no '%s' key", i, tt.key)
			continue
		}

		if n, err := s.getNumber(tt.sessionID); err != nil {
			t.Errorf("Test #%d Got error %s", i, err.Error())
		} else if n != tt.key {
			t.Errorf("Test #%d Want number: %s Got: %s", i, tt.key, n)
		}
	}
}

func TestDelete(t *testing.T) {
	addTests := []struct {
		sessionID string
		key       string
	}{
		{"hello", "1"},
		{"123", "2"},
		{"569", "3"},
	}
	delTests := []struct {
		sessionID string
	}{
		{"hello"},
		{"569"},
	}

	c := "0"
	for _, tt := range addTests {
		c = incStr(c)
		s.add(tt.sessionID, c)
	}

	if len(s.sessions) != len(addTests) {
		t.Errorf("Bad size Want: %d Got: %d\n%v", len(addTests), len(s.sessions), s.sessions)
		return
	}

	for i, tt := range delTests {
		s.delete(tt.sessionID)
		if _, ok := s.sessions[tt.sessionID]; ok {
			t.Errorf("Test #%d Key '%s' wasn't deleted", i, tt.sessionID)
		}
	}
}
