package loms

import (
	"context"
	"route256/loms/internal/models"
)

func (s *Service) CreateOrder(ctx context.Context, order models.OrderData) (int64, error) {
	orderID, err := s.repository.CreateOrder(ctx, order)
	if err != nil {
		return 0, err
	}

	//todo резервация пока создается в отдельной горутине без обработок ошибок и прочего (будет на следующих воркшопах)
	go s.repository.ReservedItems(context.Background(), orderID)
	return orderID, nil
}
