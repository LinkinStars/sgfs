package logger

import (
	"path"
	"sgfs/util"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/onrik/logrus/filename"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var logPath = "log"

func init() {
	err := util.CreateDirIfNotExist(logPath)
	if err != nil {
		panic(err)
	}

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.00",
	})
	log.SetLevel(logrus.DebugLevel)
}

// config log file system
// maxAge: Maximum save time of files
// rotationTime: Log Cutting Interval
func ConfigFilesystemLogger(logFileName string, maxAge time.Duration, rotationTime time.Duration) {
	logPath := path.Join(logPath, logFileName)

	// init err log
	errWriter, err := rotatelogs.New(
		logPath+"_err_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		log.Errorf("init err log system fail.      %+v", errors.WithStack(err))
	}

	// init info log
	infoWriter, err := rotatelogs.New(
		logPath+"_info_%Y-%m-%d.log",
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		log.Errorf("init info log system fail.      %+v", errors.WithStack(err))
	}

	// set log hook
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: infoWriter,
		log.InfoLevel:  infoWriter,
		log.WarnLevel:  infoWriter,
		log.ErrorLevel: errWriter,
		log.FatalLevel: errWriter,
		log.PanicLevel: errWriter,
	}, &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.00",
	})
	log.AddHook(lfHook)

	// print line number
	filenameHook := filename.NewHook()
	filenameHook.Field = "line"
	log.AddHook(filenameHook)
}
