package auth

import (
	"go-echo/model/parameter"
	"net/http"
)

type AuthHelper interface {
	GenerateJWT(id string) (string, error)
	VerifyJWT(headers http.Header) (output parameter.JwtClaims, err error)
}
