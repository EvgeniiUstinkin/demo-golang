package appconfig

import (
	"log"

	"github.com/gopher-lib/config"
)

var Config AppConfig

func Init(filename string) {
	if err := config.LoadFile(&Config, filename); err != nil {
		log.Fatalf("failed to load app configuration: %v", err)
	}
}

type Db struct {
	Host     string
	Name     string
	Port     string
	User     string
	Password string
	SslMode  string
	SslCert  string
}

type Storage struct {
	BucketName string
}

type AppConfig struct {
	Port           string
	Logger         bool
	ErrorReporting bool
	DB             Db      `mapstructure:"db"`
	Storage        Storage `mapstructure:"storage"`
}
