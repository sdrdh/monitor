# Monitor
A package which contains simple ping checker for most of the DBs and Services. Can be used for Black box monitoring.

Package contains just one interface which can be implemented
```go
type PingChecker interface {
	Ping() bool
	GetFrequency() time.Duration
	Alert() error
	GetServiceName() string
}
```