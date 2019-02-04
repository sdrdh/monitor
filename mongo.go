package monitor

import (
	"time"

	"gopkg.in/mgo.v2"
)

// MongoMonitor struct is for mongo pings
type MongoMonitor struct {
	BaseMonitor
}

// Ping function checks if the mongo store is up or not
func (m *MongoMonitor) Ping() bool {
	session, err := mgo.Dial(m.URI)
	if err != nil {
		return false
	}
	if session.Ping() != nil {
		return false
	}
	return true
}

// NewMongoMonitor function returns a new instance for mongo monitoring
func NewMongoMonitor(name, uri string, duration time.Duration, alertFunc func() error) *MongoMonitor {
	m := &MongoMonitor{}
	m.Name = name
	m.URI = uri
	m.Frequency = duration
	m.AlertFunc = alertFunc
	return m
}
