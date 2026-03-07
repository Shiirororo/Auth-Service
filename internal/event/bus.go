package event

type Bus struct {
	queue chan Event
}

func NewBus(size int) *Bus {
	return &Bus{
		queue: make(chan Event, size),
	}
}

func (b *Bus) Publish(e Event) {
	b.queue <- e
}

func (b *Bus) Queue() <-chan Event {
	return b.queue
}

// type LoginEvent struct {
// 	UserID    string
// 	SessionID string
// 	IPAddress string
// 	Timestamp time.Time
// }
// type OrderEvent struct {
// }

// type AuditEvent struct {
// }
