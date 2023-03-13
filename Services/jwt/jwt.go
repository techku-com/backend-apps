package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"techku/Config"
)

func (j JwtStruct) authKey() *jwtauth.JWTAuth {
	return jwtauth.New(j.Config.Encrypt, []byte(j.Config.Key), nil)
}

func (j JwtStruct) Encode(ParamKeys jwt.MapClaims) (token string, err error) {
	_, token, err = j.authKey().Encode(ParamKeys)
	return
}

func (j JwtStruct) Decode(token string) (res map[string]interface{}, err error) {
	jwtToken, err := j.authKey().Decode(token)
	if err != nil {
		return
	}
	res = jwtToken.PrivateClaims()
	return
}

func (j JwtStruct) GetConfig() *Config.JwtSetting {
	config := Config.JwtSetting{
		Key:     j.Config.Key,
		Encrypt: j.Config.Encrypt,
	}

	return &config
}
