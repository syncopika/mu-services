package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Heartbeat struct {
	ServiceName   string  `json:"serviceName"`
	ServiceStatus string  `json:"serviceStatus"`
	Timestamp     string  `json:"timestamp"`    
}

func createHeartbeatMsg(service string) []byte {
	newHeartbeat := Heartbeat{
		ServiceName: service,
		ServiceStatus: "good",
		Timestamp: fmt.Sprint("", time.Now().UTC()),
	}
	
	// struct fields need to be capitalized when using json.Marshal!
	jsonBody, err := json.Marshal(newHeartbeat)
	
	if err != nil {
		fmt.Println("error in createHeartbeatMsg! ", err)
	}
	
	return jsonBody
}

func main() {

	// emit heartbeat every 10 sec
	observerEndpoint := "http://host.docker.internal:3000/serviceUpdate"
	ticker := time.NewTicker(10000 * time.Millisecond)
	done := make(chan bool)
	
	go func() {
		for {
			select {
			case <-done: // read value from done channel
				return
			case _ = <-ticker.C:
				// send heartbeat
				newHeartbeat := createHeartbeatMsg("timekeeper")
				
				resp, err := http.Post(observerEndpoint, "application/json", bytes.NewBuffer(newHeartbeat))
				
				if err != nil {
					fmt.Println(err)
				} else {
					defer resp.Body.Close()
				}	
			}
		}
	}()

	// server api functions
	handleRequest := func(wr http.ResponseWriter, _ *http.Request) {
		io.WriteString(wr, "Hello there! This is the timekeeper service.\n")
	}
	
	handleTimeRequest := func(wr http.ResponseWriter, _ *http.Request) {
		timeNow := fmt.Sprint("", time.Now().UTC())
		io.WriteString(wr, timeNow)
	}

	// handle routes
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/getTime", handleTimeRequest)
	
	// launch server
	fmt.Println("serving on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

