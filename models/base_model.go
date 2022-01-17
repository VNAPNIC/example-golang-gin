package model

import (
	"database/sql"
	"fmt"
	"serverhealthcarepanel/utils/setting"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	ID uint `gorm:"primary_key" json:"id"`
	BaseModelTime
}

type BaseModelTime struct {
	CreatedAt time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime" json:"updated_at"`
	DeletedAt time.Time `gorm:"type:datetime" json:"deleted_at"`
}

func Setup() {
	database, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/healthcare_panel?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}

	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)

	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: database,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	}

	/*    gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	      return TablePrefix + defaultTableName
	  }*/
	TablePrefix = setting.DatabaseSetting.TablePrefix

	err = db.Debug().AutoMigrate(&Auth{}, &Role{})
	if err != nil {
		panic(err)
	}

	fmt.Println("open  mysql  successfully!")
}
