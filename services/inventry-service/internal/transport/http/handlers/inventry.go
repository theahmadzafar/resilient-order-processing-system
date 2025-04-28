package handlers

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/services"
	"github.com/theahmadzafar/resilient-order-processing-system/services/inventry-service/internal/transport/http"
)

type InventryHandler struct {
	inventrySrv *services.InventryService
	timeout     time.Duration
}

func NewInventryHandler(inventrySrv *services.InventryService, timeout time.Duration) *InventryHandler {
	return &InventryHandler{
		timeout: timeout,
	}
}
func (h *InventryHandler) Register(router *gin.RouterGroup) {
	inventryGroup := router.Group("inventry")
	inventryGroup.GET("/stocks", h.stocks)
}

func (h *InventryHandler) stocks(ctx *gin.Context) {
	stocks, err := h.inventrySrv.Stocks()
	if err != nil {
		http.ServerError(ctx, errors.New("something went wrong"), nil)
	}

	http.OK(ctx, stocks, nil)
}
