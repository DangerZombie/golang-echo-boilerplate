package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(1 * time.Minute).Unix()
	claims["authorized"] = true
	claims["user"] = username

	secret := []byte(viper.GetString("jwt.secret-key"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(headers http.Header) (string, error) {
	tokenString := strings.Split(headers["Authorization"][0], " ")[1]

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Here you need to specify the function that will be used to verify the key.
		// In this case, we are using a shared secret key.
		return []byte(viper.GetString("jwt.secret-key")), nil
	})

	// Verify the token
	if err != nil {
		return "", err
	}
	if !token.Valid {
		err = fmt.Errorf("error: %s", "token invalid")
		return "", err
	}

	// Access the claims
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("error: %s", "token invalid")
		return "", err
	}

	return "", nil
}
