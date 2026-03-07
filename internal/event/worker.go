package event

import (
	"context"
)

type Dispatcher struct {
	eventQueue chan Event

	loginQueue chan Event
	orderQueue chan Event
	auditQueue chan Event
}

func NewDispatcher(eventQueue chan Event) *Dispatcher {
	return &Dispatcher{
		eventQueue: eventQueue,

		loginQueue: make(chan Event, 100),
		orderQueue: make(chan Event, 100),
		auditQueue: make(chan Event, 100),
	}
}
func (d *Dispatcher) Start(ctx context.Context) {
	for {
		select {
		case event := <-d.eventQueue:
			d.Dispatch(event)
		case <-ctx.Done():
			return
		}

	}
}
func (d *Dispatcher) Dispatch(event Event) {
	switch event.Type {

	case LoginEvent:
		d.loginQueue <- event

	case OrderEvent:
		d.orderQueue <- event

	case AuditEvent:
		d.auditQueue <- event
	}
}
