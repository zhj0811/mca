package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"jzsg.com/mca/common/define"
)

type Config struct {
	DBConfig `yaml:"db"`
	Log      LogConf     `json:"log" yaml:"log"`
	Admin    define.User `json:"admin" yaml:"admin"`
}

type DBConfig struct {
	Type   string `yaml:"type"`   //db type
	Name   string `yaml:"name"`   //db name
	Port   uint16 `yaml:"port"`   //db port
	User   string `yaml:"user"`   //db user
	Passwd string `yaml:"passwd"` //db password
}

type LogConf struct {
	Level string `json:"level" yaml:"level"`
	File  string `yaml:"file" json:"file"`
}

var config Config

func InitConfig(configFilePath string) error {
	log.Printf("Using config path %s\n", configFilePath)
	viper.SetConfigFile(configFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Viper read in config failed: %s\n", err)
		panic(err)
	}

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return errors.WithMessage(err, "read config file failed")
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return errors.WithMessage(err, "unmarshal yaml failed")
	}
	err = completeInitialization()
	if err != nil {
		return err
	}
	log.Println("Init config success.")
	return nil
}

func completeInitialization() error {
	if config.Log.File != "" {
		f, err := os.OpenFile(config.Log.File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		globalConf.writer = f
		log.SetOutput(globalConf.writer)
	}
	return nil
}

func GetDBConfig() DBConfig {
	return config.DBConfig
}

func GetAdminConf() define.User {
	return config.Admin
}
