package main

import (
	"github.com/LinkinStars/golang-util/gu"
	"go.uber.org/zap"

	"github.com/LinkinStars/sgfs/config"
	"github.com/LinkinStars/sgfs/service"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func main() {
	gu.InitEasyZapDefault("sgfs")

	// load conf
	config.LoadConf()

	zap.S().Info("Simple golang file server is starting...")

	// Create the uploaded file directory if it does not exist
	if err := gu.CreateDirIfNotExist(config.GlobalConfig.Upload_Path); err != nil {
		panic(err)
	}

	startStaticFileServer()

	startOperationServer()
}

func startStaticFileServer() {
	fs := &fasthttp.FS{
		Root: config.GlobalConfig.Upload_Path,

		// Generate a file directory index. If true, access to the root path can see all the files stored.
		// In a production environment, it is recommended to set false
		GenerateIndexPages: config.GlobalConfig.Generate_Index_Pages,

		// Open compression for bandwidth savings
		Compress: true,
	}

	go func() {
		if err := fasthttp.ListenAndServe(config.GlobalConfig.Visit_Port, fs.NewRequestHandler()); err != nil {
			panic(err)
		}
	}()
}

func startOperationServer() {
	router := fasthttprouter.New()

	// Add panic handler
	router.PanicHandler = func(ctx *fasthttp.RequestCtx, err interface{}) {
		zap.S().Error(err)
		service.SendResponse(ctx, -1, "Unexpected error", err)
	}
	router.POST("/upload-file", service.UploadFileHandler)
	router.POST("/delete-file", service.DeleteFileHandler)

	fastServer := &fasthttp.Server{
		Handler:            router.Handler,
		MaxRequestBodySize: config.GlobalConfig.Max_Request_Body_Size,
	}

	if err := fastServer.ListenAndServe(config.GlobalConfig.Operation_Port); err != nil {
		panic(err)
	}
}
