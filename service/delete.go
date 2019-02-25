package service

import (
	"sgfs/config"
	"sgfs/util/file_util"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func DeleteFileHandler(ctx *fasthttp.RequestCtx) {
	// authentication
	token := string(ctx.FormValue("token"))
	if strings.Compare(token, config.GlobalConfig.Operation_Token) != 0 {
		SendResponse(ctx, -1, "Token error.", nil)
		return
	}

	fileUrl := string(ctx.FormValue("fileUrl"))
	if len(fileUrl) == 0 {
		SendResponse(ctx, -1, "FileUrl error.", nil)
		return
	}

	fileUrl = config.GlobalConfig.Upload_Path + fileUrl
	if err := file_util.DeleteFile(fileUrl); err != nil {
		log.Error(err)
		SendResponse(ctx, -1, "Delete file error.", err.Error())
		return
	}

	SendResponse(ctx, 1, "Delete file success.", nil)
	return
}
