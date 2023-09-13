package handler

import (
	"basego/src/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		method := c.Request.Method
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,Content-Type")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

type PageReq struct {
	PageSize int32  `json:"pageSize"`
	Page     int32  `json:"page"`
	SortType string `json:"sortType"`
}

type StandardResp struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type StandardRespWithPage struct {
	StandardResp
	Total int64 `json:"total"`
}

const (
	RespCodeSuccess = 200
	RespCodeFailed  = 500
)

func SuccessfulJSONResp(data interface{}, msg string, c *gin.Context) {
	resp := StandardResp{
		Code: RespCodeSuccess,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}

func SuccessfulJSONRespWithPage(data interface{}, total int64, c *gin.Context) {
	resp := StandardRespWithPage{
		StandardResp: StandardResp{
			Code: RespCodeFailed,
			Data: data,
		},
		Total: total,
	}
	c.JSON(http.StatusOK, resp)
}

func FailedJSONResp(msg string, c *gin.Context) {
	resp := StandardResp{
		Code: RespCodeFailed,
		Msg:  msg,
	}
	c.JSON(http.StatusOK, resp)
}

type Handler interface {
	Handle(s *server.Server) gin.HandlerFunc
}

type TestHandler struct {
}

func (th *TestHandler) Handle(s *server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		SuccessfulJSONResp("Hello,World!", "", ctx)
	}
}
