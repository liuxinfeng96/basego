package router

import (
	"basego/src/handler"
	"basego/src/logger"
	"basego/src/server"
	"errors"
)

// 路由注册表
var routerList = []struct {
	path          string
	method        string
	isTokenVerify bool
	h             handler.Handler
}{
	// 测试接口
	{"test", "GET", false, &handler.TestHandler{}},

	// TODO 应用接口注册

}

func LoadHttpHandlers(s *server.Server) error {
	s.GinEngine().Use(handler.Cors())

	ginLogger := s.GetZapLogger("Gin")
	s.GinEngine().Use(logger.GinLogger(ginLogger))
	s.GinEngine().Use(logger.GinRecovery(ginLogger, true))

	for _, r := range routerList {
		switch r.method {
		case "POST":
			if r.isTokenVerify {
				s.GinEngine().POST(r.path, handler.JWTAuthMiddleware(), r.h.Handle(s))
			} else {
				s.GinEngine().POST(r.path, r.h.Handle(s))
			}

		case "GET":
			if r.isTokenVerify {
				s.GinEngine().GET(r.path, handler.JWTAuthMiddleware(), r.h.Handle(s))
			} else {
				s.GinEngine().GET(r.path, r.h.Handle(s))
			}

		default:
			return errors.New("unknown http request type")
		}
	}
	return nil
}
