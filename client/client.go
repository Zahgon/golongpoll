// Package client provides a client for longpoll servers serving events using
// LongpollManager.SubscriptionHandler and (optionally) LongpollManager.PublishHandler.
package client

import (
	"net/http"
	"net/url"
	"time"

	"github.com/jcuga/golongpoll"
)

// ClientOptions for configuring client behavior.
type ClientOptions struct {
	// SubscribeUrl is the longpoll server's subscription handler's URL to hit when
	// polling for events. This is a URL pointing to where
	// LongpollManager.SubscriptionHandler is being served or wrapped by another
	// handler.
	SubscribeUrl url.URL
	// Category is the subscription category to poll.
	Category string
	// PublishUrl is the optional longpoll server's publish handler's URL to
	// publish event data to. Can be omitted if publishing not used or if
	// the LongpollManager.PublishHandler is not served/exposed.
	PublishUrl url.URL
	// PollTimeoutSeconds is the timeout arg the client sends to the longpoll
	// server, which dictates how long the server should keep the request idle
	// before issuing a timeout response when there hasn't been any event data
	// to respond with. Defaults to 45 seconds.
	// NOTE: if hitting a longpoll server behind a poxy/webserver, make sure
	// the PollTimeoutSeconds is less than that proxy's HTTP timeout setting!
	PollTimeoutSeconds uint
	// ReattemptWaitSeconds controls the amount of time the client waits to
	// reattempt polling the server after a failure response.
	// Defaults to 30 seconds.
	ReattemptWaitSeconds uint
	// SuccessWaitSeconds is how long to wait, in seconds, after receiving
	// events before requesting more events. Defaults to zero, no wait.
	SuccessWaitSeconds uint
	// HttpClient is an optional http.Client to use when polling the server.
	// Defaults to go's default http.Client.
	HttpClient *http.Client
	// BasicAuthUsername is an optional username to be used for basic HTTP
	// authentication when polling the server.
	BasicAuthUsername string
	// BasicAuthUsername is an optional password to be used for basic HTTP
	// authentication when polling the server.
	BasicAuthPassword string
	// Whether or not logging should be enabled
	LoggingEnabled bool
	// Optional callback for HTTP longpoll request failures.
	// The client will stop if the provided callback returns false.
	OnFailure func(err error) bool
	// ExtraHeaders has optional HTTP headers to include in lonpoll requests.
	// Useful if you wrap LongpollManager.SubscriptionHandler with additional
	// authentication or other logic that uses headers.
	ExtraHeaders []HeaderKeyValue
}

// HeaderKeyValue is a HTTP header key-value pair.
type HeaderKeyValue struct {
	Key   string
	Value string
}

// Client for polling a longpoll server.
type Client struct {
	// Flag that signals when the event polling goroutine should quit.
	// NOTE: using sys/atomic.AddUint64 on fields require them to be 64bit
	// aligned. The start of an allocated struct are guaranteed to be aligned,
	// so placing at start here fixes a cryptic panic I got after adding
	// a new field to this struct and then runID kept breaking.
	// see: https://stackoverflow.com/questions/28670232/atomic-addint64-causes-invalid-memory-address-or-nil-pointer-dereference
	runID uint64
	// options dictating client behavior
	options ClientOptions
	// Populated with received longpoll events.
	events chan *golongpoll.Event
	// flag whether or not Client.Start has been called--enforces use-only-once.
	started bool
	// flag whether or not Client.Stop has been called--enforces use-only-once.
	stopped bool
}

// NewClient creates a new Client configured with the given ClientOptions.
// Returns new client and nil error or nil client with a non-nil error.
func NewClient(opts ClientOptions) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

// Require both basic auth user/password, or neither

// Set defaults if missing/zero/nil

// Start begins Client's longpolling request-loop goroutine. The pollSince
// argument dictates the point in time we wish to see events since.
// Callers could pass time.Now() to only see new events, or a time in the
// past to start event consumption in the past. Use Client.Stop()
// to cease polling for new events. Read incoming events from the returned
// channel. The returned channel will be closed by client.Stop() or if
// ClientOptions.OnFailure returns false.
// Clients can only be started once and will panic otherwise.
// To resume long polling after stopping, create a new client via
// NewClient(ClientOptions).
func (c *Client) Start(pollSince time.Time) <-chan *golongpoll.Event {
	_ = "STUB: not implemented"
	return nil
}

// HTTP API takes milliseconds
// last received eventID, used to fix issue #19

// Don't bother sleeping if it's time to quit

// Stop sending events if time to quit. Checking now since this did
// an HTTP request since last time we checked.

// Check if it's time to quit before sending new events to
// channel if there's any chance we may have to wait
// for channel space to become available.
// That way, if clients call client.Stop() and don't read from
// the events channel anymore, we don't get stuck in this
// goroutine waiting for callers to read from the channel so we
// can send remaining data to it.

// Only push timestamp forward if its greater than the last we checked

// here we have no events, no error message, and a zero/no timestamp
// expect to get one of those.

// Stop signals to the client's event polling goroutine to stop.
// Upon receiving the stop signal, the client's goroutine will close the
// Events channel and stop running.  This function call does not block on the
// client's goroutine stopping.  Callers can wait for the Events channel
// returned by Client.Start() to be closed to know when that goroutine has
// finished executing. Note that said channel could also be closed due to
// ClientOptions.OnFailure returning false. Either way a closed channel
// indicates the client has stopped polling for event data.
// Clients can only be stopped once, multiple calls to this will panic.
// To resume long polling, create a new client via NewClient(ClientOptions).
func (c *Client) Stop() { _ = "STUB: not implemented"; return }

// Changing the runID will have any previous goroutine ignore any events it may receive

// Checks if time to quit and if so closes event channel.
func (c *Client) doQuit(runID uint64) bool { _ = "STUB: not implemented"; return false }

func (c *Client) getCommonLogFields(since int64, lastID string) string {
	_ = "STUB: not implemented"
	return ""
}

// Relevant parts of HTTP longpolling API's resposne data
type pollResponse struct {
	Events []*golongpoll.Event `json:"events"`
	// Set for timeout responses
	Timestamp int64 `json:"timestamp"`
	// API error responses could have an informative error here. Empty on success.
	ErrorMessage string `json:"error"`
}

// Call the longpoll server to get the events since a specific timestamp
func (c Client) fetchEvents(since int64, lastID string) (*pollResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Optional headers to include in request--can be used for extra authentication
// if the LongpollManager.SubscriptionHandler is wrapped with additional checks.

type publishResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// Publish will publish data on the given category to ClientOptions.PublishUrl.
// Returns nil on success, non-nil error on failure.
// ClientOptions.PublishUrl must be non-empty.
func (c Client) Publish(category string, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Optional headers to include in request--can be used for extra authentication
// if the LongpollManager.PublishHandler is wrapped with additional checks.
