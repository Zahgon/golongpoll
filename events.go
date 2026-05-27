package golongpoll

import (
	"container/list"
	"time"

	"github.com/gofrs/uuid"
)

// Event is a longpoll event.  This type has a Timestamp as milliseconds since
// epoch (UTC), a string category, and an arbitrary Data payload.
// The category is the subscription category/topic that clients can listen for
// via longpolling.  The Data payload can be anything that is JSON serializable
// via the encoding/json library's json.Marshal function.
type Event struct {
	// Timestamp is milliseconds since epoch to match javascrits Date.getTime().
	// This is the timestamp when the event was published.
	Timestamp int64 `json:"timestamp"`
	// Category this event belongs to. Clients subscribe to a given category.
	Category string `json:"category"`
	// Event data payload.
	// NOTE: Data can be anything that is able to passed to json.Marshal()
	Data interface{} `json:"data"`
	// Event ID, used in conjunction with Timestamp to get a complete timeline
	// of event data as there could be more than one event with the same timestamp.
	ID uuid.UUID `json:"id"`
}

// eventResponse is the json response that carries longpoll events.
type eventResponse struct {
	Events []*Event `json:"events"`
}

// eventBuffer is a buffer of Events that adds new events to the front/root and
// and old events are removed from the back/tail when the buffer reaches it's
// maximum capacity.
// NOTE: this add-new-to-front/remove-old-from-back behavior is fairly
// efficient since it is implemented as a ring with root.prev being the tail.
// Unlike an array, we don't have to shift every element when something gets
// added to the front, and because our root has a root.prev reference, we can
// quickly jump from the root to the tail instead of having to follow every
// node's node.next field to finally reach the end.
// For more details on our list's implementation, see:
// https://golang.org/src/container/list/list.go
type eventBuffer struct {
	*list.List
	MaxBufferSize int
	// keeping track of this allows for more efficient event TTL expiration purges:
	// time in milliseconds since epoch since thats what Event types use
	// for Timestamps
	oldestEventTime int64
}

func newEvent(category string, data interface{}) *Event { _ = "STUB: not implemented"; return nil }

func newEventWithTime(t time.Time, category string, data interface{}) *Event {
	_ = "STUB: not implemented"
	return nil
}

// QueueEvent adds a new longpoll Event to the front of our buffer and removes
// the oldest event from the back of the buffer if we're already at maximum
// capacity.
func (eb *eventBuffer) QueueEvent(event *Event) error { _ = "STUB: not implemented"; return nil }

// Cull our buffer if we're at max capacity

// Add event to front of our list

// Update oldestEventTime with the time of our least recent event (at back)
// keeping track of this allows for more efficient event TTL expiration purges

// GetEventsSnce will return all of the Events in our buffer that occurred after
// the given input time (since).  Returns an error value if there are any
// objects that aren't an Event type in the buffer.  (which would be weird...)
// Optionally removes returned events from the eventBuffer if told to do so by
// deleteFetchedEvents argument.
func (eb *eventBuffer) GetEventsSince(since time.Time,
	deleteFetchedEvents bool, lastEventUUID *uuid.UUID) ([]*Event, error) {
	_ = "STUB: not implemented"
	return nil,

		// NOTE: events are bufferd with the most recent event at the front.
		// So we want to start our search at the front of the buffer and stop
		// searching once we've reached events that are older than the 'since'
		// argument.  But we want to return the subset of events in chronological
		// order, which is least recent in front.  So do our search from the
		// start so we can cut out early, but then iterate back from our last
		// item we want to return as a result.  Doing this avoids having to capture
		// results and then create another copy of the results but in reverse
		// order.
		nil
}

// Search forward until we reach events that are too old

// is event time after 'since' time arg? convert 'since' to epoch ms

// Event time is same as last seen event time, but ID is different, so this is more recent
// than last seen event since we see events most-recent-first

// we've reached items that are too old, they occurred before or on
// 'since' so we don't care about anything after this point.

// Now accumulate results in the correct chronological order starting from
// our oldest, valid Event that occurrs after 'since'

// Tracked outside of loop conditional to allow delete while iterating:

// we already know this event is after 'since'

// Advance iteration before List.Remove() invalidates element.prev

// Now safely remove from list if told to do so:

// element.Prev() now == nil

func (eb *eventBuffer) DeleteEventsOlderThan(olderThanTimeMs int64) error {
	_ = "STUB: not implemented"
	return nil
}

// Either no events or the the oldest event is more recent than
// olderThanTimeMs, so nothing  could possibly be expired.
// skip searching list

// Search list in reverse (starting from the back) removing expired events
// and updating eb.oldestEventTime as we remove stale events.
// NOTE: we iterate over list in reverse since oldest elements are at
// the back, newest up front.

// Advance iteration before List.Remove() invalidates element.prev

// Update oldestEventTime to the current event's Timestamp

// Now able to safely remove from list event is too old:

// element.Prev() now == nil

// element is too new, stop checking since events are only going to
// get even more recent as we get closer to the front of the list
