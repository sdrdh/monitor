package monitor

import "time"

// BaseMonitor is what most tend to follow
type BaseMonitor struct {
	Name      string
	URI       string
	Frequency time.Duration
	AlertFunc func() error
}

// GetFrequency gets the frequency at which we should check the Mongo
func (m *BaseMonitor) GetFrequency() time.Duration {
	return m.Frequency
}

// GetServiceName is used for sending the alerts
func (m *BaseMonitor) GetServiceName() string {
	return m.Name
}

// Alert alerts the users
func (m *BaseMonitor) Alert() error {
	return m.AlertFunc()
}
