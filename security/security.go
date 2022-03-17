package security

import (
	"fms/moddle"
)

type UserS struct {
	ID int	`json:"id"`
	Username string	`json:"username"`
}

type HomeUser struct {
	Username string	`json:"username"`
	Group string `json:"group"`
	Role string `json:"role"`
	Menu []moddle.UserMenu `json:"menu"`
	LoginTime	moddle.Time	`json:"login_time"`
}

type AdminGroupMem struct {
	Group string	`json:"group"`
	Username string	`json:"username"`
	Role	string	`json:"role"`
}

type AdminGroupDetails struct {
	Group string	`json:"group"`
	Username []string	`json:"username"`
}

type AdminGroup struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

func GetUserS(u *moddle.UserInfo)*UserS{
	 a := &UserS{
		ID: u.ID,
		Username: u.Username,
	}
	return a
}

func HomeUserInfo(u moddle.GroupMember,e moddle.UserInfo,m *[]moddle.UserMenu)*HomeUser{
	a := &HomeUser{
		Username: u.Username,
		Group: u.Group.GroupName,
		LoginTime: e.LoginDate,
	}
	if e.Adm.ID == 0 {
		a.Role = "用户"
	}else {
		a.Role = "管理员"
	}
	a.Menu = *m
	return a
}

func GroupMem(group *[]moddle.GroupMember)*[]AdminGroupMem {
	g := make([]AdminGroupMem,0)
	a := AdminGroupMem{}
	for _,v := range *group{
		a.Username = v.Username
		a.Group = v.Group.GroupName
		if v.Adm.ID == 0 {
			a.Role = "用户"
		}else {
			a.Role = "管理员"
		}
		g = append(g,a)
	}
	return &g
}

func GroupN(group *[]*moddle.UserGroup)*[]AdminGroup{
	g := make([]AdminGroup,0)
	a := AdminGroup{}
	for _,v := range *group {
		a.Label = v.GroupName
		a.Value = v.GroupName
		g = append(g,a)
	}
	return &g
}

func GetGroupDetails(groupM *[]moddle.GroupMember,group *[]*moddle.UserGroup)*[]AdminGroupDetails{
	g := make([]AdminGroupMem,0)
	a := AdminGroupMem{}
	c := AdminGroupDetails{}
	b := make([]AdminGroupDetails,0)
	for _,v := range *groupM{
		a.Username = v.Username
		a.Group = v.Group.GroupName
		g = append(g,a)
	}
	for _,v := range *group{
		c.Group = v.GroupName
		for _,v := range g{
			if v.Group == c.Group {
				c.Username = append(c.Username,v.Username)
			}
		}
		b = append(b,c)
		c = AdminGroupDetails{}
	}
	return &b
}
