package config

import (
	// "fmt"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	GinMode  string         `yaml:"gin_mode"`
	MySql    MySqlConfig    `yaml:"mysql"`
	Logging  LoggingConfig  `yaml:"logging"`
	Msg      MsgConfig      `yaml:"msg"`
	Redis    RedisConfig    `yaml:"redis"`
	Cache    CacheConfig    `yaml:"cache"`
	Memcache MemcacheConfig `yaml:"memcache"`
	Nsq      NsqConfig      `yaml:"nsq"`
	Sign     struct {
		Expire int  `yaml:"expire"`
		Check  bool `yaml:"check"`
		AppKey struct {
			Web     string `yaml:"web"`
			IOS     string `yaml:"ios"`
			Android string `yaml:"android"`
		} `yaml:"keys"`
	} `yaml:"sign"`
}

type MySqlConfig struct {
	ConnectTimeOut int         `yaml:"connect_time_out"`
	Charset        string      `yaml:"charset"`
	Debug          string      `yaml:"debug"`
	Prefix         string      `yaml:"prefix"`
	FileWrite      string      `yaml:"file_write"`
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

type MemcacheConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Prefix   string `yaml:"prefix"`
}

type CacheConfig struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
	Prefix   string `yaml:"prefix"`
}

type NsqConfig struct {
	TcpAddress        string `yaml:"tcp_address"`
	LookupdTcpAddress string `yaml:"lookupd_tcp_address"`
}

//初始化配置
func InitConfig() (Config, error) {
	var conf Config
	goPath := os.Getenv("GOPATH")
	data, err := ioutil.ReadFile(goPath + "/src/red_east_go/config.yaml")
	if err != nil {
		return conf, err
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}
