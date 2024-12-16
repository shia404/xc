package pkg

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var Token = &_token{}

type _token struct{}

// EncodeUId 加密
func (x *_token) EncodeUId(secretKey string, expire int64, value interface{}) (string, error) {
	iat := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + expire
	claims["iat"] = iat
	claims["uId"] = value
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tokenStr, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return tokenStr, errors.New("生成jwt失败")
	}
	return tokenStr, nil
}

func (x *_token) GetUId(ctx context.Context) int64 {
	return gconv.Int64(ctx.Value("uId"))
}
