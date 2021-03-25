package service

import (
	"server/global"
	"server/model"
)

func GetProfile(id int) (profile model.Profile, err error) {
	err = global.DB.Where("id = ?", id).First(&profile).Error
	return
}

func UpdateProfile(profile model.Profile) error {
	err := global.DB.Save(&profile).Error
	return err
}
