package endpoint

import (
	"go-echo/helper/auth"
)

type endpointImpl struct {
	authHelper auth.AuthHelper
}

func NewEndpoint(ah auth.AuthHelper) Endpoint {
	return &endpointImpl{ah}
}
