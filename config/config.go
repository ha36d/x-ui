package config

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed version
var version string

//go:embed name
var name string

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
)

func GetVersion() string {
	return strings.TrimSpace(version)
}

func GetName() string {
	return strings.TrimSpace(name)
}

func GetLogLevel() LogLevel {
	if IsDebug() {
		return Debug
	}
	logLevel := os.Getenv("XUI_LOG_LEVEL")
	if logLevel == "" {
		return Info
	}
	return LogLevel(logLevel)
}

func IsDebug() bool {
	return os.Getenv("XUI_DEBUG") == "true"
}

func GetDBUsername() string {
	return os.Getenv("XUI_MYSQL_USERNAME")
}

func GetDBPassword() string {
	return os.Getenv("XUI_MYSQL_PASSWORD")
}

func GetDBHost() string {
	return os.Getenv("XUI_MYSQL_HOST")
}

func GetDBPort() string {
	return os.Getenv("XUI_MYSQL_PORT")
}

func GetDB() string {
	return os.Getenv("XUI_MYSQL_DB")
}

func GetDBURI() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", GetDBUsername(), GetDBPassword(), GetDBHost(), GetDBPort(), GetDB())
}

func GetDBPath() string {
	return fmt.Sprintf("/etc/%s/%s.db", GetName(), GetName())
}
