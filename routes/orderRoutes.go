package routes

import (
	"orderservice/controllers"

	"github.com/kataras/iris/v12"
)

func RegisterOrderRoutes(app *iris.Application) {
	order := app.Party("/orders")
	{
		order.Post("/", controllers.CreateOrders)
		order.Patch("/{order_id}/status", controllers.UpdateOrderStatus)
		order.Get("/{order_id}", controllers.GetOrder)
	}
}
