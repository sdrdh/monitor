package monitor

import (
	"net/http"
	"time"
)

// ServiceMonitor is the abstraction around http services
type ServiceMonitor struct {
	BaseMonitor
	client *http.Client
}

// NewServiceMonitor returns the struct to be passed to CheckServices
func NewServiceMonitor(name, uri string, duration time.Duration, alertFunc func() error, client *http.Client) *ServiceMonitor {
	s := &ServiceMonitor{}
	s.Name = name
	s.URI = uri
	s.Frequency = duration
	s.AlertFunc = alertFunc
	s.client = client
	return s
}

// Ping Checks whether the service is up or not
func (s *ServiceMonitor) Ping() bool {
	url := s.URI
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return false
	}
	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}
