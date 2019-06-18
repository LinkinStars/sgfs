package config

import (
	"github.com/jinzhu/configor"
	"go.uber.org/zap"
)

var GlobalConfig = struct {
	// Maximum size of single upload file
	Max_Upload_Size int `default:"20"`

	// Maximum overall request size
	Max_Request_Body_Size int `default:"30"`

	// Upload file root directory
	Upload_Path string `default:"upload"`

	// File Operating Port
	Operation_Port string `default:"8080"`

	// File Visit port
	Visit_Port string `default:"8081"`

	// Operation certificate
	Operation_Token string `default:"123456"`

	// Generate a file directory index
	Generate_Index_Pages bool `default:"true"`
}{}

func LoadConf() {
	// load config info
	if err := configor.Load(&GlobalConfig, "conf.yml"); err != nil {
		panic(err)
	}

	log := zap.S()

	log.Infof("The current file server configuration is:  ")
	log.Infof("# Max_Upload_Size    %dM", GlobalConfig.Max_Upload_Size)
	log.Infof("# Max_Request_Body_Size    %dM", GlobalConfig.Max_Request_Body_Size)
	log.Infof("# Operation_Port     %s", GlobalConfig.Operation_Port)
	log.Infof("# Visit_Port     %s", GlobalConfig.Visit_Port)
	log.Infof("# Operation_Token     %s", GlobalConfig.Operation_Token)

	GlobalConfig.Max_Upload_Size *= MB
	GlobalConfig.Max_Request_Body_Size *= MB

	GlobalConfig.Operation_Port = ":" + GlobalConfig.Operation_Port
	GlobalConfig.Visit_Port = ":" + GlobalConfig.Visit_Port
}
