package system

import (
	"filetool/database/models"
)

// 聊天记录表
type ModelChatRecordInfo struct {
	Title      string `json:"title" gorm:"column:title;comment:聊天记录"`
	ChatRecord string `json:"chatRecord" gorm:"column:chatRecord;comment:聊天记录"`
	FilePath   string `json:"filePath" gorm:"column:filePath;size:128;comment:文件路径"`
	UserName   string `json:"userName" gorm:"column:userName;size:128;comment:用户id"`
}

type TChatRecordInfo struct {
	models.Model

	ModelChatRecordInfo

	models.ControlBy
	models.ModelTime
}

func (TChatRecordInfo) TableName() string {
	return "t_chat_info"
}

func (e *TChatRecordInfo) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *TChatRecordInfo) GetId() interface{} {
	return e.Id
}
