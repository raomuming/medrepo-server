package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Mysql struct {
	Host     string
	User     string
	Password string
	DB       string
	Port     string
}

type Wechat struct {
	Appid  string `toml:"appid"`
	Secret string `toml:"secret"`
}

type Redis struct {
	IP   string `toml:"ip"`
	Port string `toml:"port"`
}

type Config struct {
	Debug  bool `toml:"debug"`
	Port   string
	Secret string
	Mysql  `toml:"mysql"`
	Wechat `toml:"wechat"`
	Redis  `toml:"redis"`
}

var config Config

var configFile = ""

func Get() Config {
	if config.Host == "" {
		filepath := getPath(configFile)

		if _, err := toml.DecodeFile(filepath, &config); err != nil {
			log.Fatal("配置文件读取失败!", err)
		}
	}
	return config
}

func SetPath(path string) {
	configFile = path
}

func getPath(path string) string {
	if path != "" {
		return path
	}

	path = os.Getenv("MEDREPO_CONF")
	if path != "" {
		return path
	}

	filepath := "config.toml"

	return filepath
}
