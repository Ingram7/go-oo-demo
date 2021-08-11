package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var loader = viper.New()

var (
	AddConfigPath = loader.AddConfigPath
	SetConfigName = loader.SetConfigName
	SetConfigType = loader.SetConfigType
	ReadInConfig = loader.ReadInConfig
	Unmarshal = loader.Unmarshal
)

func GetPathFromCmd() string {
	pflag.String("configPath", "", "Server config file path")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	return viper.GetString("configPath")
}
