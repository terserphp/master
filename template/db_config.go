package template

var DBConfigTemplate = `// auto-generated by terser-cli
// struct: DBConfig 数据库配置类
package model

import (
	"github.com/jinzhu/gorm"
)

// DBConfig struct 数据库配置类
type DBConfig struct {
	DBName     string
	DSN        string // Database Source Name: user:pswd@tcp(127.0.0.1:3306)db_name?charset=utf8&parseTime=True&loc=Local
	DriverName string
	dbServer   *gorm.DB
}

var defaultConfig *DBConfig

func DefaultConfig() (dc *DBConfig) {
	if defaultConfig == nil {
		// todo: 读取数据库配置
		dbName := "{{.DBName}}"
		dsn := "{{.DSN}}"
		dbDriver := "{{if .DriverName}}{{.DriverName}}{{else}}mysql{{end}}"

		defaultConfig = NewDBConfig(dbName, dsn, dbDriver)
	}

	return defaultConfig
}

func NewDBConfig(dbName, dsn, dbDriver string) *DBConfig {
	return &DBConfig{
		DBName:     dbName,
		DSN:        dsn,
		DriverName: dbDriver,
	}
}

func (dc *DBConfig) initDB() {
	if dc.dbServer != nil {
		return
	}

	db, err := gorm.Open(dc.DriverName, dc.DSN)
	if err != nil {
		// todo: 处理数据连接错误
		panic(err)
	}
	dc.dbServer = db
}

func (dc *DBConfig) GetDBServer() (db *gorm.DB) {
	dc.initDB()
	return dc.dbServer
}

func (dc *DBConfig) Close() (err error) {
	if dc.dbServer != nil {
		err = dc.dbServer.Close()
	}

	return err
}

`
