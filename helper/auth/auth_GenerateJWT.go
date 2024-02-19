package auth

import (
	"go-echo/model/parameter"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func (h *authHelperImpl) GenerateJWT(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// find user and check grant access
	tx := h.baseRepo.GetBegin()
	user, err := h.userRepo.FindUserRoleByUserId(tx, parameter.FindUserRoleByUserIdInput{
		Id: id,
	})

	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	now := time.Now()
	roles := []string{}
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
	}

	claims["iss"] = user.Id
	claims["iat"] = now.Unix()
	claims["sub"] = user.Username
	claims["exp"] = now.Add(24 * time.Hour).Unix()
	claims["usr"] = user.Nickname
	claims["roles"] = roles

	secret := []byte(viper.GetString("jwt.secret-key"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
