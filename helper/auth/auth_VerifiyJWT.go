package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

// TODO: need to fix verifying JWT
func (h *authHelperImpl) VerifyJWT(headers http.Header) (string, error) {
	if headers.Get("Authorization") == "" {
		return "", errors.New("token is null, need valid token")
	}

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
		err = fmt.Errorf("errors: %s", "token invalid")
		return "", err
	}

	// Access the claims
	_, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("errors: %s", "token invalid")
		return "", err
	}

	return "", nil
}
