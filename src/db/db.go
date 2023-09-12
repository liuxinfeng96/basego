package db

import (
	"fmt"
	"ldbc-bcos/src/logger"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DBConfig struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	IP       string `mapstructure:"ip"`
	Port     string `mapstructure:"port"`
	DbName   string `mapstructure:"dbname"`
}

func getDBConfig(dbConf *DBConfig) string {
	mysqlURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConf.User, dbConf.Password, dbConf.IP, dbConf.Port, dbConf.DbName, "utf8")
	return mysqlURL
}

func GormInit(dbConf *DBConfig, tableSlice []interface{},
	zaplogger *zap.SugaredLogger) (*gorm.DB, error) {
	var err error

	glogger := logger.NewGormLogger(zaplogger, 200*time.Millisecond, false)
	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       getDBConfig(dbConf),
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: glogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := gormDb.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	// Set table options
	gormDb.Set("gorm:association_autoupdate", false).
		Set("gorm:association_autocreate", false).
		Set("gorm:table_options", "ENGINE=InnoDB")
	err = gormDb.Set("gorm:table_options", "CHARSET=utf8").
		Set("gorm:table_options", "COLLATE=utf8_general_ci").AutoMigrate(tableSlice...)
	if err != nil {
		return nil, err
	}

	return gormDb, nil
}
