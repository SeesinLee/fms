package user

//群组的相关操作：新增、修改、删除、查询展示

import (
	"fms/common"
	"fms/database"
	"fms/moddle"
	"fms/response"
	"fms/security"
	"fmt"
	"github.com/gin-gonic/gin"
)

type GroupInter interface {
	common.RestInterface
	ShowDetails()gin.HandlerFunc
}

type GroupS struct {
	a interface{}
}

func NewGroup()GroupInter{
	a := &GroupS{
		a : nil,
	}
	return a
}

func (gs *GroupS)Create()gin.HandlerFunc{	//添加群组
	return func(c *gin.Context){
		g := new(moddle.UserGroup)
		gp := new(moddle.UserGroup)
		if err := c.ShouldBind(&g);err !=nil{
			response.Fail(c,nil,"数据验证错误")
			return
		}
		if g.GroupName == "" {
			response.Fail(c,nil,"分组名不能为空")
			return
		}
		db:=database.MysqlInit()
		db.Model(moddle.UserGroup{}).Where("group_name=?", g.GroupName).First(&gp)
		if gp.GroupName != "" {
			response.Fail(c, gin.H{"group": g.GroupName}, "已存在分组")
			return
		}
		database.DB.Model(moddle.UserGroup{}).Create(&g)
		database.DB.Model(moddle.UserGroup{}).Where("group_name=?", g.GroupName).First(&gp)
		if g.ID == 0 {
			response.Fail(c, nil, "创建分组失败")
			return
		}
		response.Success(c, gin.H{"group": gp.GroupName}, "创建成功")
		_=db.Close()
	}

}
func (gs *GroupS)Update()gin.HandlerFunc{	//修改群组
	return func(c *gin.Context){
		g := new(moddle.GroupUpdate)
		gp := new(moddle.UserGroup)
		if err := c.ShouldBind(&g); err != nil {
			response.Fail(c,nil, "数据验证错误")
			return
		}
		fmt.Println(g)
		if g.NewName == "" || len(g.NewName) == 0{
			response.Fail(c,nil,"分组名不能为空")
			return
		}
		gp.GroupName = g.OldName
		db:=database.MysqlInit()
		db.Model(moddle.UserGroup{}).Where("group_name=?", gp.GroupName).First(&gp)
		if gp.ID == 0 {
			response.Fail(c, nil, "分组不存在")
			return
		}
		gp.GroupName = g.NewName
		database.DB.Model(moddle.UserGroup{}).Where("id=?",gp.ID).Update(&gp)
		database.DB.Model(moddle.UserGroup{}).Where("id=?", gp.ID).First(&gp)
		response.Success(c, gin.H{"group": gp.GroupName}, "修改成功")
		_=db.Close()
	}
}
func (gs *GroupS)Delete()gin.HandlerFunc{	//删除群组
	return func(c *gin.Context){
		g := new(moddle.UserGroup)
		gp := new(moddle.UserGroup)
		if err := c.ShouldBind(&g);err != nil {
			response.Fail(c,nil,"数据验证错误")
			return
		}
		if g.GroupName == ""{
			response.Fail(c,nil,"分组名不能为空")
			return
		}
		db:=database.MysqlInit()
		db.Model(moddle.UserGroup{}).Where("group_name=?",g.GroupName).First(&gp)
		if gp.ID == 0{
			response.Fail(c,nil,"分组不存在")
			return
		}
		database.DB.Model(moddle.UserGroup{}).Where("id=?",gp.ID).Delete(&g)
		response.Success(c,nil,"删除成功")
		_=db.Close()
	}
}
func (gs *GroupS)Show()gin.HandlerFunc{	//展示群组
	return func(c *gin.Context){
		g := []*moddle.UserGroup{}
		db:=database.MysqlInit()
		result := db.Model(moddle.UserGroup{}).Find(&g)
		b := security.GroupN(result.Value.(*[]*moddle.UserGroup))
		response.Success(c, gin.H{"data": b}, "已显示所有分组")
		_=db.Close()
	}
}

func (gs *GroupS)ShowDetails()gin.HandlerFunc{
	return func(c *gin.Context) {
		g := []*moddle.UserGroup{}
		b :=[]moddle.GroupMember{}
		db:=database.MysqlInit()
		result1 := db.Model(moddle.UserGroup{}).Find(&g)
		result2 := db.Preload("Group").Find(&b)
		d := security.GetGroupDetails(result2.Value.(*[]moddle.GroupMember),result1.Value.(*[]*moddle.UserGroup))
		response.Success(c,gin.H{"data":d},"以显示所有群组信息")
		_=db.Close()
	}
}