package event

type EventType string
type Event struct {
	Type    EventType
	Payload any
}

const (
	LoginEvent EventType = "login"
	OrderEvent EventType = "order"
	AuditEvent EventType = "audit"
)
