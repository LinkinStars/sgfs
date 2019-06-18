package service

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

type ResponseJson struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(ctx *fasthttp.RequestCtx, code int, message string, data interface{}) {
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)

	responseJson := &ResponseJson{
		Code:    code,
		Message: message,
		Data:    data,
	}

	jsonStr, err := json.Marshal(responseJson)
	if err != nil {
		zap.S().Error("marshal json fail \n" + err.Error())
	}

	ctx.Response.SetBody(jsonStr)
}
