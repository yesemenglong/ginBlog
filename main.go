package main

import (
	"fmt"
	"server/config"
	"server/global"
	"server/initialize"

	"github.com/spf13/viper"
)

func main() {
	config.Init()
	global.DB = initialize.Mysql()    // gorm连接数据库
	initialize.MysqlTables(global.DB) // 初始化表
	// 程序结束前关闭数据库链接
	db, _ := global.DB.DB()
	defer db.Close()

	router := initialize.InitRouter()
	router.Run(fmt.Sprintf(":%d", viper.GetInt("system.addr")))
}
