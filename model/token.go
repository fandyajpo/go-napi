package model

type Token struct {
	AccessToken   string
	RefreshToken  string
	TokenExpire   int64
	RefreshExpire int64
}
