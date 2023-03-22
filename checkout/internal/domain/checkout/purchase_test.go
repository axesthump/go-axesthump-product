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

func TestService_Purchase_Repository_ListCart_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.ListCartMock.Expect(ctx, 777).Return(nil, errors.New("err"))

	s := &Service{
		repository: repo,
	}

	got, err := s.Purchase(ctx, 777)
	require.Equal(t, int64(0), got)
	require.Error(t, err)
}

func TestService_Purchase_createOrderChecker_CreateOrder_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	createOrderChecker := mocks.NewCreateOrderCheckerMock(mc)

	repo.ListCartMock.Expect(ctx, 777).Return([]models.Item{}, nil)
	createOrderChecker.CreateOrderMock.Expect(ctx, 777, []models.CreateOrderItem{}).Return(0, errors.New("err"))

	s := &Service{
		repository:         repo,
		createOrderChecker: createOrderChecker,
	}

	got, err := s.Purchase(ctx, 777)
	require.Equal(t, int64(0), got)
	require.Error(t, err)
}

func TestService_Purchase_Repository_ClearCart_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	createOrderChecker := mocks.NewCreateOrderCheckerMock(mc)

	repo.ListCartMock.Expect(ctx, 777).Return([]models.Item{}, nil)
	repo.ClearCartMock.Expect(ctx, 777).Return(errors.New("err"))
	createOrderChecker.CreateOrderMock.Expect(ctx, 777, []models.CreateOrderItem{}).Return(1, nil)

	s := &Service{
		repository:         repo,
		createOrderChecker: createOrderChecker,
	}

	got, err := s.Purchase(ctx, 777)
	require.Equal(t, int64(1), got)
	require.Error(t, err)
}

func TestService_Purchase_Success(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	createOrderChecker := mocks.NewCreateOrderCheckerMock(mc)

	repo.ListCartMock.Expect(ctx, 777).Return([]models.Item{
		{
			Sku:   1,
			Count: 2,
			Name:  "Item1",
			Price: 3000,
		},
		{
			Sku:   2,
			Count: 3,
			Name:  "Item2",
			Price: 6000,
		},
	}, nil)
	repo.ClearCartMock.Expect(ctx, 777).Return(nil)
	createOrderChecker.CreateOrderMock.Expect(ctx, 777, []models.CreateOrderItem{
		{
			Sku:   1,
			Count: 2,
		},
		{
			Sku:   2,
			Count: 3,
		},
	}).Return(1, nil)

	s := &Service{
		repository:         repo,
		createOrderChecker: createOrderChecker,
	}

	got, err := s.Purchase(ctx, 777)
	require.Equal(t, int64(1), got)
	require.NoError(t, err)
}
