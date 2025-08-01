package route

import (
	"orderservice/controllers"
	"orderservice/db"
	"orderservice/services"

	"github.com/kataras/iris/v12"
)

func RegisterOrderRoutes(app *iris.Application) {
	orderService := &services.OrderService{DB: db.GetDBConnection()}
	orderController := &controllers.OrderController{Service: orderService}

	app.Post("/orders", orderController.CreateOrders)
	app.Get("/orders/{order_id}", orderController.GetOrder)
	app.Put("/orders/{order_id}/status", orderController.UpdateOrderStatus)
}
