package test

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

var myKey = []byte("toj-key")

// Generate token
func TestGenerateToken(t *testing.T) {
	UserClaim := &UserClaims{
		Identity:       "user_1",
		Name:           "Get",
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		t.Fatal(err)
	}
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InVzZXJfMSIsIm5hbWUiOiJHZXQifQ.FIla3bg5t0vGA83tdKFQOfgLTI0Lc9MUzd0C4_hKln8
	fmt.Println(tokenString)
}

// Parse token
func TestAnalyseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZGVudGl0eSI6InVzZXJfMSIsIm5hbWUiOiJHZXQifQ.FIla3bg5t0vGA83tdKFQOfgLTI0Lc9MUzd0C4_hKln8"
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if claims.Valid {
		fmt.Println(userClaim)
	}
}
