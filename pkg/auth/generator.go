package auth

import (
	logger "bitbucket.org/gomatrix/customlogger"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type PasswordGenerator interface {
	MakeRandomString(length int) string
	MakeHash(key string) (string, error)
	CompareHashAndPassword(hashedKey string, key string) error
}

type Generator struct {
	letterRunes []rune
}

func NewGenerator() *Generator {
	return &Generator{letterRunes: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")}
}

func (g *Generator) MakeRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = g.letterRunes[rand.Intn(len(g.letterRunes))]
	}
	return string(b)
}

func (g *Generator) MakeHash(key string) (string, error) {
	hashedKey, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	if err != nil {
		logger.Errorf("Error make hash: %v", err)
		return "", err
	}
	return string(hashedKey), nil
}

func (g *Generator) CompareHashAndPassword(hashedKey string, key string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedKey), []byte(key))
}
