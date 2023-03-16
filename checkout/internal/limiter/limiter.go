package limiter

import (
	"context"
	"sync"
	"time"
)

type Limiter struct {
	countRequestInSecond int           // кол-во запросов, стоящие в очереди
	countRequestLimit    int           // ограничение на количество запросов в секунду
	ch                   chan struct{} // очередь
	t                    *time.Ticker
	*sync.Mutex
}

func New(countRequestLimit int) *Limiter {
	limiter := &Limiter{
		countRequestLimit: countRequestLimit,
		ch:                make(chan struct{}),
		t:                 time.NewTicker(time.Second),
		Mutex:             &sync.Mutex{},
	}
	go limiter.resetCountRequestInSecond() // запускаем процесс обработки очереди
	return limiter
}

func (l *Limiter) resetCountRequestInSecond() {
	for {
		_, ok := <-l.t.C
		if !ok {
			return // выходим если таймер уже закрылся по контексту
		}
		l.Lock()                                          // локаем для проверки очереди
		needToRead := l.countRequestInSecond              // предпологаем, что нужно сдвинуть очередь на кол-во народу в ней
		if l.countRequestInSecond > l.countRequestLimit { // если кол-во народу в очереди больше, то продвигаем только необходимый нам лимит
			needToRead = l.countRequestLimit
		}
		for i := 0; i < needToRead; i++ { // продвигаем очередь
			<-l.ch
			l.countRequestInSecond--
		}
		l.Unlock() // отпускаем лок
	}
}

func (l *Limiter) Wait(ctx context.Context) error {
	l.Lock()
	l.countRequestInSecond++
	l.Unlock()

	select {
	case l.ch <- struct{}{}: // встаем в очередь
		return nil
	case <-ctx.Done(): // выходим из очередь по отвалу контекста
		l.Lock()
		l.countRequestInSecond--
		l.Unlock()
		return ctx.Err()
	}
}

func (l *Limiter) Close() {
	close(l.ch)
	l.t.Stop()
}
