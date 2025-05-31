package token

import (
	"errors"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/netesh5/go_ecommerce/internal/models"
)

type SignDetails struct {
	Email string
	Name  string
	ID    int
	jwt.StandardClaims
}

func getSecretKey() (string, error) {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		return "", jwt.NewValidationError("SECRET_KEY not set", jwt.ValidationErrorClaimsInvalid)
	}
	return key, nil
}

func TokenGenerator(email string, name string, id int) (token string, refreshToken string, err error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return "", "", err
	}
	claims := &SignDetails{
		Email: email,
		Name:  name,
		ID:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(time.Hour * 48).Unix(), // Token valid for 48 hours
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

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil

}

func TokenValidator(tokenString string) (claims *SignDetails, message string) {
	token, err := jwt.ParseWithClaims(tokenString, &SignDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		message = err.Error()
		return nil, message
	}
	claims, ok := token.Claims.(*SignDetails)
	if !ok || !token.Valid {
		return nil, "Invalid token"
	}

	if claims.ExpiresAt < jwt.TimeFunc().Unix() {
		return nil, "Token has expired"
	}

	return claims, message
}

//func UpdateAllTokens(token string, refreshToken string, userId int) (string, string, error) {
// var SECRET_KEY = os.Getenv("SECRET_KEY")
// if SECRET_KEY == "" {
// 	return "", "", jwt.NewValidationError("SECRET_KEY not set", jwt.ValidationErrorClaimsInvalid)
// }

// claims := &SignDetails{
// 	ID: userId,
// 	StandardClaims: jwt.StandardClaims{
// 		ExpiresAt: jwt.TimeFunc().Add(time.Hour * 48).Unix(),
// 		IssuedAt:  jwt.TimeFunc().Unix(),
// 		Issuer:    "go_ecommerce",
// 	},
// }

// refreshClaims := &SignDetails{
// 	ID: userId,
// 	StandardClaims: jwt.StandardClaims{
// 		ExpiresAt: jwt.TimeFunc().Add(time.Hour * 168).Unix(),
// 		IssuedAt:  jwt.TimeFunc().Unix(),
// 		Issuer:    "go_ecommerce",
// 	},
// }
// token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
// if err != nil {
// 	return "", "", err
// }

// refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
// if err != nil {
// 	return "", "", err
// }
// if token == "" || refreshToken == "" {
// 	return "", "", jwt.NewValidationError("Token generation failed", jwt.ValidationErrorClaimsInvalid)
// }

// return token, refreshToken, nil

//}

func TokenParser(tokenString string) (models.User, error) {
	claims := &SignDetails{}

	// Parse the token with claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return models.User{}, err
	}

	// Validate token and claims
	if !token.Valid {
		return models.User{}, errors.New("invalid token")
	}

	// Now claims is populated and can be used
	return models.User{
		ID:    claims.ID,
		Name:  claims.Name,
		Email: claims.Email,
	}, nil
}

// func ExtractTokenFromHeader(c echo.Context) (string, error) {
// 	authHeader := c.Request().Header.Get("Authorization")
// 	if authHeader == "" {
// 		return "", fmt.Errorf("no authorization header found")
// 	}
// 	splitToken := strings.Split(authHeader, " ")
// 	if len(splitToken) != 2 || strings.ToLower(splitToken[0]) != "bearer" {
// 		return "", fmt.Errorf("no authorization header found")
// 	}

// 	return splitToken[1], nil

// }
