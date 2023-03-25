package checkout

import (
	"context"
	"errors"
	"route256/checkout/internal/domain/checkout/mocks"
	"route256/checkout/internal/models"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestService_AddToCart_StocksChecker_Stocks_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	sChecker := mocks.NewStocksCheckerMock(mc)

	sChecker.StocksMock.Expect(ctx, 1).Return(nil, errors.New("err"))

	s := &Service{
		stocksChecker: sChecker,
	}

	err := s.AddToCart(ctx, 777, 1, 2)
	require.Error(t, err)
}

func TestService_AddToCart_Not_Enough_Items(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	sChecker := mocks.NewStocksCheckerMock(mc)

	sChecker.StocksMock.Expect(ctx, 1).Return([]models.Stock{}, nil)

	s := &Service{
		stocksChecker: sChecker,
	}

	err := s.AddToCart(ctx, 777, 1, 2)
	require.Error(t, err)
}

func TestService_AddToCart_Repository_Add_To_Cart_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	sChecker := mocks.NewStocksCheckerMock(mc)
	repo := mocks.NewRepositoryMock(mc)

	sChecker.StocksMock.Expect(ctx, 1).Return([]models.Stock{
		{
			WarehouseID: 1,
			Count:       2,
		},
	}, nil)
	repo.AddToCartMock.Expect(ctx, 777, 1, 2).Return(errors.New("err"))

	s := &Service{
		stocksChecker: sChecker,
		repository:    repo,
	}

	err := s.AddToCart(ctx, 777, 1, 2)
	require.Error(t, err)
}

func TestService_AddToCart_Success(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	sChecker := mocks.NewStocksCheckerMock(mc)
	repo := mocks.NewRepositoryMock(mc)

	sChecker.StocksMock.Expect(ctx, 1).Return([]models.Stock{
		{
			WarehouseID: 1,
			Count:       2,
		},
	}, nil)
	repo.AddToCartMock.Expect(ctx, 777, 1, 2).Return(nil)

	s := &Service{
		stocksChecker: sChecker,
		repository:    repo,
	}

	err := s.AddToCart(ctx, 777, 1, 2)
	require.NoError(t, err)
}
