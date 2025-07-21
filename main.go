package main

import (
	"orderservice/db"
	"orderservice/middleware"
	"orderservice/pubsub"
	router "orderservice/routes"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	godotenv.Load()

	db.InitDBConnection()
	pubsub.InitPublisher()
	app.UseRouter(middleware.AuthMiddleware())

	router.RegisterOrderRoutes(app)
	app.Configure(iris.WithLogLevel("debug"))
	app.Listen(":8081")
}
