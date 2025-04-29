package handlers

import (
	"context"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/entities"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/services"
	"github.com/theahmadzafar/resilient-order-processing-system/services/order-service/internal/transport/http"
	"go.uber.org/zap"
)

type OrderHandler struct {
	orderSvc *services.OrderService
	timeout  time.Duration
}

func NewOrderHandler(orderSvc *services.OrderService, timeout time.Duration) *OrderHandler {
	return &OrderHandler{
		timeout: timeout,
	}
}
func (h *OrderHandler) Register(router *gin.RouterGroup) {
	ordersGroup := router.Group("orders")
	ordersGroup.POST("/place-order", h.placeOrder)
	ordersGroup.GET("/order-status", h.orderStatus)
}

func (h *OrderHandler) placeOrder(ctx *gin.Context) {
	opCtx, cancel := context.WithTimeout(ctx.Request.Context(), h.timeout)
	defer cancel()

	req := entities.OrderRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		zap.S().Error("validation failed", zap.Error(err))
		http.BadRequest(ctx, err, nil)

		return
	}

	err := h.orderSvc.PlaceOrder(opCtx, req)
	if err != nil {
		http.ServerError(ctx, errors.New("something went wrong"), nil)
	}

	http.OK(ctx, true, nil)
}

func (h *OrderHandler) orderStatus(ctx *gin.Context) {
	opCtx, cancel := context.WithTimeout(ctx.Request.Context(), h.timeout)
	defer cancel()

	req := entities.GetOrderRequest{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		zap.S().Error("validation failed", zap.Error(err))
		http.BadRequest(ctx, err, nil)

		return
	}

	res, err := h.orderSvc.GetOrder(opCtx, req)
	if err != nil {
		http.ServerError(ctx, errors.New("something went wrong"), nil)
	}

	http.OK(ctx, *res, nil)
}
