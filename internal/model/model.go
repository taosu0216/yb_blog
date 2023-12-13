package model

import (
	"blog/pkg/setting"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GblueModel struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint   `json:"created_on"`
	ModifiedOn uint   `json:"modified_on"`
	DeletedOn  uint   `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(mysqlSetting *setting.MysqlSettings) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=Local",
		mysqlSetting.User,
		mysqlSetting.Password,
		mysqlSetting.Host,
		mysqlSetting.Port,
		mysqlSetting.Database,
		mysqlSetting.Charset,
		mysqlSetting.ParseTime,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
