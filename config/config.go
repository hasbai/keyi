package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type MyConfig struct {
	Mode  string `default:"dev" env:"MODE"`
	Debug bool   `default:"false" env:"DEBUG"`
	// example: user:pass@tcp(127.0.0.1:3306)/dbname?parseTime=true
	// for more detail, see https://github.com/go-sql-driver/mysql#dsn-data-source-name
	DbUrl         string `default:"" env:"DB_URL"`
	RedisURL      string `default:"redis:6379" env:"REDIS_URL"`
	EmailHost     string `default:"" env:"EMAIL_HOST"`
	EmailPort     int    `default:"587" env:"EMAIL_PORT"`
	EmailUser     string `default:"" env:"EMAIL_USER"`
	EmailPassword string `default:"" env:"EMAIL_PASSWORD"`
	SiteName      string `default:"可易" env:"SITE_NAME"`
}

var Config MyConfig

func init() { // load config from environment variables
	fmt.Println("init config...")
	configType := reflect.TypeOf(Config)
	elem := reflect.ValueOf(&Config).Elem()
	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		// get default value
		defaultValue, defaultValueExists := field.Tag.Lookup("default")
		// get env variable name
		envName, ok := field.Tag.Lookup("env")
		if !ok {
			envName = strings.ToUpper(field.Name)
		}
		// get env variable value
		env := os.Getenv(envName)
		envExists := env != ""
		if !envExists {
			if !defaultValueExists {
				panic(fmt.Sprintf("Environment variable %s must be set!", field.Name))
			}
			env = defaultValue
		}
		var value any
		var err error
		switch field.Type.Kind() {
		case reflect.String:
			value = env
		case reflect.Int:
			value, err = strconv.Atoi(env)
			if err != nil {
				panic(fmt.Sprintf("Environment variable %s must be an int!", field.Name))
			}
		case reflect.Bool:
			lower := strings.ToLower(env)
			if lower == "true" {
				value = true
			} else if lower == "false" {
				value = false
			} else {
				panic(fmt.Sprintf("Environment variable %s must be a boolean!", field.Name))
			}
		default:
			panic("Now only supports string, int and bool as config")
		}
		elem.FieldByName(field.Name).Set(reflect.ValueOf(value))
	}
	if Config.Mode != "production" {
		Config.Debug = true
	}
}
