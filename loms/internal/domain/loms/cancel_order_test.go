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

func TestService_CancelOrder(t *testing.T) {
	tests := []struct {
		name    string
		status  models.OrderStatusID
		wantErr bool
	}{
		{
			name:    "Cancel order with status New and repository err",
			status:  models.New,
			wantErr: true,
		},
		{
			name:    "Cancel order with status New and repository success",
			status:  models.New,
			wantErr: false,
		},
		{
			name:    "Cancel order with status AwaitingPayment and repository err",
			status:  models.AwaitingPayment,
			wantErr: true,
		},
		{
			name:    "Cancel order with status AwaitingPayment and repository success",
			status:  models.AwaitingPayment,
			wantErr: false,
		},
		{
			name:    "Cancel order with status Failed",
			status:  models.Failed,
			wantErr: true,
		},
		{
			name:    "Cancel order with status Payed",
			status:  models.Payed,
			wantErr: true,
		},
		{
			name:    "Cancel order with status Cancelled",
			status:  models.Cancelled,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := minimock.NewController(t)
			ctx := context.Background()
			repo := mocks.NewRepositoryMock(mc)

			if tt.wantErr {
				repo.GetStatusMock.Expect(ctx, 1).Return(tt.status, nil)
				repo.CancelOrderMock.Expect(ctx, 1).Return(errors.New("err"))
			} else {
				repo.GetStatusMock.Expect(ctx, 1).Return(tt.status, nil)
				repo.CancelOrderMock.Expect(ctx, 1).Return(nil)
			}

			s := &Service{
				repository: repo,
			}
			if err := s.CancelOrder(ctx, 1); (err != nil) != tt.wantErr {
				t.Errorf("CancelOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_CancelOrder_Repository_GetStatus_Err(t *testing.T) {
	mc := minimock.NewController(t)
	ctx := context.Background()
	repo := mocks.NewRepositoryMock(mc)
	repo.GetStatusMock.Expect(ctx, 1).Return(models.UnknownOrderStatusID, errors.New("err"))

	s := &Service{
		repository: repo,
	}

	err := s.CancelOrder(ctx, 1)
	require.Error(t, err)
}
