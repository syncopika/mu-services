package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
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

