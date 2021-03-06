package model

import (
	_ "gorm.io/gorm"
)

type Role struct {
	RoleId uint `gorm:"NOT NULL;primary_key;" json:"role_id"` // role coding
	BaseModelTime
	RoleName string `gorm:"type:varchar(128);" json:"role_name"` // role Name
	IsAdmin  bool   `gorm:"type:int(1);DEFAULT:0;NOT NULL;" json:"is_admin"`
	Status   int    `gorm:"type:int(1);DEFAULT:0;NOT NULL;" json:"status"`
	RoleKey  string `gorm:"type:varchar(128);uniqueIndex;" json:"role_key"` // role code
}

func (Role) TableName() string {
	return TablePrefix + "role"
}

func RoleExists(roleKey string) bool {
	var exists bool
	res := db.Model(&Role{}).Select("count(*) > 0").Where("role_key = ?", &roleKey).Find(&exists).Error
	return res == nil && exists
}

func CreateRole(role Role) error {
	res := db.Create(&role)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
