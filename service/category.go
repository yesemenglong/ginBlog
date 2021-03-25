package service

import (
	"server/global"
	"server/model"
)

// 添加分类
func CreateCategory(data model.Category) error {
	err := global.DB.Create(&data).Error
	return err
}

// 分页查询
func GetCategoryList(pageSize, pageNum int) (err error, list []model.Category, total int64) {
	db := global.DB.Model(&model.Category{})
	offset := (pageNum - 1) * pageSize
	err = db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Find(&list).Error
	return err, list, total
}

// 修改分类
func UpdateCategory(data model.Category) error {
	var cate model.Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := global.DB.Model(&cate).Where("id = ?", data.ID).Updates(maps).Error
	return err
}

// 删除分类
func DeleteCategory(id int) error {
	var data model.Category
	err := global.DB.Where("id = ?", id).Delete(&data).Unscoped().Error
	return err
}

func GetCategory(id uint) (err error, data model.Category) {
	err = global.DB.Where("id = ?", id).First(&data).Error
	return err, data
}
