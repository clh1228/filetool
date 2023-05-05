package models

import (
	"fmt"

	log "github.com/pion/ion-log"
	"gorm.io/gorm"
)

type TVerityModel struct {
	db *gorm.DB
}

func NewTVerityModel(db *gorm.DB) *TVerityModel {
	return &TVerityModel{
		db: db,
	}
}

// Verify the uniqueness of the field for add
func (m *TVerityModel) VerifyUniqueForAdd(table, field, fieldVue string) error {
	// 对空值不设置唯一性检验
	if fieldVue == "" {
		return nil
	}

	queryCount := fmt.Sprintf("select count(*) from %s where `%s` = '%s' and `deleted_at` is null", table, field, fieldVue)
	var num int64
	err := m.db.Raw(queryCount).Scan(&num).Error
	if err != nil {
		log.Errorf(">>func (m *TVerityModel) VerifyUniqueForAdd() Failed! Err: [%v]", err)
		err = ErrCountNum
		return err
	} else if num >= 1 {
		return ErrFielExisted
	}
	return nil
}

// Verify the unquness of the field for update
func (m *TVerityModel) VerifyUniqueForUpdate(table, field, fieldVue, id string) error {
	// 对空值不设置唯一性检验
	if fieldVue == "" {
		return nil
	}

	queryCount := fmt.Sprintf("select count(*) from %s where `%s` = '%s' and `id` != %s and `deleted_at` is null", table, field, fieldVue, id)
	var num int64
	err := m.db.Debug().Raw(queryCount).Scan(&num).Error
	if err != nil {
		log.Errorf(">>func (m *TVerityModel) VerifyUniqueForUpdate() Failed! Err: [%v]", err)
		err = ErrCountNum
		return err
	} else if num >= 1 {
		return ErrFielExisted
	}
	return nil
}
