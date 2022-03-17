package router

import (
	"fms/dashboard"
	"fms/list"
	"fms/midware"
	"fms/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var url string

func InitRouter(){
	r := gin.Default()
	r.Use(CORSHandler())
	r.POST("/login",user.Login())
	r.POST("/logout",user.Logout())
	r.Use(midware.UserAuthor(),midware.CheckBlackList())
	r.GET("/info",user.Info())
	Group := r.Group("/group")
	Group.Use(midware.AdminAuthor())
	Group.POST("/create",midware.AdminAuthor(),user.NewGroup().Create())
	Group.POST("/update",midware.AdminAuthor(),user.NewGroup().Update())
	Group.POST("/delete",midware.AdminAuthor(),user.NewGroup().Delete())
	Group.GET("/show",user.NewGroup().Show())
	Group.GET("/showDetails",midware.AdminAuthor(),user.NewGroup().ShowDetails())
	GroupMember := r.Group("/groupMember")
	GroupMember.Use(midware.AdminAuthor())
	GroupMember.POST("/create",midware.AdminAuthor(),user.NewJOinMem().Create())
	GroupMember.POST("/update",midware.AdminAuthor(),user.NewJOinMem().Update())
	GroupMember.POST("/delete",midware.AdminAuthor(),user.NewJOinMem().Delete())
	GroupMember.GET("/show",user.NewJOinMem().Show())
	List :=r.Group("/list")
	List.POST("/show",list.NewListsT().Create())
	List.POST("/showList",list.NewListsT().Show())
	List.POST("/delete",list.NewListsT().Delete())
	dashboardGroup := r.Group("/dashboard")
	dashboardGroup.GET("/sum",dashboard.InitDashboardAPI().FaultsSum())
	getLocation()
	err := r.Run(url)
	if err !=nil {
		panic(err)
		return
	}
}

func getLocation(){	//获取服务运行地址和端口
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	url = fmt.Sprintf("%s:%s",host,port)
}