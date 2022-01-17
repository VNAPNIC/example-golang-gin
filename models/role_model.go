package model

import (
	_ "gorm.io/gorm"
)

type Role struct {
	RoleId    int      `gorm:"NOT NULL;primary_key;" json:"role_id"` // role coding
	CreatedAt JSONTime `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt JSONTime `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
	DeletedAt JSONTime `gorm:"type:timestamp;default:current_timestamp" json:"deleted_at"`
	RoleName  string   `gorm:"type:varchar(128);" json:"role_name"` // role Name
	IsAdmin   bool     `gorm:"type:int(1);DEFAULT:0;NOT NULL;" json:"is_admin"`
	Status    int      `gorm:"type:int(1);DEFAULT:0;NOT NULL;" json:"status"`
	RoleKey   string   `gorm:"type:varchar(128);uniqueIndex;" json:"role_key"` // role code
}

func (Role) TableName() string {
	return TablePrefix + "role"
}

func CreateRole(auth Role) error {
	res := db.Create(&auth)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}
