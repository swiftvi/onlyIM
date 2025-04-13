package ctx

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenJwtToken(uid string, secret string, expSec int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["uid"] = uid
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Second * time.Duration(expSec)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func GetUidFromToken(ctx context.Context) string {
	if uid, ok := ctx.Value("uid").(string); ok {
		return uid
	}
	return ""
}
