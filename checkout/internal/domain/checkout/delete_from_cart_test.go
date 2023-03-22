package checkout

import (
	"context"
	"errors"
	"route256/checkout/internal/domain/checkout/mocks"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestService_DeleteFromCart_Repository_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.DeleteFromCartMock.Expect(ctx, 777, 1, 2).Return(errors.New("err"))

	s := &Service{
		repository: repo,
	}

	err := s.DeleteFromCart(ctx, 777, 1, 2)
	require.Error(t, err)
}

func TestService_DeleteFromCart_Repository_Success(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.DeleteFromCartMock.Expect(ctx, 777, 1, 2).Return(nil)

	s := &Service{
		repository: repo,
	}

	err := s.DeleteFromCart(ctx, 777, 1, 2)
	require.NoError(t, err)
}
