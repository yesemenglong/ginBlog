package service

import (
	"server/global"
	"server/model"

	"gorm.io/gorm"
)

// 添加文章
func CreateArticle(data model.Article) error {
	err := global.DB.Create(&data).Error
	return err
}

// 分页查询
func GetArticleList(pageSize, pageNum int, title string) (err error, userList []model.Article, total int64) {
	db := global.DB.Model(&model.Article{})
	offset := (pageNum - 1) * pageSize
	if title != "" {
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	err = db.Count(&total).Error
	err = db.Preload("Category").Limit(pageSize).Offset(offset).Find(&userList).Error
	return err, userList, total
}

// 修改文章
func UpdateArticle(data model.Article) error {
	var user model.Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := global.DB.Model(&user).Where("id = ?", data.ID).Updates(maps).Error
	return err
}

// 删除文章
func DeleteArticle(id int) error {
	var user model.Article
	err := global.DB.Where("id = ?", id).Delete(&user).Unscoped().Error
	return err
}

func FindArticle(id int) (err error, article model.Article) {
	err = global.DB.Preload("Category").Where("id = ?", id).First(&article).Error
	global.DB.Model(&article).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	return
}

func GetCateArticleList(id, pageSize, pageNum int) (err error, cateArticleList []model.Article, total int64) {
	offset := (pageNum - 1) * pageSize
	db := global.DB.Model(&model.Article{})
	err = db.Where("cid = ?", id).Count(&total).Error
	err = db.Preload("Category").Limit(pageSize).Offset(offset).Where("cid = ?", id).Find(&cateArticleList).Error
	return
}
