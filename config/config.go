package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type MyConfig struct {
	SiteName  string `env:"SITE_NAME" envDefault:"可易"`
	BaseURL   string `env:"BASE_URL" envDefault:"http://localhost:8000"`
	Mode      string `env:"MODE" envDefault:"dev"`
	Debug     bool   `env:"DEBUG" envDefault:"false"`
	AppID     string `env:"APP_ID" envDefault:""`
	AppSecret string `env:"APP_SECRET" envDefault:""`
	// example: user:pass@tcp(127.0.0.1:3306)/dbname?parseTime=true
	// for more detail, see https://github.com/go-sql-driver/mysql#dsn-data-source-name
	DbType       string `env:"DB_TYPE" envDefault:"sqlite"`
	DbUrl        string `env:"DB_URL" envDefault:""`
	RedisURL     string `env:"REDIS_URL" envDefault:""` // redis:6379
	SecretKey    string `env:"SECRET_KEY" envDefault:""`
	SmtpHost     string `env:"SMTP_HOST" envDefault:""`
	SmtpPort     int    `env:"SMTP_PORT" envDefault:"587"`
	SmtpUser     string `env:"SMTP_USER" envDefault:""`
	SmtpPassword string `env:"SMTP_PASSWORD" envDefault:""`
	FromEmail    string `env:"FROM_EMAIL" envDefault:""`
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
	if Config.SecretKey == "" {
		fmt.Println("secret key not set!!!")
	}
	if Config.FromEmail == "" {
		Config.FromEmail = Config.SmtpUser
	}
}

func init() { // load config from environment variables
	initConfig()
}
