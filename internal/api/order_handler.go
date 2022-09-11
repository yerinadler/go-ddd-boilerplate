package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yerinadler/go-ddd/internal/application"
)

type OrderHandler struct {
	service *application.OrderApplicationService
}

func NewOrderHandler(g *gin.Engine, svc *application.OrderApplicationService) {
	h := OrderHandler{
		service: svc,
	}

	g.GET("/orders/:id", h.GetOrderById)
	g.POST("/orders", h.CreateOrder)
	g.PUT("/orders/:id/paid", h.MarkOrderAsPaid)
}

func (h *OrderHandler) GetOrderById(c *gin.Context) {
	ctx := c.Request.Context()
	order_id := c.Param("id")
	order, err := h.service.GetOrderById(ctx, order_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Error("500", err.Error(), "can not get an order with the requested id"))
	}
	c.JSON(http.StatusOK, Ok(order, "successfully retrieved an order"))
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	ctx := c.Request.Context()

	var body application.CreateOrderRequestDTO

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, Error("500", err.Error(), "can not create an order"))
		return
	}

	err := h.service.CreateOrder(ctx, body.CustomerId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, Error("500", err.Error(), "can not create an order"))
		return
	}

	c.JSON(http.StatusCreated, Ok(nil, "order created"))
}

func (h *OrderHandler) MarkOrderAsPaid(c *gin.Context) {
	order_id := c.Param("id")
	ctx := c.Request.Context()
	err := h.service.MarkOrderAsPaid(ctx, order_id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, Error("500", err.Error(), "can not create an order"))
		return
	}

	c.JSON(http.StatusAccepted, Ok(nil, "order with the id of "+order_id+" has been marked as paid"))
}
