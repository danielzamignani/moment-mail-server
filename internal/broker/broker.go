package broker

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

type Event struct {
	Type string      `json:"Type"`
	Data interface{} `json:"data,omitempty"`
}

type EventBroker struct {
	mutex   sync.Mutex
	clients map[uuid.UUID]chan Event
}

func NewEventBroker() *EventBroker {
	return &EventBroker{
		clients: make(map[uuid.UUID]chan Event),
	}
}

func (b *EventBroker) Subscribe(inboxId uuid.UUID) <-chan Event {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	ch := make(chan Event)
	b.clients[inboxId] = ch
	log.Printf("New channel subscriber: %s", inboxId)

	return ch
}

func (b *EventBroker) Unsubscribe(inboxId uuid.UUID) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if ch, ok := b.clients[inboxId]; ok {
		close(ch)
		delete(b.clients, inboxId)
		log.Printf("Client %s removed", inboxId)
	}
}

func (b *EventBroker) Publish(inboxId uuid.UUID, event Event) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if ch, ok := b.clients[inboxId]; ok {
		ch <- event
	}
}
