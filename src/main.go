package main

import (
	"ldbc-bcos/src/config"
	"ldbc-bcos/src/db"
	"ldbc-bcos/src/logger"
	"ldbc-bcos/src/router"
	"ldbc-bcos/src/server"
)

func main() {

	conf, err := config.InitConfig("")
	if err != nil {
		panic(err)
	}

	logBus := logger.NewLoggerBus(conf.LogConfig)

	zlog := logBus.GetZapLogger("Gorm")

	gormDb, err := db.GormInit(conf.DBConfig, db.TableSlice, zlog)
	if err != nil {
		panic(err)
	}

	s, err := server.NewServer(
		server.WithConfig(conf),
		server.WithGinEngin(),
		server.WithGormDb(gormDb),
		server.WithLog(&logBus),
	)
	if err != nil {
		panic(err)
	}

	err = router.LoadHttpHandlers(s)
	if err != nil {
		panic(err)
	}

	err = s.Start()
	if err != nil {
		panic(err)
	}
}
