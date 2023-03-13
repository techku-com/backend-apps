package Dto

import (
	"techku/Modules"
	"techku/Services/jwt"
)

type Utilities struct {
	Jwt     jwt.JwtServices
	Modules Modules.Modules
}
