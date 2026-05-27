package golongpoll

// This priority queue manages eventBuffers that expire after a certain
// period of inactivity (no new events).
type expiringBuffer struct {
	// references an eventBuffer
	eventBufferPtr *eventBuffer
	// The subscription category for the given event buffer
	// This is needed so we can clean up our category-to-Item map
	// by doing a simple key lookup and removing the eventBuffer ref
	category string
	// The priority of the item in the queue.
	// For our purposes, this is milliseconds since epoch
	priority int64
	// index is needed by update and is maintained by heap.Interface
	// The index of this item in the heap.
	index int
}

// A Priority Queue (min heap) implemented with go's heap container.
// Adapted from go's example at: https://golang.org/pkg/container/heap/
//
// This priorityQueue is used to keep track of eventBuffer objects in order of
// oldest last-event-timestamp so that we can more efficiently purge buffers
// that have expired events.
//
// The priority here will be a timestamp in milliseconds since epoch (int64)
// with lower values (older timestamps) being at the top of the heap/queue and
// higher values (more recent timestamps) being further down.
// So this is a Min Heap.
//
// A priorityQueue implements heap.Interface and holds Items.
type priorityQueue []*expiringBuffer

func (pq priorityQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (pq priorityQueue) Less(i, j int) bool {
	_ = "STUB: not implemented"
	// We want Pop to give us the lowest priority, so less uses < here:
	return false
}

func (pq priorityQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (pq *priorityQueue) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (pq *priorityQueue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// for safety

// update modifies the priority of an item and updates the heap accordingly
func (pq *priorityQueue) updatePriority(item *expiringBuffer, priority int64) {
	_ = "STUB: not implemented"
	return

	// NOTE: fix is a slightly more efficient version of calling Remove() and
	// then Push()
}

// get the priority of the heap's top item.
func (pq *priorityQueue) peekTopPriority() (int64, error) { _ = "STUB: not implemented"; return 0, nil }
