package service

import (
	"path"
	"sgfs/config"
	"sgfs/util"
	"sgfs/util/date_util"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func UploadFileHandler(ctx *fasthttp.RequestCtx) {
	// Get the file from the form
	header, err := ctx.FormFile("file")
	if err != nil {
		SendResponse(ctx, -1, "No file was found.", nil)
		return
	}

	// Check File Size
	if header.Size > int64(config.GlobalConfig.Max_Upload_Size) {
		SendResponse(ctx, -1, "File size exceeds limit.", nil)
		return
	}

	// authentication
	token := string(ctx.FormValue("token"))
	if strings.Compare(token, config.GlobalConfig.Operation_Token) != 0 {
		SendResponse(ctx, -1, "Token error.", nil)
		return
	}

	// Check upload File Path
	uploadSubPath := string(ctx.FormValue("uploadSubPath"))
	if strings.Index(uploadSubPath, "/") != -1 {
		SendResponse(ctx, -1, "UploadSubPath error.", nil)
		return
	}

	visitPath := "/" + uploadSubPath + "/" + date_util.GetCurTimeFormat(date_util.YYYYMMDD)

	dirPath := config.GlobalConfig.Upload_Path + visitPath
	if err := util.CreateDirIfNotExist(dirPath); err != nil {
		log.Error(err)
		SendResponse(ctx, -1, "Failed to create folder.", nil)
		return
	}

	suffix := path.Ext(header.Filename)

	filename := createFileName(suffix)

	fileAllPath := dirPath + "/" + filename

	// Guarantee that the filename does not duplicate
	for {
		if !util.CheckPathIfNotExist(fileAllPath) {
			break
		}
		filename = createFileName(suffix)
		fileAllPath = dirPath + "/" + filename
	}

	// Save file
	if err := fasthttp.SaveMultipartFile(header, fileAllPath); err != nil {
		log.Error(err)
		SendResponse(ctx, -1, "Save file fail.", err.Error())
	}

	SendResponse(ctx, 1, "Save file success.", visitPath+"/"+filename)
	return
}

func createFileName(suffix string) string {
	// Date and Time + _ + Random Number + File Suffix
	return date_util.GetCurTimeFormat(date_util.YYYYMMddHHmmss) + "_" + util.GenerateRandomNumber(10) + suffix
}
