// Super simple chat server with some pre-defined rooms and login-less posting
// using display names. No attempt is made at security.
// This shows how one could use categories and the longpoll pub-sub to make
// a chat server with individual rooms/topics.
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jcuga/golongpoll"
)

func main() {
	listenAddress := flag.String("serve", "127.0.0.1:8080", "address:port to serve.")
	staticClientJs := flag.String("clientJs", "./js-client/client.js", "where the static js-client/client.js is located relative to where this binary runs")
	flag.Parse()

	http.HandleFunc("/js/client.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *staticClientJs)
	})

	// NOTE: to have chats persist across program runs, use Options.AddOn: FilePersistorAddOn.
	manager, err := golongpoll.StartLongpoll(golongpoll.Options{
		// How many chats per topic to hang on to:
		MaxEventBufferSize: 1000,
		LoggingEnabled:     true,
	})
	if err != nil {
		log.Fatalf("Failed to create chat longpoll manager: %q\n", err)
	}

	http.HandleFunc("/", indexPage)
	http.HandleFunc("/topic/news", topicPage("news"))
	http.HandleFunc("/topic/sports", topicPage("sports"))
	http.HandleFunc("/topic/politics", topicPage("politics"))
	http.HandleFunc("/topic/humor", topicPage("humor"))
	// NOTE: using the plain publish handler.  If one wanted to add
	// additional behavior like escaping or validating data, one could make a
	// http handler function that has a closure capturing the longpoll manager
	// and then call manager.Publish directly.
	http.HandleFunc("/post", manager.PublishHandler)
	http.HandleFunc("/events", manager.SubscriptionHandler)
	log.Printf("Launching chat server on http://%s\n", *listenAddress)
	http.ListenAndServe(*listenAddress, nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func topicPage(topic string) func(http.ResponseWriter, *http.Request) {
	_ = "STUB: not implemented"
	return nil
}
