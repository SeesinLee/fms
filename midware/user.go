package midware

import (
	"fms/database"
	"fms/moddle"
	"fms/response"
	"fms/security"
	"fms/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func UserAuthor() gin.HandlerFunc {		//用户登陆中间件
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer "){
			response.Request(c,http.StatusUnauthorized,401,nil,"权限不足!")
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token,claim,err := user.ParseToken(tokenString)
		if err != nil || !token.Valid{
			response.Request(c,http.StatusUnauthorized,401,nil,"权限不足!")
			c.Abort()
			return
		}
		u := new(moddle.UserInfo)
		u.ID = claim.ID
		db:=database.MysqlInit()
		db.Model(moddle.UserInfo{}).Where("id=?",u.ID).First(&u)
		if u.Username == "" {
			response.Request(c,http.StatusUnauthorized,401,nil,"权限不足!")
			c.Abort()
			return
		}
		_=db.Close()
		a := security.GetUserS(u)
		c.Set("userInfo",a)
		c.Next()
	}
}
