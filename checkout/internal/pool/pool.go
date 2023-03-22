package pool

import (
	"context"
	"route256/checkout/internal/models"
	"sync"
)

type ProductsChecker interface {
	GetProduct(ctx context.Context, sku uint32) (models.Product, error)
}

type WorkerPool struct {
	ctx            context.Context
	productChecker ProductsChecker
	Out            chan models.Item
	in             chan models.Item
	Err            chan error
	countWorkers   uint
	wg             *sync.WaitGroup
}

func New(ctx context.Context, productChecker ProductsChecker, countWorkers uint) *WorkerPool {
	pool := &WorkerPool{
		ctx:            ctx,
		productChecker: productChecker,
		countWorkers:   countWorkers,
		Out:            make(chan models.Item), // Канал для чтения item извне (вне пула)
		in:             make(chan models.Item), // Канал для передачи item внутри (внутри пула)
		Err:            make(chan error),       // Канал по которому прокидываем ошибки
		wg:             &sync.WaitGroup{},      // Для коректного закрытия каналов
	}
	for i := 0; i < int(pool.countWorkers); i++ {
		go pool.work() // запускаем воркеров
	}
	return pool
}

// Submit отправляет во внутренний канал in пришедшие item
func (w *WorkerPool) Submit(items []models.Item) {
	go func() {
		for _, item := range items {
			w.in <- item
			w.wg.Add(1)
		}
		go w.close()
	}()
}

func (w *WorkerPool) work() {
	for {
		select {
		case item, ok := <-w.in: // слушаем приходящие item
			if !ok {
				return // если канал закрыт вываливаемся
			}
			product, err := w.productChecker.GetProduct(w.ctx, item.Sku) // сходили за продуктом
			if err != nil {
				select {
				case w.Err <- err: // пытаемся записать в ерр
				default: // кейс когда, ерр уже закрыт (так как один из прошлых запросов на продукт ответил ошибкой)
				}
				w.wg.Done()
				return
			}
			item.Price = product.Price
			item.Name = product.Name
			w.Out <- item
			w.wg.Done()
		case <-w.ctx.Done():
			return
		}
	}
}

func (w *WorkerPool) close() {
	w.wg.Wait() // будет ждать пока все запущенные операции не отработали
	close(w.Err)
	close(w.Out)
	close(w.in)
}
