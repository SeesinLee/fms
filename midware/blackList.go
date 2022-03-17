package midware

import (
	"fms/response"
	"fms/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CheckBlackList()gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString:=c.GetHeader("Authorization")
		if len(tokenString) == 0 || !strings.HasPrefix(tokenString,"Bearer "){
			response.Request(c,http.StatusUnauthorized,401,nil,"未授权!")
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		b := util.GetExist(tokenString)
		if b == int64(1) {
			response.Request(c,http.StatusUnauthorized,401,nil,"未授权！")
			c.Abort()
			return
		}
		c.Next()
	}
}