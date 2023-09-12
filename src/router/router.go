package router

import (
	"errors"
	"github.com/gin-gonic/gin"
	"ldbc-bcos/src/handler"
	"ldbc-bcos/src/logger"
	"ldbc-bcos/src/server"
)

func LoadHttpHandlers(s *server.Server) error {
	s.GinEngine().Use(handler.Cors())

	ginLogger := s.GetZapLogger("Gin")
	s.GinEngine().Use(logger.GinLogger(ginLogger))
	s.GinEngine().Use(logger.GinRecovery(ginLogger, true))

	for _, h := range handler.HttpHandlerList {
		group := s.GinEngine().Group(h.RouterGroup())
		err := loadRouterGroupHandler(group, h, s)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadRouterGroupHandler(routerGroup *gin.RouterGroup, h handler.Handler, s *server.Server) error {
	switch h.HttpMethod() {
	case "POST":
		if h.IsTokenVerify() {
			routerGroup.POST(h.RouterName(), handler.JWTAuthMiddleware(), h.Handle(s))
		} else {
			routerGroup.POST(h.RouterName(), h.Handle(s))
		}
	case "GET":
		if h.IsTokenVerify() {
			routerGroup.GET(h.RouterName(), handler.JWTAuthMiddleware(), h.Handle(s))
		} else {
			routerGroup.GET(h.RouterName(), h.Handle(s))
		}
	default:
		return errors.New("unknown http request type")
	}
	return nil
}
