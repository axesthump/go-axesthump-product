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

func TestService_OrderPayed(t *testing.T) {
	tests := []struct {
		name    string
		status  models.OrderStatusID
		wantErr bool
	}{
		{
			name:    "OrderPayed with status New and repository err",
			status:  models.New,
			wantErr: true,
		},
		{
			name:    "OrderPayed with status AwaitingPayment and repository success",
			status:  models.AwaitingPayment,
			wantErr: false,
		},
		{
			name:    "OrderPayed with status Failed",
			status:  models.Failed,
			wantErr: true,
		},
		{
			name:    "OrderPayedr with status Payed",
			status:  models.Payed,
			wantErr: true,
		},
		{
			name:    "OrderPayed with status Cancelled",
			status:  models.Cancelled,
			wantErr: true,
		},
		{
			name:    "OrderPayed with status Cancelled",
			status:  models.UnknownOrderStatusID,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := minimock.NewController(t)
			ctx := context.Background()
			repo := mocks.NewRepositoryMock(mc)

			repo.GetStatusMock.Expect(ctx, 1).Return(tt.status, nil)
			repo.OrderPayedMock.Expect(ctx, 1).Return(nil)

			s := &Service{
				repository: repo,
			}
			if err := s.OrderPayed(ctx, 1); (err != nil) != tt.wantErr {
				t.Errorf("CancelOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_OrderPayed_Repository_GetStatus_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.GetStatusMock.Expect(ctx, 1).Return(models.UnknownOrderStatusID, errors.New("err"))

	s := &Service{
		repository: repo,
	}

	err := s.OrderPayed(ctx, 1)
	require.Error(t, err)
}

func TestService_OrderPayed_Repository_OrderPayed_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.GetStatusMock.Expect(ctx, 1).Return(models.AwaitingPayment, errors.New("err"))
	repo.OrderPayedMock.Expect(ctx, 1).Return(errors.New("err"))
	s := &Service{
		repository: repo,
	}

	err := s.OrderPayed(ctx, 1)
	require.Error(t, err)
}

func TestService_OrderPayed_Success(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.GetStatusMock.Expect(ctx, 1).Return(models.AwaitingPayment, nil)
	repo.OrderPayedMock.Expect(ctx, 1).Return(nil)
	s := &Service{
		repository: repo,
	}

	err := s.OrderPayed(ctx, 1)
	require.NoError(t, err)
}
