package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type MyConfig struct {
	Mode  string `envDefault:"dev" env:"MODE"`
	Debug bool   `envDefault:"false" env:"DEBUG"`
	// example: user:pass@tcp(127.0.0.1:3306)/dbname?parseTime=true
	// for more detail, see https://github.com/go-sql-driver/mysql#dsn-data-source-name
	DbUrl string `envDefault:"" env:"DB_URL"`
}

var Config MyConfig

func initConfig() { // load config from environment variables
	fmt.Println("init config...")
	err := env.Parse(&Config)
	if err != nil {
		panic(err)
	}
	if Config.Mode != "production" {
		Config.Debug = true
	}
}

func init() { // load config from environment variables
	initConfig()
}
