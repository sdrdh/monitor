package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sdrdh/monitor"
)

type Service struct {
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	URI      string `json:"uri,omitempty"`
	Duration int    `json:"duration,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage is monitor {fileName}")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	var ss []Service
	err = json.NewDecoder(file).Decode(&ss)
	if err != nil {
		panic(err)
	}
	var pp []monitor.PingChecker
	for _, s := range ss {
		switch strings.ToUpper(s.Type) {
		case "MONGO":
			m := monitor.NewMongoMonitor(s.Name, s.URI, time.Duration(s.Duration)*time.Second, AlertFunc(s.Name))
			pp = append(pp, m)
		case "SERVICE":
			client := &http.Client{
				Timeout: 10 * time.Second,
			}
			s := monitor.NewServiceMonitor(s.Name, s.URI, time.Duration(s.Duration)*time.Second, AlertFunc(s.Name), client)
			pp = append(pp, s)
		}
	}
	monitor.CheckServices(pp)
}

func AlertFunc(name string) func() error {
	return func() error {
		fmt.Println("Trying to alert for", name)
		return nil
	}
}
