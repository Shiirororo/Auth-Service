package event

type EventType string
type Event struct {
	Type    EventType
	Payload any
}

type RegisterSuccessPayload struct {
	Username string
	Email    string
	Password string
}

type LoginPayload struct {
	UserID []byte
}
type PaymentSuccess struct {
}

const (
	LoginEvent           EventType = "login"
	RegisterSuccessEvent EventType = "register"
	OrderEvent           EventType = "order"
	AuditEvent           EventType = "audit"
)
