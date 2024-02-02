package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

// TODO: need to fix generating JWT with valid step
func (h *authHelperImpl) GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = username

	secret := []byte(viper.GetString("jwt.secret-key"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
