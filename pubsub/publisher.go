package pubsub

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"orderservice/models"

	"cloud.google.com/go/pubsub"
)

var topic *pubsub.Topic

func InitPublisher() {
	projectID := os.Getenv("GCP_PROJECT")
	topicID := os.Getenv("ORDER_TOPIC_ID")
	client, err := pubsub.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatal("PubSub init error:", err)
	}
	topic = client.Topic(topicID)
}

func PublishOrder(ctx context.Context, order *models.Order) {
	data, _ := json.Marshal(order)
	res := topic.Publish(ctx, &pubsub.Message{Data: data})
	if _, err := res.Get(ctx); err != nil {
		log.Println("Publish error:", err)
	}
}
