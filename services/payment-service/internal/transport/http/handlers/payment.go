package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/entities"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/services"
	"github.com/theahmadzafar/resilient-order-processing-system/services/payment-service/internal/transport/http"
	"go.uber.org/zap"
)

type PaymentHandler struct {
	orderSvc *services.PaymentService
	timeout  time.Duration
}

func NewPaymentHandler(orderSvc *services.PaymentService, timeout time.Duration) *PaymentHandler {
	return &PaymentHandler{
		timeout:  timeout,
		orderSvc: orderSvc,
	}
}
func (h *PaymentHandler) Register(router *gin.RouterGroup) {
	ordersGroup := router.Group("payment")
	ordersGroup.POST("/pay", h.pay)
}

func (h *PaymentHandler) pay(ctx *gin.Context) {
	opCtx, cancel := context.WithTimeout(ctx.Request.Context(), h.timeout)
	defer cancel()

	req := entities.PaymentRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		zap.S().Error("validation failed", zap.Error(err))
		http.BadRequest(ctx, err.Error(), nil)

		return
	}

	err := h.orderSvc.Pay(opCtx, req)
	if err != nil {
		http.ServerError(ctx, err.Error(), nil)

		return
	}

	http.OK(ctx, true, nil)
}
