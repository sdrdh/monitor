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

Alert does not take any parameters. Any parameters needed are expected to be sent in a closure implementing this function signature. This will give the ability to choose whatever the caller wants.

Examples can be found in the ```examples``` directory