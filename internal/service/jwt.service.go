package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtService struct {
	secret     string
	accessTTL  time.Duration
	refreshTTL time.Duration
}
type JWTService interface {
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	VerifyJWT(ctx context.Context, tokenString string) (*Claims, error)
	ParseRefreshToken(tokenString string) (*Claims, error)
}

func NewJWTService(secret string) *jwtService {
	return &jwtService{
		secret:     secret,
		accessTTL:  15 * time.Minute,
		refreshTTL: 7 * 24 * time.Hour,
	}
}

type Claims struct {
	UserID    string `json:"user_id"`
	TokenType string `json:"type"`
	jwt.RegisteredClaims
}

func (s *jwtService) GenerateAccessToken(userID string) (string, error) {
	if len(s.secret) == 0 {
		return "", errors.New("JWT_SECRET is not set in environment variables")
	}

	now := time.Now()
	claims := Claims{
		UserID:    userID,
		TokenType: "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(), // --> JIT
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "authorizor_api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *jwtService) GenerateRefreshToken(userID string) (string, error) {
	if len(s.secret) == 0 {
		return "", errors.New("JWT_SECRET is not set in environment variables")
	}

	now := time.Now()
	claims := Claims{
		UserID:    userID,
		TokenType: "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.NewString(),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.refreshTTL)), // 7 days
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Issuer:    "backend_api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

// VerifyJWT parses and validates the token, checking against the blacklist
func (s *jwtService) VerifyJWT(ctx context.Context, tokenString string) (*Claims, error) {
	return s.parseToken(tokenString)
}

func (s *jwtService) ParseRefreshToken(tokenString string) (*Claims, error) {
	claims, err := s.parseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != "refresh" {
		return nil, errors.New("invalid token type")
	}

	return claims, nil
}

func (s *jwtService) parseToken(tokenString string) (*Claims, error) {
	if len(s.secret) == 0 {
		return nil, errors.New("configuration error: JWT_SECRET is not set")
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return s.secret, nil
		},
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token has expired")
		}
		return nil, errors.New("invalid token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token contents")
	}

	return claims, nil
}
