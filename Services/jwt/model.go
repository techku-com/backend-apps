package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"techku/Config"
)

type JwtStruct struct {
	Config *Config.JwtSetting
}

type JwtServices interface {
	authKey() *jwtauth.JWTAuth
	Encode(ParamKeys jwt.MapClaims) (token string, err error)
	Decode(token string) (res map[string]interface{}, err error)
}
