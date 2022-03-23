package user

//用户与群组的关联操作：添加用户到群组、从群组中删除用户、修改用户所在群组、展示用户所在群组

import (
	"fms/common"
	"fms/database"
	"fms/moddle"
	"fms/response"
	"fms/security"
	"github.com/gin-gonic/gin"
)

type JoinMember interface {
	common.RestInterface
}

type GroupMe struct {
	a interface{}
}

func NewJOinMem () JoinMember{
	a := &GroupMe{
		a : nil,
	}
	return a
}

func (gm *GroupMe)Create()gin.HandlerFunc{return func(c *gin.Context) {
		u := new(security.AdminGroupMem)
		d := new(moddle.UserInfo)
		b := new(moddle.UserGroup)
		if err := c.ShouldBind(&u);err != nil {
			response.Fail(c,nil,"数据验证错误")
			return
		}
		if u.Username == ""{
			response.Fail(c,nil,"用户名不不能为空")
			return
		}
		if u.Group == "" {
			response.Fail(c,nil,"分组名不能为空")
			return
		}
		d.Username = u.Username
		db := database.MysqlInit()
		db.Model(moddle.UserInfo{}).Where("username=?",d.Username).First(&d)
		if d.ID == 0 {
			response.Fail(c,nil,"用户不存在，请确认用户名")
			return
		}
		b.GroupName = u.Group
		db.Model(moddle.UserGroup{}).Where("group_name=?",b.GroupName).First(&b)
		if b.ID == 0 {
			response.Fail(c,nil,"分组不存在，请确认组名")
			return
		}
		y := &moddle.GroupMember{
			Username: u.Username,
			GroupID: b.ID,
			AdmUsername: u.Username,
		}
		j := new(moddle.GroupMember)
		db.Model(moddle.GroupMember{}).Where("username=?",y.Username).First(&j)
		if j.GroupID != 0 {
			response.Fail(c,nil,"用于已存在分组")
			return
		}
		db.Model(moddle.GroupMember{}).Create(&y)
		db.Model(moddle.UserGroup{}).Where("group_name=?",b.GroupName).First(&b)
		response.Success(c,gin.H{"group":b},"加入分组成功")
		_=db.Close()
}}
func (gm *GroupMe)Update()gin.HandlerFunc{
	return func(c *gin.Context) {
	a := new(moddle.UserInfo)
	b := new(moddle.UserGroup)
	e := new(moddle.GroupMember)
	u := new(security.AdminGroupMem)
	if err := c.ShouldBind(&u);err != nil {
		response.Fail(c,nil,"数据验证错误")
		return
	}
	if len(u.Username) == 0 {
		response.Fail(c,nil,"用户名不能为空")
		return
	}
	if len(u.Group) == 0 {
		response.Fail(c,nil,"分组名不能为空")
		return
	}
	a.Username = u.Username
	db := database.MysqlInit()
	db.Model(moddle.UserInfo{}).Where("username=?",a.Username).First(&a)
	if a.ID == 0 {
		response.Fail(c,nil,"用户不存在，请确认用户名")
		return
	}
	b.GroupName = u.Group
	db.Model(moddle.UserGroup{}).Where("group_name=?",b.GroupName).First(&b)
	if b.ID == 0{
		response.Fail(c,nil,"分组不存在，请确认分组名")
		return
	}
	e.Username = u.Username
	db.Model(moddle.GroupMember{}).Where("username=?",e.Username).First(&e)
	d := &moddle.GroupMember{Username: u.Username,GroupID: b.ID,ID: e.ID}
	db.Model(moddle.GroupMember{}).Where("id=?",d.ID).Update(&d).First(&d)
	response.Success(c,gin.H{"data":d},"已修改")
	_=db.Close()

}}
func (gm *GroupMe)Delete()gin.HandlerFunc{
	return func(c *gin.Context) {
		j := new(moddle.GroupMember)
		a := new(security.AdminGroupMem)
		if err := c.ShouldBind(&a);err != nil {
			response.Fail(c,nil,"数据验证错误")
			return
		}
	db := database.MysqlInit()
	db.Model(moddle.GroupMember{}).Where("username=?",a.Username).First(&j)
	if j.GroupID == 0{
		response.Fail(c,nil,"该用户未分组")
		return
	}
	u := new(moddle.GroupMember)
	db.Model(moddle.GroupMember{}).Where("username=?",a.Username).First(&u).Delete(&u)
	response.Success(c,nil,"删除成功")
		_=db.Close()
}}
func (gm *GroupMe)Show()gin.HandlerFunc{
	return func(c *gin.Context) {
		g := []moddle.GroupMember{}
		db := database.MysqlInit()
		result := db.Preload("Group").Preload("Adm").Find(&g)
		a := security.GroupMem(result.Value.(*[]moddle.GroupMember))
		response.Success(c,gin.H{"groups":a},"查询成功")
		_=db.Close()
	}}
