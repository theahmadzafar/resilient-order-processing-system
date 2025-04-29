package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Response struct {
	Status  int                    `json:"status"`
	Success bool                   `json:"success"`
	Meta    map[string]interface{} `json:"meta"`
	Data    interface{}            `json:"data"`
}

func new(status int, meta map[string]interface{}, data interface{}) *Response {
	success := false
	if status >= 200 && status <= 299 {
		success = true
	}

	response := &Response{
		Status:  status,
		Success: success,
		Meta:    meta,
		Data:    data,
	}

	if response.Data == nil {
		response.Data = http.StatusText(status)
	}

	if v, ok := data.(error); ok {
		response.Data = v.Error()
	}

	return response
}

func OK(ctx *gin.Context, data interface{}, meta map[string]interface{}) {
	r := new(http.StatusOK, meta, data)
	ctx.JSON(r.Status, r)
}

func BadRequest(ctx *gin.Context, data interface{}, meta map[string]interface{}) {
	zap.S().Error(data)
	r := new(http.StatusBadRequest, meta, data)
	ctx.JSON(r.Status, r)
}

func Unauthorized(ctx *gin.Context, data interface{}, meta map[string]interface{}) {
	zap.S().Error(data)
	r := new(http.StatusUnauthorized, meta, data)
	ctx.JSON(r.Status, r)
}

func Forbidden(ctx *gin.Context, data interface{}, meta map[string]interface{}) {
	zap.S().Error(data)
	r := new(http.StatusForbidden, meta, data)
	ctx.JSON(r.Status, r)
}

func NotFound(ctx *gin.Context, data interface{}, meta map[string]interface{}) {
	zap.S().Error(data)
	r := new(http.StatusNotFound, meta, data)
	ctx.JSON(r.Status, r)
}

func ServerError(ctx *gin.Context, data interface{}, meta map[string]interface{}) {
	zap.S().Error(data)
	r := new(http.StatusInternalServerError, meta, data)
	ctx.JSON(r.Status, r)
}

func ResponseByCode(ctx *gin.Context, code int, data interface{}, meta map[string]interface{}) {
	zap.S().Info(data)
	zap.S().Info("code: " + strconv.Itoa(code))
	r := new(code, meta, data)
	ctx.JSON(r.Status, r)
}
