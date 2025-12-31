package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type Manager interface {
	Verify(token string) (Payload, error)
	Generate(payload Payload, duration time.Duration) (string, error)
}

type Payload struct {
	jwt.StandardClaims
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

type implManager struct {
	secretKey string
}

func NewManager(secretKey string) Manager {
	return &implManager{
		secretKey: secretKey,
	}
}

// Verify verifies the token and returns the payload
func (m implManager) Verify(token string) (Payload, error) {
	if token == "" {
		return Payload{}, ErrInvalidToken
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			log.Printf("jwt.ParseWithClaims: %v", ErrInvalidToken)
			return nil, ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		log.Printf("jwt.ParseWithClaims: %v", err)
		return Payload{}, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		log.Printf("Parsing to Payload: %v", ErrInvalidToken)
		return Payload{}, ErrInvalidToken
	}

	return *payload, nil
}

// Generate generates a new JWT token with the given payload and duration
func (m implManager) Generate(payload Payload, duration time.Duration) (string, error) {
	// Set expiration time
	payload.ExpiresAt = time.Now().Add(duration).Unix()
	payload.IssuedAt = time.Now().Unix()

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// Sign token with secret key
	tokenString, err := token.SignedString([]byte(m.secretKey))
	if err != nil {
		log.Printf("jwt.Generate: %v", err)
		return "", ErrGenerateToken
	}

	return tokenString, nil
}
