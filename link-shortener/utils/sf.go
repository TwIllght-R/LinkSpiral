package utils

import (
	"errors"
	"sync"
	"time"
)

// Snowflake struct represents the Snowflake generator
type Snowflake struct {
	mutex        sync.Mutex
	nodeID       int64
	sequence     int64
	lastUnixTime int64
}

// NewSnowflake initializes and returns a new Snowflake generator
func NewSnowflake(nodeID int64) *Snowflake {
	return &Snowflake{
		nodeID: nodeID,
	}
}

// GenerateID generates a unique Snowflake ID
func (s *Snowflake) GenerateID() (int64, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	currentTime := time.Now().UnixNano() / 1000000 // convert to milliseconds
	if currentTime < s.lastUnixTime {
		return 0, errors.New("clock moved backwards")
	}

	if currentTime == s.lastUnixTime {
		s.sequence++
		if s.sequence > 4095 {
			time.Sleep(time.Millisecond)
			currentTime = time.Now().UnixNano() / 1000000
			s.sequence = 0
		}
	} else {
		s.sequence = 0
	}

	s.lastUnixTime = currentTime

	id := (currentTime << 22) | (s.nodeID << 12) | s.sequence
	return id, nil
}
