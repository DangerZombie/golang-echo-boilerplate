package auth

type authHelperImpl struct {
}

func NewAuthHelper() AuthHelper {
	return &authHelperImpl{}
}
