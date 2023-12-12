package queue

import (
	"sync"
)

type QueueEntry struct {
	Prompt string
}

type Queue struct {
	Entries []QueueEntry
	lenght  uint
	mu      sync.Mutex
}

var Q = Queue{}

func (queue *Queue) Push(entry QueueEntry) {
	queue.mu.Lock()
	defer queue.mu.Unlock()

	queue.Entries = append(queue.Entries, entry)
	queue.lenght++
}

func (queue *Queue) Pop() *QueueEntry {
    queue.mu.Lock()
    defer queue.mu.Unlock()

    if queue.lenght == 0 {
        return nil
    }

    entry := queue.Entries[0]
    queue.Entries = queue.Entries[1:]
    queue.lenght--

    return &entry
}

func (queue *Queue) Peek() *QueueEntry {
    queue.mu.Lock()
    defer queue.mu.Unlock()

    if queue.lenght == 0 {
        return nil
    }

    entry := queue.Entries[0]

    return &entry
}

func (queue *Queue) Lenght() uint {
    queue.mu.Lock()
    defer queue.mu.Unlock()

    return queue.lenght
}
