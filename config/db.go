package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var DB *gorm.DB

const (
	DBTypeMysql = iota
	DBTypeSqlite
)

var DBType uint

var gormConfig = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
	},
	PrepareStmt: true, // use PrepareStmt for `Save`, `Update` and `Delete`
	//DisableForeignKeyConstraintWhenMigrating: true,
}

func mysqlDB() (*gorm.DB, error) {
	DBType = DBTypeMysql
	fmt.Println("db type: mysql")
	return gorm.Open(mysql.Open(Config.DbUrl), gormConfig)
}

func sqliteDB() (*gorm.DB, error) {
	DBType = DBTypeSqlite
	fmt.Println("db type: sqlite")
	err := os.MkdirAll("data", 0750)
	if err != nil {
		panic(err)
	}
	return gorm.Open(sqlite.Open("data/sqlite.db"), gormConfig)
}

func memoryDB() (*gorm.DB, error) {
	DBType = DBTypeSqlite
	fmt.Println("db type: memory")
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), gormConfig)
}

func init() {
	fmt.Println("init db config...")
	var err error
	if Config.DbUrl != "" {
		DB, err = mysqlDB()
	} else {
		if Config.Mode == "test" {
			DB, err = memoryDB()
		} else {
			DB, err = sqliteDB()
		}
	}
	if Config.Debug {
		DB = DB.Debug()
	}
	if err != nil {
		panic(err)
	}
}
