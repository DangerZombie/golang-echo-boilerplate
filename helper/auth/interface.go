package auth

import "net/http"

type AuthHelper interface {
	GenerateJWT(username string) (string, error)
	VerifyJWT(headers http.Header) (string, error)
}
