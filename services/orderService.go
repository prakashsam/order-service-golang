package services

import (
	"context"
	"orderservice/db"
	"orderservice/models"
	"orderservice/pubsub"
)

func CreateOrders(ctx context.Context, orders []models.Order) ([]models.Order, error) {
	DB := db.GetDBConnection()
	for i := range orders {
		if err := DB.Create(&orders[i]).Error; err != nil {
			return nil, err
		}
		pubsub.PublishOrder(ctx, &orders[i])
	}
	return orders, nil
}

func GetOrder(orderID string) models.Order {
	var order models.Order
	DB := db.GetDBConnection()
	DB.Where("order_id = ?", orderID).First(&order)
	return order
}

func UpdateOrderStatus(orderID string, status string) (models.Order, error) {
	var order models.Order
	DB := db.GetDBConnection()

	err := DB.First(&order, "order_id = ?", orderID).Error
	if err != nil {
		return order, err
	}

	order.Status = status

	err = DB.Save(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
