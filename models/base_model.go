package models

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"serverhealthcarepanel/utils/setting"
	"time"
)

var (
	db          *gorm.DB
	TablePrefix string
)

// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

type BaseModel struct {
	ID        uint     `gorm:"autoIncrement;primary_key" json:"id"`
	CreatedAt JSONTime `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt JSONTime `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
	DeletedAt JSONTime `gorm:"type:timestamp;default:current_timestamp" json:"deleted_at"`
}

func Setup() {
	database, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/healthcare_panel?charset=utf8")
	if err != nil {
		panic(err)
	}

	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)

	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: database,
	}), &gorm.Config{})

	if db == nil {
		return
	}

	/*    gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	      return TablePrefix + defaultTableName
	  }*/
	TablePrefix = setting.DatabaseSetting.TablePrefix

	err = db.AutoMigrate(&Auth{}, &Role{})
	if err != nil {
		panic(err)
	}

	fmt.Println("open  mysql  successfully!")
}
