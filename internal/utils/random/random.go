package random

import (
	"math/rand"
	"time"
)

func GenerateOPT6Digit() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := 100000 + rng.Intn(900000)
	return otp
}

func HashEmail(email string) string {
	return " "
}
