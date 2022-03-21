package database

import (
	"fms/moddle"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/url"
)

var (
	DB *gorm.DB
	err error
)

func MysqlInit()*gorm.DB{	//初始化mysql
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	urlA := viper.GetString("mysql.url")
	port := viper.GetString("mysql.port")
	loc := viper.GetString("mysql.loc")
	databaseName := viper.GetString("mysql.databaseName")
	driverName := viper.GetString("mysql.driverName")
	charSet := viper.GetString("mysql.charSet")
	n := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&loc=%s&parseTime=True",
		username,
		password,
		urlA,
		port,
		databaseName,
		charSet,
		url.QueryEscape(loc))
	DB,err = gorm.Open(driverName,n)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	return DB
}

func CreatTable(){
	DB.AutoMigrate(moddle.UserInfo{})
	DB.AutoMigrate(moddle.Administrator{})
	DB.AutoMigrate(moddle.UserGroup{})
	DB.AutoMigrate(moddle.GroupMember{})
	DB.AutoMigrate(moddle.UserMenu{})
	DB.AutoMigrate(moddle.FaultList{})
}