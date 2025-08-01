package main

import (
	"orderservice/caching"
	"orderservice/db"

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
	caching.InitializeRedisClient()

	router.RegisterOrderRoutes(app)
	app.Configure(iris.WithLogLevel("debug"))
	app.Listen(":8082")
}
