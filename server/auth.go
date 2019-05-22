package server

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const tokenExpireTime = time.Hour * 12

type TokenClaims struct {
	jwt.StandardClaims
	Id int `json:"id"`
}

// comparePassword returns nil if the given plain password matches the secret
// and returns an error otherwise.
func comparePassword(pw, secret string) error {
	return bcrypt.CompareHashAndPassword([]byte(secret), []byte(pw))
}

// hashPassword creates a secret from the given password using bcrypt hashing.
// Returns an error if password hashing fails.
func hashPassword(pw string) (string, error) {
	secretBytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	return string(secretBytes), err
}

// newSignedJWT returns a signed JWT containing the given ID.
func newSignedJWT(id int, issuer string, key []byte) (string, error) {
	sClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenExpireTime).Unix(),
		Issuer:    issuer,
		IssuedAt:  time.Now().Unix(),
	}
	claims := &TokenClaims{
		StandardClaims: sClaims,
		Id:             id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

// verifyJWT verifies the given JWT with the given key.
func verifyJWT(encodedToken string, key []byte) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
}
