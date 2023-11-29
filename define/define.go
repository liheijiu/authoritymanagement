package define

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	// JWTKey 秘钥
	JWTKey = "sys-admin"
	// TokenExpire 有效期7天
	TokenExpire = time.Now().Add(time.Second * 3600 * 24 * 7).Unix()
	// ReFreshTokenExpire 刷新Token有限期14天
	ReFreshTokenExpire = time.Now().Add(time.Second * 3600 * 24 * 14).Unix()
	// DefaultSize 默认分页没有显示条数
	DefaultSize = 10
)

type UserClaim struct {
	Id      uint
	Name    string
	IsAdmin bool //是否超管
	jwt.StandardClaims
}
