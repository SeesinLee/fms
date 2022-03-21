package ini

import (
	"fms/database"
	"fms/errorLog"
	"fms/prometheusAPI"
	"fms/router"
	"fms/util"
	"github.com/spf13/viper"
)

func Init (){
	errorLog.InitErrorLog() //初始化错误日志
	initViper()	//初始化viper
	database.MysqlInit()
	database.CreatTable()
	_=database.DB.Close()
	go prometheusAPI.InitGetCode() //监控业务应用状态
	util.NewChan()	//初始化检测管道
	go util.GetUserStatus() //获取用户状态
	router.InitRouter()		//初始化路由
}

func initViper(){		//viper初始化
	path := "./config"
	viper.SetConfigName("config")	//配置文件名
	viper.SetConfigType("yml")		//配置文件后缀
	viper.AddConfigPath(path)		//添加配置文件地址
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}