package projectConfig

import "time"

type BaseInfo struct {
	env string
}

type Config struct {
	BaseConfig     `ini:"BASE_CONFIG"`
	ServerConfig   `ini:"SERVER_CONFIG"`
	LogConfig      `ini:"LOG"`
	DatabaseConfig `ini:"DATABASE"`
	RedisConfig    `ini:"REDIS"`
	ElasticConfig  `ini:"ELASTIC"`
	WXAPPConfig    `ini:"WXAPP"`
	UserConfig     `ini:"USER_CONFIG"`
}

type ServerConfig struct {
	RUN_MODE      string
	HTTP_PORT     int
	READ_TIMEOUT  int
	WRITE_TIMEOUT int
}

type BaseConfig struct {
	ENV string

	QR_CODE_SAVE_PATH string
	RUNTIME_ROOT_PATH string
	IMAGE_SAVE_PATH   string
	IMAGE_MAX_SIZE    int
	IMAGE_ALLOW_EXTS  string
	EXPORT_SAVE_PATH  string
	PREFIX_URL        string

	PAGE_SIZE int

	JWT_SECRET      string
	JWT_VERIFY_TYPE string

	REDIS_COMMON_EXPIRATION_TIME time.Duration
}

type UserConfig struct {
	USER_LOGIN_EXPIRATION_TIME time.Duration
}

type LogConfig struct {
	LOG_SAVE_URL  string
	LOG_SAVE_NAME string
}

type DatabaseConfig struct {
	DRIVER_NAME string
	HOST        string
	PORT        string
	DATABASE    string
	USER_NAME   string
	PASSWORD    string
	CHARSET     string
}

type RedisConfig struct {
	HOST         string
	PORT         string
	PASSWORD     string
	CHARSET      string
	PoolSize     int
	PoolTimeout  int
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
}

type ElasticConfig struct {
	HOST string
	PORT string
}

type WXAPPConfig struct {
	APPID  string
	SECRET string
}
