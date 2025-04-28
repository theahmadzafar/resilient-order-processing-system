package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/transport/http"
)

var tag = "no tag"

type metaHandler struct {
}

func NewMetaHandler() http.Handler {
	return &metaHandler{}
}

func (h *metaHandler) Register(route *gin.RouterGroup) {
	route.GET("health", h.health)
	route.GET("info", h.info)
}
func (h *metaHandler) health(ctx *gin.Context) {
	http.OK(ctx, struct{ Success string }{Success: "ok"}, nil)
}

func (h *metaHandler) info(ctx *gin.Context) {
	http.OK(ctx, struct{ Tag string }{Tag: tag}, nil)
}
