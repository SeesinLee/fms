package moddle

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Administrator struct {
	ID int `json:"id"`
	Username string	`json:"username" gorm:"unique"`
}

type UserInfo struct {		//用户登陆信息
	ID int	`json:"id"`
	Username string	`json:"username"`
	Password string	`json:"password"`
	Status int `json:"status"`
	Adm	Administrator	`json:"administrator" gorm:"ForeignKey:AdmUsername"`
	AdmUsername string
	LoginDate Time `json:"login-date" gorm:"type:timestamp"`
}

type GetUseLoginStatus struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status int 	`json:"status"`
	LoginDate Time `json:"login-date" gorm:"type:timestamp"`
}

type UserGroup struct {		//群组信息
	ID	int	`json:"id"`
	GroupName string	`json:"groupName"`
}


type GroupUpdate struct {	//用于修改组名的对象
	OldName string `json:"oldName"`
	NewName string `json:"newName"`
}

type UaG struct {
	Username string	`json:"username"`
	GroupName string	`json:"groupName"`
}

type GroupMember struct {	//添加用户到群组的对象
	ID uuid.UUID	`json:"id" gorm:"type:char(36);primary_key"`
	Username string	`json:"username"`
	Group UserGroup	`json:"group" gorm:"ForeignKey:GroupID"`
	GroupID	int	`json:"group_id"`
	Adm Administrator `json:"adm" gorm:"ForeignKey:AdmUsername"`
	AdmUsername	string	`json:"adm_username"`
}
func (GM *GroupMember) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

type EditPassword struct {
	Username string	`json:"username"`
	OldPassword string	`json:"oldPassword"`
	NewPassword string	`json:"newPassword"`
	CorrectPassword string	`json:"correctPassword"`
}