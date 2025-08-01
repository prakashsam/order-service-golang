package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"orderservice/models"
	"orderservice/config"

	"cloud.google.com/go/pubsub"
)

var topic *pubsub.Topic

func InitPublisher() {
	cfg := config.Load()
	client, err := pubsub.NewClient(context.Background(), cfg.ProjectID)
	if err != nil {
		log.Fatal("PubSub init error:", err)
	}
	topic = client.Topic(cfg.ORDERTOPICID)
}

func PublishOrder(ctx context.Context, order *models.Order) {
	data, _ := json.Marshal(order)
	res := topic.Publish(ctx, &pubsub.Message{Data: data})
	if _, err := res.Get(ctx); err != nil {
		log.Println("Publish error:", err)
	}
}
