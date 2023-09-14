package config

import (
	"basego/src/db"
	"basego/src/logger"
	"errors"
	"flag"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort  string            `mapstructure:"server_port"`
	LogConfig   *logger.LogConfig `mapstructure:"log_config"`
	DBConfig    *db.DBConfig      `mapstructure:"db_config"`
	TmpFilePath string            `mapstructure:"tmp_file_path"`
}

const (
	DefaultServerPort  = "9606"
	DefaultTmpFilePath = "../tmp"
)

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

	if len(conf.TmpFilePath) == 0 {
		conf.TmpFilePath = DefaultTmpFilePath
	}

	return &conf, nil
}
