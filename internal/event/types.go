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

// CheckEmailPayload is dispatched by the User service to ask the Auth service
// whether an email already exists. The caller blocks on ReplyCh.
type CheckEmailPayload struct {
	Email   string
	ReplyCh chan bool // true = email exists
}

// CheckUsernamePayload is dispatched by the User service to ask the Auth service
// whether a username already exists. The caller blocks on ReplyCh.
type CheckUsernamePayload struct {
	Username string
	ReplyCh  chan bool // true = username exists
}

type PaymentSuccess struct {
}

const (
	LoginEvent             EventType = "login"
	RegisterSuccessEvent   EventType = "register"
	CheckEmailEvent        EventType = "check_email"
	CheckUsernameEvent     EventType = "check_username"
	OrderEvent             EventType = "order"
	AuditEvent             EventType = "audit"
)
