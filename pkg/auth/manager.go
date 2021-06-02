package auth

import (
	logger "bitbucket.org/gomatrix/customlogger"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/maksymchuk-mm/crm/pkg/errors"
	"math/rand"
	"time"
)

type TokenManager interface {
	NewJWT(userPublicID uuid.UUID, ttl time.Duration) (string, errors.CrmError)
	Parse(accessToken string) (*Claims, errors.CrmError)
	NewRefreshToken() (string, errors.CrmError)
}

type Manager struct {
	signingKey string
}

type Claims struct {
	jwt.StandardClaims
	UserID uuid.UUID `json:"UserID"`
}

func NewManager(signingKey string) (*Manager, errors.CrmError) {
	if signingKey == "" {
		return nil, errors.Errors.EmptySigningKey
	}
	return &Manager{signingKey: signingKey}, nil
}

func (m *Manager) NewJWT(userPublicID uuid.UUID, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			Subject:   userPublicID.String(),
		},
		UserID: userPublicID,
	})
	tokenString, err := token.SignedString([]byte(m.signingKey))
	if err != nil {
		logger.Errorf(err.Error())
		return "", errors.Errors.TokenGenerationError
	}
	return tokenString, nil
}

func (m *Manager) Parse(accessToken string) (*Claims, errors.CrmError) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(m.signingKey), nil
	})
	if err != nil {
		logger.Errorf(err.Error())
		return nil, errors.Errors.TokenParseError
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.Errors.NoUserInClaims
	}

	return claims, nil
}

func (m *Manager) NewRefreshToken() (string, errors.CrmError) {
	b := make([]byte, 32)
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	if _, err := r.Read(b); err != nil {
		logger.Errorf(err.Error())
		return "", errors.Errors.TokenGenerationError
	}
	return fmt.Sprintf("%x", b), nil
}
