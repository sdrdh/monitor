package monitor

import (
	"database/sql"
	"time"
)

// SQLDBMonitor is the instance of SQL Monitor
type SQLDBMonitor struct {
	sqlType string
	BaseMonitor
}

// NewSQLMonitor returns the new PingChecker for sql types
func NewSQLMonitor(name, uri, sqlType string, duration time.Duration, alertFunc func() error) *SQLDBMonitor {
	s := &SQLDBMonitor{}
	s.Name = name
	s.URI = uri
	s.Frequency = duration
	s.AlertFunc = alertFunc
	s.sqlType = sqlType
	return s
}

// Ping checks whether the SQL type DB is up or down
func (s *SQLDBMonitor) Ping() bool {
	db, err := sql.Open(s.sqlType, s.URI)
	if err != nil {
		return false
	}
	if err := db.Ping(); err != nil {
		return false
	}
	return true
}
