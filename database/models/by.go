package models

import (
	"time"

	"gorm.io/gorm"
)

type ControlBy struct {
	CreateBy string `json:"createBy" gorm:"column:createdBy; index;comment:创建者"`
	UpdateBy string `json:"updateBy" gorm:"column:updateBy; index;comment:更新者"`
}

func (e *ControlBy) SetCreateBy(createBy string) {
	e.CreateBy = createBy
}

func (e *ControlBy) SetUpdateBy(updateBy string) {
	e.UpdateBy = updateBy
}

type Model struct {
	Id int `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
}

type ModelTime struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"column:createdAt;comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updatedAt;comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}
