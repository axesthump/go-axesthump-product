package reservationchecker

import (
	"context"
	"log"
	"route256/loms/internal/models"
	"route256/loms/internal/pool"
	"time"
)

type Repository interface {
	CancelOrder(ctx context.Context, orderID int64) error
	GetAwaitingPaymentOrdersIDs(ctx context.Context) ([]models.OrderTimestamp, error)
}

type ReservationChecker struct {
	ctx        context.Context
	ticker     *time.Ticker
	repository Repository
	pool       *pool.WorkerPool
}

func New(ctx context.Context, repository Repository) *ReservationChecker {
	checker := &ReservationChecker{
		ctx:        ctx,
		ticker:     time.NewTicker(time.Minute),
		repository: repository,
		pool:       pool.New(ctx, repository, 10),
	}
	go checker.start()
	go checker.listenErr()
	return checker
}

func (r *ReservationChecker) start() {
	defer r.ticker.Stop()
	for {
		_, ok := <-r.ticker.C
		if !ok {
			return // выходим по отключению тикера извне
		}
		orders, err := r.repository.GetAwaitingPaymentOrdersIDs(r.ctx)
		if err != nil {
			log.Println("Cant get AwaitingPaymentOrders:", err.Error()) // логаем что не смогли получить, через тик все равно пробуем еще раз
		}
		if len(orders) != 0 {
			r.pool.Submit(orders)
		}
	}
}

func (r *ReservationChecker) listenErr() {
	for {
		err, ok := <-r.pool.Err
		if !ok {
			return
		}
		log.Println("Err canceling order:", err.Error()) // через тик будем пробовать еще раз
	}
}

func (r *ReservationChecker) Stop() {
	r.ticker.Stop()
}
