package auth

type Auth struct{}

type AuthInterface interface {
	CreateToken(string) string
}

func NewAuth() *Auth {
	return &Auth{}
}

func (i *Auth) CreateToken(f string) string {
	return f
}
