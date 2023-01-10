package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var DB *gorm.DB

const (
	DBTypeMysql    = "mysql"
	DBTypeSqlite   = "sqlite"
	DBTypePostgres = "postgres"
)

var gormConfig = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
	},
	PrepareStmt: true, // use PrepareStmt for `Save`, `Update` and `Delete`
	//DisableForeignKeyConstraintWhenMigrating: true,
}

func postgresDB() (*gorm.DB, error) {
	fmt.Println("db type: postgres")
	return gorm.Open(postgres.Open(Config.DbUrl), gormConfig)
}

func mysqlDB() (*gorm.DB, error) {
	fmt.Println("db type: mysql")
	return gorm.Open(mysql.Open(Config.DbUrl), gormConfig)
}

func sqliteDB() (*gorm.DB, error) {
	fmt.Println("db type: sqlite")
	err := os.MkdirAll("data", 0750)
	if err != nil {
		panic(err)
	}
	return gorm.Open(sqlite.Open("data/sqlite.db"), gormConfig)
}

func memoryDB() (*gorm.DB, error) {
	fmt.Println("db type: memory")
	return gorm.Open(sqlite.Open("file::memory:?cache=shared"), gormConfig)
}

func init() {
	fmt.Println("init db config...")
	var err error

	if Config.Mode == "test" {
		DB, err = memoryDB()
	} else {
		switch Config.DbType {
		case DBTypeMysql:
			DB, err = mysqlDB()
		case DBTypeSqlite:
			DB, err = sqliteDB()
		case DBTypePostgres:
			DB, err = postgresDB()
		default:
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
