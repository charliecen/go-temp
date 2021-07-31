package user

import (
	"errors"
	uuid2 "github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
	"go-temp/global"
	"go-temp/model/user"
	"go-temp/utils"
	"gorm.io/gorm"
)

type UserService struct {
}

// Register 用户注册
func (us *UserService) Register(u user.User) (err error, uo user.User) {
	var user user.User
	// 判断用户名是否已注册
	if !errors.Is(global.DB.Where("username = ?", u.Username).
		First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已注册"), uo
	}
	// 附加uuid， md5密码加密
	u.Password = utils.MD5V([]byte(u.Password))
	u.UUID = uuid2.UUID(uuid.NewV4())
	err = global.DB.Create(&u).Error
	return err, u
}

func (us *UserService) FindUserById(id int) (err error, user *user.User) {
	err = global.DB.Where("`id` = ?", id).First(&user).Error
	return
}

func (us *UserService) FindUserByUuid(uuid string) (err error, user *user.User) {
	if err = global.DB.Where("`uuid` = ?", uuid).First(&user).
		Error; err != nil {
		return errors.New("用户不存在"), user
	}
	return nil, user
}

func (us *UserService) Login(u *user.User) (err error, uo *user.User) {
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).
		Preload("Authorities").
		Preload("Authority").
		First(&uo).Error
	return
}
