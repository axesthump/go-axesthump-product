package pool

import (
	"context"
	"route256/loms/internal/models"
	"sync"
	"time"
)

type OrderCanceling interface {
	CancelOrder(ctx context.Context, orderID int64) error
}

type WorkerPool struct {
	ctx            context.Context
	in             chan models.OrderTimestamp
	orderCanceling OrderCanceling
	Err            chan error
	countWorkers   uint
	wg             *sync.WaitGroup
}

func New(ctx context.Context, orderCanceling OrderCanceling, countWorkers uint) *WorkerPool {
	pool := &WorkerPool{
		ctx:            ctx,
		countWorkers:   countWorkers,
		in:             make(chan models.OrderTimestamp),
		orderCanceling: orderCanceling,
		Err:            make(chan error),
		wg:             &sync.WaitGroup{},
	}
	for i := 0; i < int(pool.countWorkers); i++ {
		go pool.work()
	}
	return pool
}

func (w *WorkerPool) Submit(orders []models.OrderTimestamp) {
	go func() {
		for _, order := range orders {
			w.in <- order
			w.wg.Add(1)
		}
	}()
}

func (w *WorkerPool) work() {
	for {
		select {
		case order, ok := <-w.in:
			if !ok {
				return
			}
			if time.Since(order.CreateAt) > time.Minute*10 { // проверяем надо ли отменять заказ
				err := w.orderCanceling.CancelOrder(w.ctx, order.ID)
				if err != nil {
					select {
					case w.Err <- err:
					default:
					} // если не получилось отменить заказ откидываем ошибку и идем дальше, на следующий заход будет повторная попытка
				}
			}
			w.wg.Done()
		case <-w.ctx.Done():
			return
		}
	}
}

func (w *WorkerPool) Close() {
	w.wg.Wait()
	close(w.Err)
	close(w.in)
}
