package token

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type SignDetails struct {
	Email string
	Name  string
	ID    int
	jwt.StandardClaims
}

func TokenGenerator(email string, name string, id int) (token string, refreshToken string, err error) {
	claims := &SignDetails{
		Email: email,
		Name:  name,
		ID:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(time.Hour * 48).Unix(), // Token valid for 24 hours
			IssuedAt:  jwt.TimeFunc().Unix(),                     // Token issued at current time
			Issuer:    "go_ecommerce",
		},
	}

	refreshClaims := &SignDetails{
		Email: email,
		Name:  name,
		ID:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(time.Hour * 168).Unix(), // Refresh token valid for 72 hours
			IssuedAt:  jwt.TimeFunc().Unix(),                      // Refresh token issued at current time
			Issuer:    "go_ecommerce",
		},
	}

	var SECRET_KEY = os.Getenv("SECRET_KEY")
	if SECRET_KEY == "" {
		return "", "", jwt.NewValidationError("SECRET_KEY not set", jwt.ValidationErrorClaimsInvalid)
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil

}

func TokenValidator() {

}

func UpdateAllTokens() {

}
