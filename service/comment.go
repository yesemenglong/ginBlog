package service

import (
	"server/global"
	"server/model"

	"gorm.io/gorm"
)

// 创建
func CreateComment(data *model.Comment) (err error) {
	err = global.DB.Create(&data).Error
	return err
}

// 查询单个
func FindComment(id uint) (c model.Comment, err error) {
	err = global.DB.Where("id = ?", id).First(&c).Error
	return c, err
}

// 查询所有评论列表
func GetCommentList(pageSize, pageNum int) (err error, commentList []model.Comment, total int64) {
	db := global.DB.Model(&model.Comment{})
	offset := (pageNum - 1) * pageSize
	err = db.Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Order("Created_At DESC").
		Select("comment.id, article.title,user_id,article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").
		Joins("LEFT JOIN article ON comment.article_id = article.id").
		Joins("LEFT JOIN user ON comment.user_id = user.id").Scan(&commentList).Error
	return
}

// 获取评论数量
func GetCommentCount(id int) (err error, total int64) {
	var c model.Comment
	err = global.DB.Find(&c).Where("article_id = ?", id).Where("status = ?", 1).Count(&total).Error
	return
}

// 当前页面的评论列表
func GetCommentListFront(id, pageSize, pageNum int) (err error, commentList []model.Comment, total int64) {
	db := global.DB.Model(&model.Comment{})
	offset := (pageNum - 1) * pageSize
	err = db.Where("article_id = ?", id).Where("status = ?", 1).Count(&total).Error
	err = db.Limit(pageSize).Offset(offset).Order("Created_At DESC").
		Select("comment.id, article.title, user_id, article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").
		Joins("LEFT JOIN article ON comment.article_id = article.id").
		Joins("LEFT JOIN user ON comment.user_id = user.id").
		Where("article_id = ?", id).Where("status = ?", 1).Scan(&commentList).Error
	return
}

// 删除评论
func DeleteComment(id uint) (err error) {
	var c model.Comment
	err = global.DB.Where("id = ?", id).Delete(&c).Error
	return
}

// 通过评论
func CheckComment(id int, data *model.Comment) (err error) {
	var c model.Comment
	var res model.Comment
	var a model.Article
	var maps = make(map[string]interface{})

	maps["status"] = data.Status
	err = global.DB.Model(&c).Where("id = ?", id).Updates(maps).First(&res).Error
	global.DB.Model(&a).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
	return
}

// 撤销评论
func UncheckComment(id int, data *model.Comment) (err error) {
	var c model.Comment
	var res model.Comment
	var article model.Article
	var maps = make(map[string]interface{})
	maps["status"] = data.Status

	err = global.DB.Model(&c).Where("id = ?", id).Updates(maps).First(&res).Error
	global.DB.Model(&article).Where("id = ?", res.ArticleId).UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	return
}
