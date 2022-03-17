package user

import (
	"fms/database"
	"fms/moddle"
	"fms/response"
	"fms/security"
	"fms/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func Login()gin.HandlerFunc{
	return func(c *gin.Context){
		u := new(moddle.UserInfo)
		err := c.ShouldBind(&u)		//绑定获取到的参数
		if err != nil {
			response.Fail(c,nil,"获取参数失败!")
			return
		}
		if len(u.Username) == 0 {
			response.Fail(c, nil, "账号不能为空")
			return
		}
		if len(u.Password) == 0 {
			response.Fail(c, nil, "密码不能为空")
			return
		}
		db:=database.MysqlInit()
		defer db.Close()
		a := new(moddle.UserInfo)
		db.Model(moddle.UserInfo{}).Where("username=?", u.Username).First(&a)	//数据库寻找用户信息
		if a.ID == 0 {
			response.Fail(c, nil, "账号不存在")
			return
		}
		//err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(u.Password)) //密码hash过的

		if u.Password != a.Password {
			response.Fail(c, nil, "密码不正确")
			return
		}
		u.ID = a.ID
		token, err := ReleaseToken(u)
		if err != nil {
			response.Fail(c, nil, "token生成失败")
			return
		}
		response.Success(c, gin.H{"token": token}, "登陆成功")
		util.StatusChan.ID = u.ID
		util.StatusChan.Set(2)
	}
}

func Info()gin.HandlerFunc{
	return func(c *gin.Context) {
		u,b:= c.Get("userInfo")
		if b == false {
			response.Fail(c,nil,"未授权!")
			c.Abort()
			return
		}
		a := moddle.GroupMember{}
		e := moddle.UserInfo{}
		m := []moddle.UserMenu{}
		a.Username = u.(*security.UserS).Username
		e.Username = a.Username
		db:=database.MysqlInit()
		defer db.Close()
		db.Preload("Group").Where("username=?",a.Username).First(&a)
		db = database.MysqlInit()
		db.Preload("Adm").Where("username=?",e.Username).First(&e)
		if e.Adm.ID == 0{
			db = database.MysqlInit()
			db.Model(moddle.UserMenu{}).Where("role=?","1").Find(&m)
		} else {
			db = database.MysqlInit()
			db.Model(moddle.UserMenu{}).Find(&m)
		}
		d := security.HomeUserInfo(a,e,&m)
		response.Success(c,gin.H{"data":d},"获取用户信息成功!")
	}
}

func Logout()gin.HandlerFunc{
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if len(tokenString) == 0 || !strings.HasPrefix(tokenString,"Bearer ") {
			response.Fail(c,nil,"获取失败!")
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		database.InitRedis()
		defer database.DBR.Close()
		r,_:= database.DBR.Get().Do("SET",tokenString,tokenString,"EX","30000")
			response.Success(c,gin.H{"data":r},"退出成功!")
		util.StatusChan.Set(1)
	}
}
