package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const ConfigPath = "../../conf/config.yaml"

func TestInitConf(t *testing.T) {
	conf, err := InitConfig(ConfigPath)
	require.Nil(t, err)

	fmt.Printf("Config: %+v\n", conf)
	fmt.Printf("LogConfig: %+v\n", conf.LogConfig)
	fmt.Printf("DBConfig: %+v\n", conf.DBConfig)

}
