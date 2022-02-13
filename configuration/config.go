package configuration

import (
	"log"
	"os"
)

type DbConfig struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	SslMode  string
}

type WebConfig struct {
	Port string
}

type Config struct {
	Web      WebConfig
	Database DbConfig
	Timezone string
	Logger   log.Logger
}

var Configuration = InitiateConfig()

func InitiateConfig() (r Config) {
	r.Database = initiateDBConfig()
	r.Web = initiateWebConfig()
	r.Timezone = loadEnvvar("TZ", "Etc/UTC")
	r.Logger = *log.Default()
	return
}

func initiateDBConfig() (r DbConfig) {
	r.Host = loadEnvvar("DB_HOST", "localhost")
	r.Port = loadEnvvar("DB_PORT", "5432")
	r.Name = loadEnvvar("DB_NAME", "planwey")
	r.User = loadEnvvar("DB_USER", "planwey")
	r.Password = loadEnvvar("DB_PASSWORD", "planwey")
	r.SslMode = loadEnvvar("DB_SSL", "disable")
	return
}

func initiateWebConfig() (r WebConfig) {
	r.Port = loadEnvvar("WEB_PORT", "3000")
	return
}

func loadEnvvar(k string, default_return string) string {
	v, b := os.LookupEnv(k)
	if b {
		return v
	} else {
		return default_return
	}
}
