package golongpoll

import (
	"time"
)

// FilePersistorAddOn implements the AddOn interface to provide
// file persistence for longpoll events. Use NewFilePersistor(string, int, int)
// to create a configured FilePersistorAddOn.
//
// NOTE: uses bufio.NewScanner which has a default max scanner buffer size of
// 64kb (65536). So any event whose JSON is that large will fail to be
// read back in completely. These events, and any other events whose JSON
// fails to unmarshal will be skipped.
type FilePersistorAddOn struct {
	// Filename to use for storing events. The file will be created if it
	// does not already exist.
	filename string
	// How large the underlying bufio.Writer's buffer is.
	writeBufferSize int
	// How often to flush buffered data to disk.
	writeFlushPeriodSeconds int
	// Channel for incoming published events.
	// To avoid blocking LongpollManager when it calls OnPublish(), we send
	// events to channel for processing in a separate goroutine.
	publishedEvents chan *Event
	// Channel used to signal when done flushing to disk during OnShutdown().
	shutdownDone chan bool
}

// NewFilePersistor creates a new FilePersistorAddOn with the provided options.
// filename is the file to use for storing event data.
// writeBufferSize is how large a buffer is used to buffer output before
// writing to file. This is simply the underlying bufio.Writer's buffer size.
// writeFlushPeriodSeconds is how often to flush buffer to disk even when
// the output buffer is not filled completely. This helps avoid data loss
// in the event of a program crash.
// Returns a new FilePersistorAddOn and nil error on success, or a nil
// addon and a non-nil error explaining what went wrong creating a new addon.
func NewFilePersistor(filename string, writeBufferSize int, writeFlushPeriodSeconds int) (*FilePersistorAddOn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure we can use filename

// OnLongpollStart returns a channel of events that is populated
// via a separate goroutine so that the calling LongpollManager can
// begin consuming events while we're still reading more events in
// from file. Note that the goroutine that adds events to this returned
// channel will call close() on the channel when it is out of
// initial events. The LongpollManager will wait for more events
// until the channel is closed.
func (fp *FilePersistorAddOn) OnLongpollStart() <-chan *Event {
	_ = "STUB: not implemented"
	// return a channel to send initial events to.
	return nil
}

// populate input events channel in own goroutine, which will
// call close(ch) once it sends all events.

// launch the FilePersistorAddOn's run goroutine. This will read from
// a different channel to handle OnPublish() events.

// OnPublish will write new events to file.
// Events are sent via channel to a separate goroutine so that the calling
// LongpollManager does not block on the file writing.
func (fp *FilePersistorAddOn) OnPublish(event *Event) { _ = "STUB: not implemented"; return }

// OnShutdown will signal the run goroutine to flush any remaining event data
// to disk and wait for it to complete.
func (fp *FilePersistorAddOn) OnShutdown() {
	_ = "STUB: not implemented"
	// Signal to stop working flush to disk.
	return
}

// wait for signal that flush to disk has finished (channel will be closed)

// Reads previously stored events from file and sends them to the channel
// returned by OnLongpollStart().
// NOTE: uses bufio.NewScanner which has a default max scanner buffer size of
// 64kb (65536). So any event whose JSON is that large will fail to be
// read back in completely. These events, and any other events whose JSON
// fails to unmarshal will be skipped.
func (fp *FilePersistorAddOn) getOnStartInputEvents(ch chan *Event) {
	_ = "STUB: not implemented"
	return
}

// skip any blank lines as we prepend newline before writing event json.
// NOTE: doing prepend instead of append so if an event getting written
// out is stopped prematurely due to an ungraceful shutdown, we will
// start a subsequent write on it's own line.

// NOTE: important to close--or calling LongpollManager will hang waiting
// for more channel data.

// FilePersistorAddOn's run goroutine that reads OnPublish()'s events from a
// channel. This allows OnPublish to return without handling the file writing
// directly, thus LongpollManager is not blocking on file writes.
// This will also flush to disk any buffered data at the configured time
// interval. When OnShtudown() is called, this will flush any remaining data
// and stop.
func (fp *FilePersistorAddOn) run() { _ = "STUB: not implemented"; return }

// NOTE: O_APPEND is critical here as I was an idiot and could not figure
// out why I was getting weird file data wrapping around across restarts!
// Without this, we'll start at the beginning of the file instead of adding
// to the end. *smacks palm on forehead*

// issue#30: use time.NewTimer with Reset to avoid creating a new timer channel per loop iteration
// that lingers (isn't closed) until duration exceeded.

// channel closed, flush any buffered data to disk and stop

// signal finished shutting down, OnShutdown is blocking waiting for channel action

// NOTE: adding newline before instead of after in case last event wasn't fully written
// out to completion

func isTimeToFlush(lastFlushTime time.Time, writeFlushPeriodSeconds int) bool {
	_ = "STUB: not implemented"
	return false
}
