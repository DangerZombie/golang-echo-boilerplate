package parameter

import "go-echo/model/entity"

type JwtClaims struct {
	Issuer  string
	Subject string
	User    string
	Roles   []string
}

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

type CreateUserInput struct {
	entity.User
}

type CreateUserOutput struct {
	Id string
}

type FindUserRoleByUserIdInput struct {
	Id string
}

type FindUserRoleByUserIdOutput entity.User
