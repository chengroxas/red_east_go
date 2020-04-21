package config

import (
	// "fmt"
	"io/ioutil"

	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MySql   MySqlConfig   `yaml:"mysql"`
	Logging LoggingConfig `yaml:"logging"`
	Msg     MsgConfig     `yaml:"msg"`
	Sign    struct {
		Expire int  `yaml:"expire"`
		Check  bool `yaml:"check"`
		AppKey struct {
			Web     string `yaml:"web"`
			IOS     string `yaml:"ios"`
			Android string `yaml:"android"`
		} `yaml:"keys"`
	} `yaml:"sign"`
	Redis RedisConfig `yaml:"redis"`
}

type MySqlConfig struct {
	ConnectTimeOut int         `yaml:"connect_time_out"`
	Charset        string      `yaml:"charset"`
	Main           MainConfig  `yaml:"main"`
	Admin          AdminConfig `yaml:"admin"`
	Msg            MsgConfig   `yaml:"msg"`
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

type MsgConfig struct {
	Account string `yaml:"account"`
	Secret  string `yaml:"secret"`
	Auth    string `yaml:"auth"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	Prefix   string `yaml:"prefix"`
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
