package broker

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

type Event struct {
	Type string      `json:"type"`
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

func (eventBroker *EventBroker) Subscribe(inboxId uuid.UUID) <-chan Event {
	eventBroker.mutex.Lock()
	defer eventBroker.mutex.Unlock()

	ch := make(chan Event, 1)
	eventBroker.clients[inboxId] = ch
	log.Printf("New client subscribed inboxID: %s", inboxId)
	return ch
}

func (eventBroker *EventBroker) Unsubscribe(inboxId uuid.UUID) {
	eventBroker.mutex.Lock()
	defer eventBroker.mutex.Unlock()

	if ch, ok := eventBroker.clients[inboxId]; ok {
		close(ch)
		delete(eventBroker.clients, inboxId)
		log.Printf("Client inboxID %s unsubscribed", inboxId)
	}
}

func (eventBroker *EventBroker) Publish(inboxId uuid.UUID, event Event) {
	eventBroker.mutex.Lock()
	defer eventBroker.mutex.Unlock()

	if ch, ok := eventBroker.clients[inboxId]; ok {
		select {
		case ch <- event:
		default:
			log.Printf("Message for inboxID %s discarted", inboxId)
		}
	}
}
