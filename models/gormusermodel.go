package models

import (
	"filetool/pkg/utils"
	"time"

	log "github.com/pion/ion-log"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type TGormCyUserModel struct {
	db *gorm.DB
}

// NewTGormCyUserModel
//
//	@Description: 构造函数，为其封装数据库连接
//	@return *TGormCyUserModel
func NewTGormCyUserModel(db *gorm.DB) *TGormCyUserModel {
	return &TGormCyUserModel{
		db: db,
	}
}

// Insert
//
//	@Description: 插入一条用户记录
//	@param data   待插入的用户信息
//	@return error 成功返回nil,失败返回error
func (m *TGormCyUserModel) Insert(data *UserInfo) error {
	log.Infof(">>(m *TGormCyUserModel) Insert db: %v", m.db)

	timeLastAt := time.Now()

	var queryInfo UserInfo

	results := m.db.Table(data.TableName()).Where("userName = ?", data.UserName).First(&queryInfo)
	if errInsert := results.Error; errInsert != nil {
		if results.Error == gorm.ErrRecordNotFound {
			log.Infof(">>(m *TGormCyUserModel) Insert Create")

			data.CreatedAt = timeLastAt
			data.UpdatedAt = timeLastAt

			err := m.db.Table(data.TableName()).Create(&data).Error
			if err != nil {
				logx.Errorf("insert by %s error, reason: %s", data.TableName(), err)
				return err
			}
		} else {
			logx.Errorf("insert by %s error, reason: %s", data.TableName(), errInsert.Error())
			return errInsert
		}
	} else {
		logx.Errorf(">>(m *TGormCyUserModel) Insert error, reason:  %s", ErrFielExisted.Error())
		return ErrFielExisted
	}
	return nil
}

// Delete
//
//	@Description: 删除一条用户信息记录
//	@param id  待删除的用户信息id
//	@return error 成功返回nil,失败返回error
func (m *TGormCyUserModel) Delete(id int) error {
	var cyUser UserInfo
	err := m.db.Model(&cyUser).Where("id = ?", id).Delete(&cyUser).Error
	if err != nil {
		logx.Errorf("Delete by %s error, reason: %s", cyUser.TableName(), err)
		return err
	}
	return nil
}

// FindOne
//
//	@Description: 根据id查询相应的记录
//	@param  id 为主键
//	@return *CyUser 返回一条用户记录
//	@return error 成功返回nil，失败返回error
func (m *TGormCyUserModel) FindOne(id int) (*UserInfo, error) {
	var cyUser UserInfo
	cyUser = UserInfo{}
	err := m.db.Table(cyUser.TableName()).Where("id = ?", id).First(&cyUser).Error
	if err != nil {
		logx.Errorf("FindOne by %s error, reason: %s", cyUser.TableName(), err)
		return &cyUser, err
	}
	return &cyUser, nil
}

// Update
//
//	@Description: 更新用户记录
//	@param c 期望的用户记录
//	@return error 成功则返回nil，失败则返回error
func (m *TGormCyUserModel) Update(c *UserInfo) error {
	verity := NewTVerityModel(m.db)
	if err := verity.VerifyUniqueForUpdate(c.TableName(), "userName", c.UserName, utils.Int2str(c.Id)); err != nil {
		return err
	}

	err := m.db.Model(&c).Updates(c).Error
	if err != nil {
		logx.Errorf("Update by %s error, reason: %s", c.TableName(), err)
		return err
	}
	return nil
}

// List
//
//	@Description: 根据条件查询用户记录
//	@param fields 欲查询的条件
//	@param page 分页
//	@param size 页面大小
//	@return *[]CyUser 符合条件的用户记录
//	@return int64 成功返回符合条件的记录数, 失败返回0
//	@return error 成功则返回nil, 失败返回error
func (m *TGormCyUserModel) List(fieldsStr string, page, size int) (*[]UserInfo, int64, error) {
	var cyUser UserInfo
	var users []UserInfo

	offset := (page - 1) * size
	if offset < 0 {
		offset = 0
	}

	var count int64
	if fieldsStr == NULL {
		fieldsStr = " 1=1 "
	}
	err := m.db.Where(fieldsStr).Limit(size).Offset(offset).Find(&users).Error
	if err != nil {
		logx.Errorf("List by %s error, reason: %s", cyUser.TableName(), err)
		return &users, 0, err
	}

	//  单独计数，count和limit、offset不能混用
	//  count默认会统计软删除的记录，需要手动加上
	fieldsStr = fieldsStr + " and deleted_at IS NULL"
	m.db.Table(cyUser.TableName()).Where(fieldsStr).Count(&count)

	return &users, count, nil
}

// FindByUsername
//
//	@Description: 根据用户名查询用户记录
//	@param userName	用户名
//	@return *[]CyUser 满足条件的用户记录
//	@return error 成功返回nil，失败返回error
func (m *TGormCyUserModel) FindByUsername(userName string) (*UserInfo, error) {
	var cyUser UserInfo
	var user UserInfo

	err := m.db.Where("userName = ?", userName).First(&user).Error
	if err != nil {
		logx.Errorf("find by %s error, reason: %s", cyUser.TableName(), err)
		return &user, err
	}

	return &user, nil
}

// Update
//
//	@Description: 更新用户记录
//	@param c 期望的用户记录
//	@return error 成功则返回nil，失败则返回error
func (m *TGormCyUserModel) UpdatePwdByUsername(userName string, newPwd string) error {
	var c UserInfo
	err := m.db.Model(&c).Where("userName = ?", userName).Update("password", newPwd).Error
	if err != nil {
		logx.Errorf("Update by %s error, reason: %s", c.TableName(), err)
		return err
	}
	return nil
}
