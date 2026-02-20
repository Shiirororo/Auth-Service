package response

const (
	ErrCodeSuccess  = 201 // Success
	ErrParamInvalid = 303 //Invalid
	ErrUnauthorized = 401
	ErrInvalidToken = 301
)

var msg = map[int]string{
	ErrCodeSuccess:  "Success",
	ErrParamInvalid: "Invalid Parameter",
	ErrUnauthorized: "Unauthorized",
	ErrInvalidToken: "Invalid Token",
}
