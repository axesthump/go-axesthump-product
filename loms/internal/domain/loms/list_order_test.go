package loms

import (
	"context"
	"errors"
	"route256/loms/internal/domain/loms/mocks"
	"route256/loms/internal/models"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestService_ListOrder_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.ListOrderMock.Expect(ctx, 1).Return(models.OrderInfo{}, errors.New("err"))

	s := &Service{
		repository: repo,
	}

	got, err := s.ListOrder(ctx, 1)

	require.Equal(t, models.OrderInfo{}, got)
	require.Error(t, err)
}

func TestService_ListOrder_Success(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.ListOrderMock.Expect(ctx, 1).Return(models.OrderInfo{
		Status: 1,
		User:   2,
		Items: []models.Item{
			{
				Sku:   1,
				Count: 2,
			},
		},
	}, nil)

	s := &Service{
		repository: repo,
	}

	got, err := s.ListOrder(ctx, 1)

	require.Equal(t, models.OrderInfo{
		Status: 1,
		User:   2,
		Items: []models.Item{
			{
				Sku:   1,
				Count: 2,
			},
		},
	}, got)
	require.NoError(t, err)
}
