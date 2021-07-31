package user

import (
	"context"
	"errors"
	"go-temp/global"
	"go-temp/model/user"
	"gorm.io/gorm"
	"time"
)

type JwtService struct {
}

func (jwtService *JwtService) JsonInBlacklist(jwtList user.JwtBlacklist) (err error) {
	err = global.DB.Create(&jwtList).Error
	return
}

//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	err := global.DB.Where("jwt = ?", jwt).First(&user.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}

//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string

func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.Redis.Get(context.Background(), userName).Result()
	return err, redisJWT
}

//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.CONFIG.JWT.ExpiresTime) * time.Second
	err = global.Redis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}
