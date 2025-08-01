package controllers

import (
	"orderservice/models"
	"orderservice/services"

	"github.com/kataras/iris/v12"
)

type OrderController struct {
	Service *services.OrderService
}

func (c *OrderController) CreateOrders(ctx iris.Context) {
	var orders []models.Order
	if err := ctx.ReadJSON(&orders); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid request"})
		return
	}

	createdOrders, err := c.Service.CreateOrders(ctx.Request().Context(), orders)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(createdOrders)
}

func (c *OrderController) GetOrder(ctx iris.Context) {
	orderID := ctx.Params().Get("order_id")
	order, err := c.Service.GetOrder(ctx.Request().Context(), orderID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	if order.ID == "" {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	ctx.JSON(order)
}

func (c *OrderController) UpdateOrderStatus(ctx iris.Context) {
	orderID := ctx.Params().Get("order_id")
	var update models.Order
	if err := ctx.ReadJSON(&update); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid status update request"})
		return
	}

	order, err := c.Service.UpdateOrderStatus(ctx.Request().Context(), orderID, update.Status)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.JSON(order)
}
