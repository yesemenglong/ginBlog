package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;unique" json:"username" validator:"required"`
	Password string `gorm:"not null;comment:用户登录密码" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}
