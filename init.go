package database

import (
	"github.com/238Studio/child-nodes-error-manager/errpack"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitSQLiteDatabase 初始化SQLite数据库
// 传入：使用数据库名称，数据库地址
// 传出：数据库对象，错误
func InitSQLiteDatabase(dataBaseName, dataBaseURL string) (*DatabaseAPP, error) {
	databaseApp := new(DatabaseAPP)
	databaseApp.databaseConfig.databaseName = dataBaseName //设定数据库名称
	databaseApp.databaseConfig.databaseURL = dataBaseURL   //设定数据库地址

	db, err := gorm.Open(sqlite.Open(dataBaseName), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		//数据库初始化异常，界定为严重错误
		return nil, errpack.NewError(errpack.FatalException, errpack.Database, err)
	}
	databaseApp.db = db

	return databaseApp, nil
}
