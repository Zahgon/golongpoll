package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/jcuga/golongpoll"
)

func main() {
	staticClientJs := flag.String("clientJs", "./js-client/client.js", "where the static js-client/client.js is located relative to where this binary runs")
	serveAddr := flag.String("serve", "127.0.0.1:8080", "Address to serve HTTP on.")
	category := flag.String("category", "testing", "Longpoll category to display on webpage.")
	flag.Parse()

	manager, err := golongpoll.StartLongpoll(golongpoll.Options{
		LoggingEnabled: true,
	})
	if err != nil {
		log.Fatalf("Failed to create manager: %q", err)
	}

	http.HandleFunc("/", getStressorHomepage(*category))
	http.HandleFunc("/js/client.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *staticClientJs)
	})
	http.HandleFunc("/events", manager.SubscriptionHandler)
	http.HandleFunc("/publish", manager.PublishHandler)
	fmt.Printf("Serving webpage at http://%s\n", *serveAddr)
	http.ListenAndServe(*serveAddr, nil)
}

func getStressorHomepage(category string) func(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return nil
}
