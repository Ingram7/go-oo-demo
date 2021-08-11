package app

import (
	"fmt"
	"go-oo-demo/pkg/config"
	"log"
	"path/filepath"
)

type LogDBConfig struct {
	Host      string
	Name    string
	Username  string
	Password  string
	MaxIdle int
	MaxOpen int
	Table  string
}

func (c *LogDBConfig) String() string {
	return fmt.Sprintf("{Host:%s Name:%s Username:%s Password:%s MaxIdle:%d MaxOpen:%d Table:%s}", c.Host, c.Name, c.Username, c.Password, c.MaxIdle, c.MaxOpen, c.Table)
}

type LogConfig struct {
	EnableHook      bool
	ReportCaller bool
	Hook    string
	Database  *LogDBConfig
}

func (c *LogConfig) String() string {
	return fmt.Sprintf("{EnableHook:%v ReportCaller:%v Hook:%s Gorm:%+v}", c.EnableHook, c.ReportCaller, c.Hook, c.Database)
}

type ServerConfig struct {
	Host string
	Port string
}

func (c *ServerConfig) String() string {
	return fmt.Sprintf("{Host:%s Port:%s}", c.Host, c.Port)
}

type DBConfig struct {
	Host      string
	Name    string
	Username  string
	Password  string
	MaxIdle int
	MaxOpen int
}

func (c *DBConfig) String() string {
	return fmt.Sprintf("{Host:%s Name:%s Username:%s Password:%s MaxIdle:%d MaxOpen:%d}", c.Host, c.Name, c.Username, c.Password, c.MaxIdle, c.MaxOpen)
}

type Config struct {
	Mode     string
	Database *DBConfig
	Server   *ServerConfig
	Log *LogConfig

}

func newConfig() *Config {
	config := new(Config)
	return config
}

type configOption struct {
	configName string
	configType string
}

type configOptFunc func(opt *configOption)

func withConfigName(configName string) configOptFunc {
	return func(opt *configOption) {
		opt.configName = configName
	}
}

func withConfigType(configType string) configOptFunc {
	return func(opt *configOption) {
		opt.configType = configType
	}
}

func (c *Config) load(options ...configOptFunc) {

	opt := configOption{
		configName: "config",
		configType: "yaml",
	}

	for _, f := range options {
		f(&opt)
	}

	path := config.GetPathFromCmd()
	config.AddConfigPath(path)
	config.SetConfigName(opt.configName)
	config.SetConfigType(opt.configType)
	if err := config.ReadInConfig(); err != nil {
		log.Println(filepath.Abs(path))
		log.Fatalf("load config file error: %s, path: %s", err.Error(), path)
	}

	if err := config.Unmarshal(c); err != nil {
		log.Fatalf("config unmarshal error: %s", err.Error())
	}
}


func (c *Config) print() {
	fmt.Printf("config=%+v\n", c)
}

func (c *Config) mode() string {
	return c.Mode
}

func (c *Config) logConfig() *LogConfig {
	return c.Log
}

func (c *Config) dbConfig() *DBConfig {
	return c.Database
}

func (c *Config) serverConfig() *ServerConfig {
	return c.Server
}



