package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractBearerToken(c *gin.Context, requiredHeader string) (string, error) {
	Header := c.GetHeader(requiredHeader)
	if Header == "" {
		return "", errors.New("No" + requiredHeader)
	}

	parts := strings.SplitN(Header, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid authorization format")
	}

	return parts[1], nil
}
func abortUnauthorized(c *gin.Context, msg string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"error": msg,
	})
}

// func emailValidationMiddleware(c *gin.Context, email string) error  {
// 	emailString, err := extractBearerToken(email, "Email")
// 	if err != nil {
// 		return errors.New("No Email")
// 	}
// 	if email
// }
