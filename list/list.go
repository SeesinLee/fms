package list

import (
	"fms/database"
	"fms/moddle"
	"fms/response"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type ListsI interface {
	Create()gin.HandlerFunc
	Show()gin.HandlerFunc
	Delete()gin.HandlerFunc
}

type ListsT struct {
	a interface{}
}

func NewListsT() ListsI {
	a := &ListsT{
		a: nil,
	}
	return a
}

func (lt *ListsT)Create()gin.HandlerFunc{
	return func(c *gin.Context) {
		a := new(moddle.FaultList)
		d := new(moddle.FaultList)
		if err := c.ShouldBind(&d);err != nil {
			response.Fail(c,nil,"数据获取失败!")
			return
		}
		if d.ID == a.ID{
			if len(d.FaultSketch) == 0 {
				response.Fail(c,nil,"故障简述不能为空!")
				return
			}
			db := database.MysqlInit()
			db.Model(moddle.FaultList{}).Create(&d)
			_=db.Close()
			response.Success(c,gin.H{"data":d},"已保存!")
			c.Abort()
			return
		}
		db:=database.MysqlInit()
		db.Model(moddle.FaultList{}).Where("id=?",d.ID).Update(&d)
		db.Model(moddle.FaultList{}).Where("id=?",d.ID).First(&d)
		_=db.Close()
		response.Success(c,gin.H{"data":d},"已保存!")

	}
}

func (lt *ListsT)Delete()gin.HandlerFunc{
	return func(c *gin.Context) {
		a := new(moddle.FaultList)
		if err:=c.ShouldBind(&a);err !=nil {
			response.Fail(c,nil,"数据获取失败!")
			return
		}
		db := database.MysqlInit()
		db.Model(moddle.FaultList{}).Where("id=?",a.ID).First(&a).Delete(&a)
		_= db.Close()
		response.Success(c,nil,"删除成功!")

	}
}

func (lt *ListsT)Show()gin.HandlerFunc{
	return func(c *gin.Context) {
		//time.Sleep(time.Millisecond * 500)
		//a := []*moddle.FaultList{}
		//db:= database.MysqlInit()
		//db.Model(moddle.FaultList{}).Find(&a)
		//response.Success(c,gin.H{"data":a},"成功!")
		pageNum,_ :=strconv.Atoi(c.DefaultQuery("pageNum","1"))
		pageSize,_ := strconv.Atoi(c.DefaultQuery("pageSize","10"))

		a := []*moddle.FaultList{}
		var b int
		time.Sleep(time.Millisecond * 200)
		db := database.MysqlInit()
		db.Order("created_at desc").Offset((pageNum - 1)*pageSize).Limit(pageSize).Find(&a)
		db.Model(moddle.FaultList{}).Count(&b)
		_ = db.Close()
		response.Success(c,gin.H{"data":a,"total":b},"成功!")
	}
}