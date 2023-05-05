package system

import (
	"filetool/database/models"
)

// 用户表
type ModelUserInfo struct {
	UserName string `json:"userName" gorm:"column:userName; size:128; comment:用户账号"`
	UserPwd  string `json:"userPwd" gorm:"column:userPwd; size:128; comment:用户访问密码"`
	NickName string `json:"nickName" gorm:"column:nickName; comment:用户昵称"`
	Salt     string `json:"salt" gorm:"column:salt; size:128; comment:加密盐"`
	Phone    string `json:"phone" gorm:"column:phone; size:128; comment:用户电话"`
	Avatar   string `json:"avatar" gorm:"column:avatar; comment:头像"`
	Sex      string `json:"sex" gorm:"column:sex; size:12; comment:性别"`
	Email    string `json:"email" gorm:"column:email; size:128; comment:邮箱"`
}

type TUserInfo struct {
	models.Model

	ModelUserInfo

	models.ControlBy
	models.ModelTime
}

func (TUserInfo) TableName() string {
	return "t_user_info"
}

func (e *TUserInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TUserInfo) GetId() interface{} {
	return e.Id
}
