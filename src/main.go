package main

import (
	"basego/src/api"
	"basego/src/config"
	"basego/src/db"
	"basego/src/logger"
	"basego/src/server"
)

func main() {

	conf, err := config.InitConfig("")
	if err != nil {
		panic(err)
	}

	logBus := logger.NewLoggerBus(conf.LogConfig)

	zlog, err := logBus.GetZapLogger("Gorm")
	if err != nil {
		panic(err)
	}

	gormDb, err := db.GormInit(conf.DBConfig, db.TableSlice, zlog)
	if err != nil {
		panic(err)
	}

	s, err := server.NewServer(
		server.WithConfig(conf),
		server.WithGinEngin(),
		server.WithGormDb(gormDb),
		server.WithLog(logBus),
	)
	if err != nil {
		panic(err)
	}

	err = api.LoadHttpHandlers(s)
	if err != nil {
		panic(err)
	}

	err = s.Start()
	if err != nil {
		panic(err)
	}
}
