package server

import (
	"basego/src/config"
	"basego/src/db"
	"basego/src/logger"
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	ctx       context.Context
	logBus    *logger.LoggerBus
	ginEngine *gin.Engine
	config    *config.Config
	gormDb    *gorm.DB
}
type Option func(s *Server)

func WithGinEngin() Option {
	return func(s *Server) {
		g := gin.New()
		s.ginEngine = g
	}
}

func WithConfig(cfg *config.Config) Option {
	return func(s *Server) {
		s.config = cfg
	}
}

func WithLog(logBus *logger.LoggerBus) Option {
	return func(s *Server) {
		s.logBus = logBus
	}
}

func WithContext(ctx context.Context) Option {
	return func(s *Server) {
		s.ctx = ctx
	}
}

func NewServer(opts ...Option) (*Server, error) {
	server := new(Server)
	for _, opt := range opts {
		opt(server)
	}

	return server, nil
}

func (s *Server) Start() error {
	zlog, err := s.logBus.GetZapLogger("Gorm")
	if err != nil {
		return err
	}

	gormDb, err := db.GormInit(s.config.DBConfig, db.TableSlice, zlog)
	if err != nil {
		return err
	}

	s.gormDb = gormDb

	go s.GinEngine().Run(":" + s.SeverPort())

	return nil
}

func (s *Server) GetZapLogger(name ...string) (*zap.SugaredLogger, error) {
	return s.logBus.GetZapLogger(name...)
}

func (s *Server) GinEngine() *gin.Engine {
	return s.ginEngine
}

func (s *Server) Db() *gorm.DB {
	return s.gormDb
}

func (s *Server) SeverPort() string {
	return s.config.ServerPort
}

func (s *Server) TmpFilePath() string {
	return s.config.TmpFilePath
}
