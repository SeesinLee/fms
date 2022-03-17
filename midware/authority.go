package midware

import (
	"fms/database"
	"fms/moddle"
	"fms/response"
	"fms/security"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminAuthor()gin.HandlerFunc  {	//查询是否具有管理员权限
	return func(c *gin.Context) {
		u,b := c.Get("userInfo")
		if b == false {
			response.Fail(c,nil,"获取用户信息失败!")
			c.Abort()
			return
		}
		a := new(moddle.Administrator)
		a.Username = u.(*security.UserS).Username
		db := database.MysqlInit()
		db.Model(moddle.Administrator{}).Where("username=?",a.Username).First(&a)
		if a.ID == 0 {
			response.Request(c,http.StatusUnauthorized,404,nil,"您没有管理权限!")
			c.Abort()
			return
		}
		_=db.Close()
		c.Next()
	}
}