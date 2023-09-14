package handler

import (
	"basego/src/server"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UploadFileResp struct {
	FileId string `json:"fileId"`
}

type UploadFileHandler struct{}

func (*UploadFileHandler) Handle(s *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {

		zLog, err := s.GetZapLogger("UploadFileHandler")
		if err != nil {
			FailedJSONResp(RespMsgLogServerError, c)
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			zLog.Errorf("get the file failed, err: [%s]", err.Error())
			FailedJSONResp(RespMsgServerError, c)
			return
		}

		fileId := uuid.New()

		dirPath := path.Join(s.TmpFilePath(), fileId.String())

		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			zLog.Errorf("mkdir the dir failed, err: [%s]", err.Error())
			FailedJSONResp(RespMsgServerError, c)
			return
		}

		filePath := path.Join(dirPath, file.Filename)

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			zLog.Errorf("save the file failed, err: [%s]", err.Error())
			FailedJSONResp(RespMsgServerError, c)
			os.RemoveAll(filePath)
			return
		}

		resp := new(UploadFileResp)
		resp.FileId = fileId.String()

		SuccessfulJSONResp(resp, "", c)
	}
}
