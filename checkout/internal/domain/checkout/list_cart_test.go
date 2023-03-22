package checkout

import (
	"context"
	"errors"
	"reflect"
	"route256/checkout/internal/domain/checkout/mocks"
	"route256/checkout/internal/models"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestService_ListCart_Repository_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.ListCartMock.Expect(ctx, 777).Return(nil, errors.New("err"))

	s := &Service{
		repository: repo,
	}

	got, err := s.ListCart(ctx, 777)
	require.Equal(t, models.CartInfo{}, got)
	require.Error(t, err)
}

func TestService_ListCart_With_Empty_Items(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	getter := mocks.NewProductInfoGetterMock(mc)

	getter.GetProductsInfoMock.Expect(ctx, []models.Item{}).Return([]models.Item{}, nil)
	repo.ListCartMock.Expect(ctx, 777).Return([]models.Item{}, nil)

	s := &Service{
		repository:        repo,
		productInfoGetter: getter,
	}

	got, err := s.ListCart(ctx, 777)
	require.Equal(t, models.CartInfo{Items: []models.Item{}}, got)
	require.NoError(t, err)
}

func TestService_ListCart_FillItems_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	getter := mocks.NewProductInfoGetterMock(mc)

	repo.ListCartMock.Expect(ctx, 777).Return([]models.Item{}, nil)
	getter.GetProductsInfoMock.Expect(ctx, []models.Item{}).Return(nil, errors.New("err"))

	s := &Service{
		repository:        repo,
		productInfoGetter: getter,
	}

	got, err := s.ListCart(ctx, 777)
	require.Nil(t, got.Items)
	require.Error(t, err)
}

func TestService_ListCart_With_Items(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	getter := mocks.NewProductInfoGetterMock(mc)

	items := []models.Item{
		{
			Sku:   1,
			Count: 5,
		},
		{
			Sku:   2,
			Count: 3,
		},
	}

	repo.ListCartMock.Expect(ctx, 777).Return(items, nil)
	getter.GetProductsInfoMock.Expect(ctx, items).Return([]models.Item{
		{
			Sku:   1,
			Count: 5,
			Name:  "Item1",
			Price: 3000,
		},
		{
			Sku:   2,
			Count: 3,
			Name:  "Item2",
			Price: 5000,
		},
	}, nil)

	s := &Service{
		repository:        repo,
		productInfoGetter: getter,
	}

	expect := models.CartInfo{
		Items: []models.Item{
			{
				Sku:   1,
				Count: 5,
				Name:  "Item1",
				Price: 3000,
			},
			{
				Sku:   2,
				Count: 3,
				Name:  "Item2",
				Price: 5000,
			},
		},
		TotalPrice: 30000,
	}
	got, err := s.ListCart(ctx, 777)
	require.Equal(t, expect.TotalPrice, got.TotalPrice)
	require.True(t, reflect.DeepEqual(expect.Items, got.Items))
	require.NoError(t, err)
}
