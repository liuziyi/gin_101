package setting

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

type Server struct {
	RunMode      string        `yaml:"RunMode"`
	HttpPort     int           `yaml:"httpPort"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

type JWT struct {
	ExpireDuration time.Duration `yaml:"expireDuration"`
}

type Config struct {
	ServerConfig Server `yaml:"server"`
	JWTConfig    JWT    `yaml:"jwt"`
}

var ConfigSetting *Config

// Setup initialize the configuration instance
func Setup() {
	File, err := ioutil.ReadFile("conf/config.yaml")
	if err != nil {
		log.Fatalf("read ConfigSetting file failed, error: %v", err)
	}
	err = yaml.Unmarshal(File, &ConfigSetting)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/config.yaml': %v", err)
	}

	ConfigSetting.ServerConfig.ReadTimeout = ConfigSetting.ServerConfig.ReadTimeout * time.Second
	ConfigSetting.ServerConfig.WriteTimeout = ConfigSetting.ServerConfig.WriteTimeout * time.Second

	ConfigSetting.JWTConfig.ExpireDuration = ConfigSetting.JWTConfig.ExpireDuration * time.Hour
}
