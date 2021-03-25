package service

import (
	"errors"
	"server/global"
	"server/model"
	"server/utils"

	"gorm.io/gorm"
)

// 添加用户
func CreateUser(data model.User) error {
	data.Password = utils.MD5V([]byte(data.Password))
	if !errors.Is(global.DB.Where("username = ?", data.Username).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同用户名")
	}

	return global.DB.Create(&data).Error
}

func GetUser(id uint) (err error, user model.User) {
	err = global.DB.Where("id = ?", id).First(&user).Error
	return err, user
}

// 分页查询
func GetUserList(username string, pageSize, pageNum int) (err error, list []model.User, total int64) {
	offset := (pageNum - 1) * pageSize
	db := global.DB.Model(&model.User{})
	if username != "" {
		db = db.Where("username LIKE ?", "%"+username+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return err, nil, 0
	} else {
		err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	}

	return err, list, total
}

// 修改用户
func UpdateUser(data model.User) error {
	var user model.User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	//if !errors.Is(global.DB.Where("username = ?", data.Username).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
	//	return errors.New("存在相同用户名")
	//}
	return global.DB.Model(&user).Where("id = ?", data.ID).Updates(maps).Error
}

// 删除用户
func DeleteUser(id int) error {
	var user model.User
	err := global.DB.Where("id = ?", id).Delete(&user).Unscoped().Error
	return err
}

func CheckLogin(u *model.User) (error, *model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err := global.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&u).Error
	return err, &user
}

// 修改密码
func ChangePassword(u model.User, newPassword string) (error, model.User) {
	var user model.User
	u.Password = utils.MD5V([]byte(u.Password))
	err := global.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, user
}
