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

func TestService_Stocks_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.StocksMock.Expect(ctx, 1).Return(nil, errors.New("err"))

	s := &Service{
		repository: repo,
	}

	got, err := s.Stocks(ctx, 1)

	require.Nil(t, got)
	require.Error(t, err)
}

func TestService_Stocks_Success(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.StocksMock.Expect(ctx, 1).Return([]models.Stock{
		{
			WarehouseID: 1,
			Count:       2,
		},
		{
			WarehouseID: 2,
			Count:       5,
		},
	}, nil)

	s := &Service{
		repository: repo,
	}

	got, err := s.Stocks(ctx, 1)

	require.Equal(t, []models.Stock{
		{
			WarehouseID: 1,
			Count:       2,
		},
		{
			WarehouseID: 2,
			Count:       5,
		},
	}, got)
	require.NoError(t, err)
}
