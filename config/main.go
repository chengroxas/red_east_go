package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	Conf Config
)

type Config struct {
	MySql MySqlConfig `yaml:"mysql"`
}

type MySqlConfig struct {
	ConnectTimeOut string      `yaml:"connect_time_out"`
	Main           MainConfig  `yaml:"main"`
	Admin          AdminConfig `yaml:"admin"`
}

type MainConfig struct {
	Dns      string `yaml:"dns"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

type AdminConfig struct {
	Dns      string `yaml:"dns"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

//初始化全部
func InitConfig() error {
	data, err := ioutil.ReadFile("./confi.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &Conf)
	if err != nil {
		return err
	}
	return nil
}
