package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var GlobalConfig = struct {
	// Maximum size of single upload file
	MaxUploadSize int `mapstructure:"max_upload_size"`

	// Maximum overall request size
	MaxRequestBodySize int `mapstructure:"max_request_body_size"`

	// Upload file root directory
	UploadPath string `mapstructure:"upload_path"`

	// File Operating Port
	OperationPort string `mapstructure:"operation_port"`

	// File Visit port
	VisitPort string `mapstructure:"visit_port"`

	// Operation certificate
	OperationToken string `mapstructure:"operation_token"`

	// Generate a file directory index
	GenerateIndexPages bool `mapstructure:"generate_index_pages"`
}{}

func LoadConf() {
	configVip := viper.New()
	configVip.SetConfigFile("conf.yml")

	if err := configVip.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := configVip.Unmarshal(&GlobalConfig); err != nil {
		panic(err)
	}

	log := zap.S()

	log.Info("The current file server configuration is:")
	log.Infof("%+v", GlobalConfig)

	GlobalConfig.MaxUploadSize *= MB
	GlobalConfig.MaxRequestBodySize *= MB

	GlobalConfig.OperationPort = ":" + GlobalConfig.OperationPort
	GlobalConfig.VisitPort = ":" + GlobalConfig.VisitPort
}
