package models

import (
	_ "gorm.io/gorm"
)

type Auth struct {
	BaseModel
	RoleId     int      `gorm:"Size:4;DEFAULT:0;NOT NULL;" json:"role_id"`
	Status     int      `gorm:"type:int(1);DEFAULT:0;NOT NULL;" json:"status"`
	Username   string   `gorm:"Size:20;UNIQUE_INDEX;NOT NULL;" json:"user_name"`
	Password   string   `gorm:"Size:50;NOT NULL;" json:"-"`
	RoleName   string   `gorm:"-" json:"role_name"`
	LoggedInAt JSONTime `gorm:"type:timestamp;default:current_timestamp" json:"logged_in_at"`
	Role       Role     `gorm:"foreignKey:RoleId;association_foreignkey:RoleId" json:"-"`
}
