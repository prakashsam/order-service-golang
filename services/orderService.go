package services

import (
	"context"
	"orderservice/caching"
	"orderservice/models"
	"orderservice/pubsub"

	"gorm.io/gorm"
)

type OrderService struct {
	DB *gorm.DB
}

func (s *OrderService) CreateOrders(ctx context.Context, orders []models.Order) ([]models.Order, error) {
	errChan := make(chan error, len(orders))

	for i := range orders {
		if err := s.DB.Create(&orders[i]).Error; err != nil {
			return nil, err
		}

		go func(order models.Order) {
			defer func() { errChan <- nil }()

			pubsub.PublishOrder(ctx, &order)
			caching.GetRedisClient().HSetData(ctx, order.ID, order)
		}(orders[i])
	}

	for i := 0; i < len(orders); i++ {
		<-errChan
	}

	return orders, nil
}

func (s *OrderService) GetOrder(ctx context.Context, orderID string) (models.Order, error) {
	var order models.Order

	if err := caching.GetRedisClient().GetData(ctx, orderID, &order); err == nil {
		return order, nil
	}

	if err := s.DB.Where("order_id = ?", orderID).First(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}

func (s *OrderService) UpdateOrderStatus(ctx context.Context, orderID string, status string) (models.Order, error) {
	var order models.Order

	if err := s.DB.First(&order, "order_id = ?", orderID).Error; err != nil {
		return order, err
	}

	order.Status = status

	if err := s.DB.Save(&order).Error; err != nil {
		caching.GetRedisClient().HSetData(ctx, order.ID, order)
		return order, err
	}

	return order, nil
}
