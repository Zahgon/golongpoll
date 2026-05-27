// This example uses the FilePersistorAddOn to persist event data to file,
// allowing us to retain events across multiple program runs.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jcuga/golongpoll"
)

func main() {
	// Tell http.ServeFile where to get the client js file.
	// If running from root of checkout, use: -clientJs ./js-client/client.js
	// if in ./examples, use: -clientJs ../js-client/client.js
	// if in ./examples/filepersist, use: -clientJs ../../js-client/client.js
	// If using go 1.16 or higher, can simply use the embed directive instead of having to do this here, but supporting older versions.
	staticClientJs := flag.String("clientJs", "./js-client/client.js", "where the static js-client/client.js is located relative to where this binary runs")
	persistFilename := flag.String("persistTo", "./filepersist_example.data", "where to store event json data.")
	flag.Parse()

	// Note the two options after filename: writeBufferSize and writeFlushPeriodSeconds.
	// Instead of immediately writing data to disk, it is buffered with periodic flushes.
	filePersistor, err := golongpoll.NewFilePersistor(*persistFilename, 4096, 2)
	if err != nil {
		fmt.Printf("Failed to create file persistor, error: %v", err)
		return
	}

	manager, err := golongpoll.StartLongpoll(golongpoll.Options{
		LoggingEnabled: true,
		AddOn:          filePersistor,
	})
	if err != nil {
		log.Fatalf("Failed to create manager: %q", err)
	}

	http.HandleFunc("/filepersist", filePersistorExampleHomepage)
	http.HandleFunc("/js/client.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *staticClientJs)
	})
	http.HandleFunc("/filepersist/events", manager.SubscriptionHandler)
	http.HandleFunc("/filepersist/publish", manager.PublishHandler)
	fmt.Println("Serving webpage at http://127.0.0.1:8102/filepersist")
	http.ListenAndServe("127.0.0.1:8102", nil)
}

func getPublishHandler(manager *golongpoll.LongpollManager) func(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	// Creates closure that captures the LongpollManager
	return nil
}

func filePersistorExampleHomepage(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
