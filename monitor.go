package monitor

import (
	"log"
	"sync"
	"time"
)

// PingChecker is the interface that the services have to implement
type PingChecker interface {
	Ping() bool
	GetFrequency() time.Duration
	Alert() error
	GetServiceName() string
}

func checkService(p PingChecker) {
	ticker := time.NewTicker(p.GetFrequency())
	for range ticker.C {
		if up := p.Ping(); !up {
			err := p.Alert()
			if err != nil {
				log.Println("Error Pinging: ", p.GetServiceName(), err)
				ticker.Stop()
			}
		}
	}
}

// CheckServices is what checks the services
func CheckServices(pp []PingChecker) {
	wg := sync.WaitGroup{}
	for _, p := range pp {
		wg.Add(1)
		go func(x PingChecker) {
			defer wg.Done()
			checkService(x)
		}(p)
	}
	wg.Wait()
}
