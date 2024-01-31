package parameter

import "go-echo/model/entity"

type FindUserByUsernameAndPasswordInput struct {
	Username string
	Password string
}

type FindUserByUsernameAndPasswordOutput entity.User

type FindUserByIdInput struct {
	Id string
}

type FindUserByIdOutput struct {
	Id       string
	Nickname string
}
