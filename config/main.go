package config

import (
	// "fmt"
	"io/ioutil"

	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MySql     MySqlConfig   `yaml:"mysql"`
	Logging   LoggingConfig `yaml:"logging"`
	CheckSign bool          `yaml:"check_sign"`
}

type MySqlConfig struct {
	ConnectTimeOut int         `yaml:"connect_time_out"`
	Charset        string      `yaml:"charset"`
	Main           MainConfig  `yaml:"main"`
	Admin          AdminConfig `yaml:"admin"`
}

type MainConfig struct {
	Dns      string `yaml:"dns"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DataBase string `yaml:"database"`
}

type AdminConfig struct {
	Dns      string `yaml:"dns"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DataBase string `yaml:"database"`
}

type LoggingConfig struct {
	FilePath     string        `yaml:"file_path"`
	FileWrite    bool          `yaml:"file_write"`
	FileMaxAge   time.Duration `yaml:"file_max_age"`
	RotationTime time.Duration `yaml:"file_rotation_time"`
}

//初始化配置
func InitConfig() (Config, error) {
	var conf Config
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
