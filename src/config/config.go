package config

import (
	"errors"
	"flag"
	"ldbc-bcos/src/db"
	"ldbc-bcos/src/logger"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string            `mapstructure:"server_port"`
	LogConfig  *logger.LogConfig `mapstructure:"log_config"`
	DBConfig   *db.DBConfig      `mapstructure:"db_config"`
}

const DefaultServerPort = "9512"

// GetFlagPath --Specify the path and name of the configuration file (flag)
func GetFlagPath() string {
	var configPath string
	flag.StringVar(&configPath, "config", "../conf/config.yaml", "please input the system config file path")
	flag.Parse()
	return configPath
}

// InitConfig --Set config path and file name
func InitConfig(configPath string) (*Config, error) {
	var err error
	var conf Config
	if len(configPath) == 0 {
		configPath = GetFlagPath()
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(configPath)
	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	//var conf Config
	err = v.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}

	if conf.LogConfig == nil {
		conf.LogConfig = new(logger.LogConfig)
	}

	if conf.DBConfig == nil {
		return nil, errors.New("not found the db config")
	}

	if len(conf.ServerPort) == 0 {
		conf.ServerPort = DefaultServerPort
	}
	return &conf, nil
}
