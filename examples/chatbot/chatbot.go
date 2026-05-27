// This example creates a dummy chatbot while demonstrating the following features:
// 1) Golang client used by the trivial chatbot
// 2) Javascript client used by UI.
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
	// if in ./examples/chatbot, use: -clientJs ../../js-client/client.js
	// If using go 1.16 or higher, can simply use the embed directive instead of having to do this here, but supporting older versions.
	staticClientJs := flag.String("clientJs", "./js-client/client.js", "where the static js-client/client.js is located relative to where this binary runs")
	flag.Parse()

	manager, err := golongpoll.StartLongpoll(golongpoll.Options{
		LoggingEnabled: true,
	})
	if err != nil {
		log.Fatalf("Failed to create manager: %q", err)
	}

	http.HandleFunc("/chatbot", chatBotExampleHomepage)
	http.HandleFunc("/js/client.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *staticClientJs)
	})
	http.HandleFunc("/chatbot/events", manager.SubscriptionHandler)
	http.HandleFunc("/chatbot/send", manager.PublishHandler)
	fmt.Println("Serving webpage at http://127.0.0.1:8101/chatbot")
	go beChatbot(manager)
	http.ListenAndServe("127.0.0.1:8101", nil)
}

func beChatbot(lpManager *golongpoll.LongpollManager) { _ = "STUB: not implemented"; return }

// chatbot will only listen to new events (on or after now())

// assuming all events are strings--which they are in this example

// Special responses for certain keywords, otherwise a random canned response.

// Note: would only get here if c.Stop() was called--which in this example it never is.
// You could put c.Stop() in one of the if blocks above to test this out.
// For example, if talking about cat's is the last straw, the bot could stop responding.

func chatBotExampleHomepage(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
