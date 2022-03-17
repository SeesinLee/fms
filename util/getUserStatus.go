package util

import (
	"fms/database"
	"fms/moddle"
	"github.com/jinzhu/gorm"
)

var b = new(moddle.UserInfo)
var d =new(gorm.DB)

func GetUserStatus(){
	for {
		i := StatusChan.Watch()
		if i == 0 {
			continue
		}
		b.ID = StatusChan.ID
		d = database.MysqlInit()
		d.Model(moddle.UserInfo{}).Where("id=?",b.ID).First(&b)
		b.Status = i
		d = database.MysqlInit()
		d.Model(moddle.UserInfo{}).Where("id=?",b.ID).Update(&b)
		_=d.Close()
	}
}