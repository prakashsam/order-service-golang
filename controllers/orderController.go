package controllers

import (
	"context"
	"orderservice/models"
	"orderservice/services"

	"github.com/kataras/iris/v12"
)

func CreateOrders(ctx iris.Context) {
	var orders []models.Order
	if err := ctx.ReadJSON(&orders); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid request"})
		return
	}

	createdOrders, err := services.CreateOrders(context.Background(), orders)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(createdOrders)
}

func GetOrder(ctx iris.Context) {
	orderID := ctx.Params().Get("order_id")
	order := services.GetOrder(orderID)
	if order.ID == "" {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}
	ctx.JSON(order)
}

func UpdateOrderStatus(ctx iris.Context) {
	orderID := ctx.Params().Get("order_id")
	var update models.Order
	ctx.ReadJSON(&update)
	order, err := services.UpdateOrderStatus(orderID, update.Status)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.JSON(order)
}
