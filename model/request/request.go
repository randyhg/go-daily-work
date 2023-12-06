package request

import (
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	Email       string
	AuthorityId string
	BufferTime  int64
	jwt.StandardClaims
}
