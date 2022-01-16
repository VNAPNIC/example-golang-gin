package models

import (
	_ "gorm.io/gorm"
)

type Role struct {
	RoleId    uint     `gorm:"primary_key;autoIncrement" json:"role_id"` // role coding
	CreatedAt JSONTime `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt JSONTime `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
	DeletedAt JSONTime `gorm:"type:timestamp;default:current_timestamp" json:"deleted_at"`
	RoleName  string   `gorm:"type:varchar(128);" json:"role_name"` // role Name
	IsAdmin   bool     `gorm:"type:int(1);DEFAULT:0;NOT NULL;" json:"is_admin"`
	Status    int      `gorm:"type:int(1);DEFAULT:0;NOT NULL;" json:"status"`
	RoleKey   string   `gorm:"type:varchar(128);UNIQUE_INDEX;" json:"role_key"` // role code
	RoleSort  int      `gorm:"type:int(4);" json:"role_sort"`                   // role ordering
	Remark    string   `gorm:"type:varchar(255);" json:"remark"`                // remark
}
