package models

import (
	"time"

	log "github.com/pion/ion-log"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type TGormCyChatRecordModel struct {
	db *gorm.DB
}

// NewTGormCyChatRecordModel
//
//	@Description: 构造函数，为其封装数据库连接
//	@return *TGormCyChatRecordModel
func NewTGormCyChatRecordModel(db *gorm.DB) *TGormCyChatRecordModel {
	return &TGormCyChatRecordModel{
		db: db,
	}
}

// Insert
//
//	@Description: 插入一条聊天记录
//	@param data   待插入的聊天信息
//	@return error 成功返回nil,失败返回error
func (m *TGormCyChatRecordModel) Insert(data *ChatRecord) error {
	log.Infof(">>(m *TGormCyChatRecordModel) Insert db: %v", m.db)

	timeLastAt := time.Now()
	data.CreatedAt = timeLastAt
	data.UpdatedAt = timeLastAt
	err := m.db.Table(data.TableName()).Create(&data).Error
	if err != nil {
		logx.Errorf("insert by %s error, reason: %s", data.TableName(), err)
		return err
	}

	return nil
}

// Delete
//
//	@Description: 删除一条聊天信息记录
//	@param id  待删除的聊天信息id
//	@return error 成功返回nil,失败返回error
func (m *TGormCyChatRecordModel) Delete(id int) error {
	var cyChatRecord ChatRecord
	err := m.db.Model(&cyChatRecord).Where("id = ?", id).Delete(&cyChatRecord).Error
	if err != nil {
		logx.Errorf("Delete by %s error, reason: %s", cyChatRecord.TableName(), err)
		return err
	}
	return nil
}

// FindOne
//
//	@Description: 根据id查询相应的记录
//	@param  id 为主键
//	@return *CyChatRecord 返回一条聊天记录
//	@return error 成功返回nil，失败返回error
func (m *TGormCyChatRecordModel) FindOne(id int) (*ChatRecord, error) {
	var cyChatRecord ChatRecord
	cyChatRecord = ChatRecord{}
	err := m.db.Table(cyChatRecord.TableName()).Where("id = ?", id).First(&cyChatRecord).Error
	if err != nil {
		logx.Errorf("FindOne by %s error, reason: %s", cyChatRecord.TableName(), err)
		return &cyChatRecord, err
	}
	return &cyChatRecord, nil
}

// Update
//
//	@Description: 更新聊天记录
//	@param c 期望的聊天记录
//	@return error 成功则返回nil，失败则返回error
func (m *TGormCyChatRecordModel) Update(c *ChatRecord) error {
	err := m.db.Model(&c).Where("id = ?", c.Id).Updates(c).Error
	if err != nil {
		logx.Errorf("Update by %s error, reason: %s", c.TableName(), err)
		return err
	}
	return nil
}

// List
//
//	@Description: 根据条件查询聊天记录
//	@param fields 欲查询的条件
//	@param page 分页
//	@param size 页面大小
//	@return *[]CyChatRecord 符合条件的聊天记录
//	@return int64 成功返回符合条件的记录数, 失败返回0
//	@return error 成功则返回nil, 失败返回error
func (m *TGormCyChatRecordModel) List(fieldsStr string, page, size int) (*[]ChatRecord, int64, error) {
	var cyChatRecord ChatRecord
	var chatRecords []ChatRecord

	offset := (page - 1) * size
	if offset < 0 {
		offset = 0
	}

	var count int64
	if fieldsStr == NULL {
		fieldsStr = " 1=1 "
	}
	err := m.db.Order("updatedAt desc").Where(fieldsStr).Limit(size).Offset(offset).Find(&chatRecords).Error
	if err != nil {
		logx.Errorf("List by %s error, reason: %s", cyChatRecord.TableName(), err)
		return &chatRecords, 0, err
	}

	//  单独计数，count和limit、offset不能混用
	//  count默认会统计软删除的记录，需要手动加上
	fieldsStr = fieldsStr + " and deleted_at IS NULL"
	m.db.Table(cyChatRecord.TableName()).Where(fieldsStr).Count(&count)

	return &chatRecords, count, nil
}

// FindByChatRecordname
//
//	@Description: 根据聊天名查询聊天记录
//	@param chatRecordName	聊天名
//	@return *[]CyChatRecord 满足条件的聊天记录
//	@return error 成功返回nil，失败返回error
func (m *TGormCyChatRecordModel) FindByChatRecordname(chatRecordName string) (*ChatRecord, error) {
	var cyChatRecord ChatRecord
	var chatRecord ChatRecord

	err := m.db.Where("chatRecordName = ?", chatRecordName).First(&chatRecord).Error
	if err != nil {
		logx.Errorf("find by %s error, reason: %s", cyChatRecord.TableName(), err)
		return &chatRecord, err
	}

	return &chatRecord, nil
}
