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

func TestService_CreateOrder_Repository_CreateOrder_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.CreateOrderMock.Expect(ctx, models.OrderData{}).Return(0, errors.New("err"))

	s := &Service{
		repository: repo,
	}

	got, err := s.CreateOrder(ctx, models.OrderData{})

	require.Equal(t, int64(0), got)
	require.Error(t, err)
}

func TestService_CreateOrder_Repository_ReservedItems_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.CreateOrderMock.Expect(ctx, models.OrderData{}).Return(1, nil)
	repo.ReservedItemsMock.Expect(ctx, 1).Return(errors.New("err"))

	s := &Service{
		repository: repo,
	}

	got, err := s.CreateOrder(ctx, models.OrderData{})

	require.Equal(t, int64(0), got)
	require.Error(t, err)
}

func TestService_CreateOrder_Success(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.CreateOrderMock.Expect(ctx, models.OrderData{}).Return(1, nil)
	repo.ReservedItemsMock.Expect(ctx, 1).Return(nil)

	s := &Service{
		repository: repo,
	}

	got, err := s.CreateOrder(ctx, models.OrderData{})

	require.Equal(t, int64(1), got)
	require.NoError(t, err)
}
