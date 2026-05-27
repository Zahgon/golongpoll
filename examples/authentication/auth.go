// Provides example of how one can wrap SubscriptionHandler and PublishHandler
// to provide authentication. One could do other things like limit what categories
// can be published to, or what sort of data can be published and by whom.
// This also provides examples of how the javascript and golang clients
// can provide http basic auth or other header data for authentication.
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
	flag.Parse()

	manager, err := golongpoll.StartLongpoll(golongpoll.Options{
		LoggingEnabled: true,
	})
	if err != nil {
		log.Fatalf("Failed to create manager: %q", err)
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/js/client.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *staticClientJs)
	})
	http.HandleFunc("/basic-events", withBasicAuth(manager.SubscriptionHandler))
	http.HandleFunc("/basic-publish", withBasicAuth(manager.PublishHandler))
	http.HandleFunc("/header-events", withMockHeaderAuth(manager.SubscriptionHandler))
	http.HandleFunc("/header-publish", withMockHeaderAuth(manager.PublishHandler))
	fmt.Printf("Serving webpage at http://%s\n", *serveAddr)

	go clientWithBasicAuth(*serveAddr)
	go clientWithMockHeaderAuth(*serveAddr)

	http.ListenAndServe(*serveAddr, nil)
}

func withBasicAuth(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: could add other checks/logic here as desired.

// passed auth check, execute oiginal handler

func withMockHeaderAuth(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: could add other checks/logic here as desired.

// passed auth check, execute oiginal handler

// Demonstrates how a client can be configured to use http basic auth.
// This will subscribe to auth-protected events and then publish to that
// same auth-protected category acknowledging that it was able to see the event.
func clientWithBasicAuth(serveAddr string) { _ = "STUB: not implemented"; return }

// Demonstrates how a client can be configured to use header based auth.
// This will subscribe to auth-protected events and then publish to that
// same auth-protected category acknowledging that it was able to see the event.
func clientWithMockHeaderAuth(serveAddr string) { _ = "STUB: not implemented"; return }

func homePage(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }
