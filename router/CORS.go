package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin","http://192.168.12.177:8080")
		c.Writer.Header().Set("Access-Control-Max-Age","86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods","*")
		c.Writer.Header().Set("Access-Control-Allow-Headers","*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials","true")
		if c.Request.Method == http.MethodOptions{
			c.AbortWithStatus(http.StatusOK)
		}else {
			c.Next()
		}
	}
}
