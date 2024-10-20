package http

import (
	"go-echo/helper/auth"
)

type httpImpl struct {
	authHelper auth.AuthHelper
}

func NewHttp(ah auth.AuthHelper) Http {
	return &httpImpl{ah}
}
