package response

//响应封装

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context,data interface{},msg string){
	c.JSON(http.StatusOK,gin.H{
		"code" : 200,
		"data" : data,
		"msg" : msg,
	})
}

func Fail(c *gin.Context,data interface{},msg string){
	c.JSON(http.StatusBadRequest,gin.H{
		"code" : 400,
		"data" : data,
		"msg" : msg,
	})
}

func Request(c *gin.Context,httpStatus int,code int,data interface{},msg string){
	c.JSON(httpStatus,gin.H{
		"code" : code,
		"data" : data,
		"msg" : msg,
	})
}