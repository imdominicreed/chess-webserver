package broadcaster

import (
	"log"
	"sync"
)

type Broadcaster struct {
	channels map[int]chan string
	mu       sync.Mutex
	id       int
}

func (b *Broadcaster) AddChannel() int {
	b.mu.Lock()
	b.channels[b.id] = make(chan string)
	b.id += 1
	b.mu.Unlock()
	return b.id
}

func (b *Broadcaster) Broadcast(msg string) {
	b.mu.Lock()
	for id, ch := range b.channels {
		log.Printf("Broadcasting: %d", id)
		ch <- msg
	}
	b.mu.Unlock()
}

func (b *Broadcaster) DeleteChannel(id int) {
	b.mu.Lock()
	delete(b.channels, id)
	b.mu.Unlock()
}

func (b *Broadcaster) AwaitMessage(id int) string {

	return <-b.channels[id]
}

func NewBroadcaster() Broadcaster {
	return Broadcaster{channels: make(map[int]chan string)}
}
