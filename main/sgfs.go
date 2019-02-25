package main

import (
	"sgfs/config"
	"sgfs/logger"
	"sgfs/service"
	"sgfs/util"
	"time"

	"github.com/buaazp/fasthttprouter"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func main() {
	// init log system
	logger.ConfigFilesystemLogger("sgfs", 7*24*time.Hour, 24*time.Hour)

	// load conf
	config.LoadConf()

	log.Info("Simple golang file server is starting...")

	// Create the uploaded file directory if it does not exist
	if err := util.CreateDirIfNotExist(config.GlobalConfig.Upload_Path); err != nil {
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
		GenerateIndexPages: true,

		// Open compression for bandwidth savings
		Compress: true,
	}

	go func() {
		if err := fasthttp.ListenAndServe(config.GlobalConfig.Visit_Port, fs.NewRequestHandler()); err != nil {
			log.Panic(err)
		}
	}()
}

func startOperationServer() {
	router := fasthttprouter.New()

	router.POST("/upload-file", service.UploadFileHandler)
	router.POST("/delete-file", service.DeleteFileHandler)

	fastServer := &fasthttp.Server{
		Handler:            router.Handler,
		MaxRequestBodySize: config.GlobalConfig.Max_Request_Body_Size,
	}

	if err := fastServer.ListenAndServe(config.GlobalConfig.Operation_Port); err != nil {
		log.Panic(err)
	}
}
