package event

import (
	"context"
	"sync"
)

// EventHandler is an interface that any worker or service can implement to handle a specific event.
type EventHandler interface {
	Handle(ctx context.Context, event Event) error
}

type Dispatcher struct {
	eventQueue chan Event
	handlers   map[EventType][]EventHandler
	mu         sync.RWMutex
}

func NewDispatcher(eventQueue chan Event) *Dispatcher {
	return &Dispatcher{
		eventQueue: eventQueue,
		handlers:   make(map[EventType][]EventHandler),
	}
}

// Register allows new workers or components to subscribe to specific events dynamically.
func (d *Dispatcher) Register(eventType EventType, handler EventHandler) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[eventType] = append(d.handlers[eventType], handler)
}

func (d *Dispatcher) Start(ctx context.Context) {
	for {
		select {
		case event := <-d.eventQueue:
			d.Dispatch(ctx, event)
		case <-ctx.Done():
			return
		}
	}
}

func (d *Dispatcher) Dispatch(ctx context.Context, event Event) {
	d.mu.RLock()
	handlers, exists := d.handlers[event.Type]
	d.mu.RUnlock()

	if !exists {
		return
	}

	// Detach context cancellation since this is an async background operation
	detachedCtx := context.WithoutCancel(ctx)

	// Fan out the event to all registered handlers for this event type
	for _, handler := range handlers {
		// We execute the handler in a new goroutine to avoid blocking the dispatcher loop.
		// The EventHandler implementation itself should handle its own worker pooling if needed.
		go handler.Handle(detachedCtx, event)
	}
}
