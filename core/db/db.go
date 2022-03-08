package db

import (
	"fmt"

	"jzsg.com/mca/core/server/config"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func InitDb() error {
	dbConf := config.GetDBConfig()
	config := &gorm.Config{NamingStrategy: &schema.NamingStrategy{SingularTable: true}}
	var err error
	//DB, err = gorm.Open(sqlite.Open("db/dbzl.db"), config)

	if dbConf.Type == "" {
		db, err = gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", dbConf.Name)), config)
	} else {
		dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConf.User, dbConf.Passwd, dbConf.Port, dbConf.Name)
		db, err = gorm.Open(mysql.Open(dsn), config)
	}

	if err != nil {
		return errors.WithMessage(err, "open db failed")
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(10)
	err = autoMigrateTable()
	if err != nil {
		return err
	}

	err = initSql()
	if err != nil {
		return err
	}
	return nil
}

func autoMigrateTable() error {

	db.AutoMigrate(&TUser{})
	//return DB.AutoMigrate(&TPolicy{})
	db.AutoMigrate(&TWfType{})
	db.AutoMigrate(&TLog{})
	db.AutoMigrate(&TRole{})
	db.AutoMigrate(&TUserRole{})
	//db.AutoMigrate(&CryptoDownload{})
	//db.AutoMigrate(&CryptoReport{})
	//db.AutoMigrate(&CryptoUsed{})
	////db.AutoMigrate(&GenesisCryptoUsage{})
	//db.AutoMigrate(&TConsistence{})
	//db.AutoMigrate(&CryptoFile{})
	//db.AutoMigrate(&MixConsistencyInfo{})
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func initSql() error {
	err := InitRole()
	if err != nil {
		return err
	}
	err = initWft()
	if err != nil {
		return err
	}

	err = initAdmin()
	if err != nil {
		return err
	}
	//fmt.Println("init admin success")
	return nil
}
