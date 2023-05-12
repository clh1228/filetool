package svc

import (
	"filetool/api/config"
	"filetool/database"
	"filetool/models"
	"filetool/pkg/p3000"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Db                     *gorm.DB
	Config                 config.Config
	P3000Sync              *p3000.P3000Conn
	TGormUserModel         *models.TGormCyUserModel
	TGormCyChatRecordModel *models.TGormCyChatRecordModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	p3000Conn := p3000.NewP3000Client(c)

	database.SetupDatabase()
	db := database.GetDataBase()

	return &ServiceContext{
		Db:                     db,
		Config:                 c,
		P3000Sync:              p3000Conn,
		TGormUserModel:         models.NewTGormCyUserModel(db),
		TGormCyChatRecordModel: models.NewTGormCyChatRecordModel(db),
	}
}
