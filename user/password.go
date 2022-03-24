package user

import (
	"fms/database"
	"fms/moddle"
	"fms/response"
	"github.com/gin-gonic/gin"
)

type PassInterface interface {
	Edit()gin.HandlerFunc
}

type password moddle.EditPassword

func InitPassword() PassInterface {
	a := new(password)
	return a
}

func (p password)Edit()gin.HandlerFunc{
	return func(c *gin.Context) {
		if err := c.ShouldBind(&p);err != nil {
			response.Fail(c,nil,"数据获取错误!")
			return
		}
		if len(p.OldPassword) == 0 || len(p.NewPassword) == 0 || len(p.CorrectPassword) == 0{
			response.Fail(c,nil,"密码不能为空!")
			return
		}
		if p.CorrectPassword != p.NewPassword {
			response.Fail(c,nil,"两次密码必须一致!")
			return
		}
		u := moddle.UserInfo{
			Username: p.Username,
		}
		db := database.MysqlInit()
		db.Model(moddle.UserInfo{}).Where("username=?",u.Username).First(&u)
		u.Password = p.NewPassword
		db = database.MysqlInit()
		db.Model(moddle.UserInfo{}).Where("username=?",u.Username).Update(&u)
		_=db.Close()
		response.Success(c,gin.H{"data":u},"修改成功")

	}
}
