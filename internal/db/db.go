package db

import (
	"ema_sound_clone_api/config"
	"ema_sound_clone_api/pkg/model/entity"
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
)

var (
	db *gorm.DB
)

func Connect(env config.Env) {
	var err error

	dialect := mysql.Open(dsn(env))

	db, err = gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		log.Infof("Failed Connecting to database : %s", err)
	}

	db.Logger = db.Logger.LogMode(logger.Info)

	fmt.Printf("Connected mysql")
}

func Migrate() error {
	return GetDb().AutoMigrate(
		&entity.Account{},
		&entity.AccountRefreshToken{},
		&entity.AdminUser{},
	)
}

func dsn(env config.Env) string {
	var (
		user      = env.MysqlUser
		password  = env.MysqlPassword
		protocol  = env.MysqlProtocol
		dbName    = env.MysqlDatabase
		charset   = "utf8mb4"
		parseTime = "true"
		loc       = url.PathEscape("UTC") // default timezone is UTC
	)
	return fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=%s&loc=%s&sql_safe_updates=1", user, password, protocol, dbName, charset, parseTime, loc)
}

func GetDb() *gorm.DB {
	return db
}
