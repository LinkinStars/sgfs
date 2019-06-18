package service

import (
	"strings"

	"go.uber.org/zap"

	"github.com/LinkinStars/sgfs/config"
	"github.com/LinkinStars/sgfs/util/file_util"
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
		zap.S().Error(err)
		SendResponse(ctx, -1, "Delete file error.", err.Error())
		return
	}

	SendResponse(ctx, 1, "Delete file success.", nil)
	return
}
