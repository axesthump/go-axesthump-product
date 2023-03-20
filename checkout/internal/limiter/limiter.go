package limiter

import (
	"context"
	"time"
)

type Limiter struct {
	ch chan struct{}
	t  *time.Ticker
}

func New(countRequestLimit int) *Limiter {
	limiter := &Limiter{
		ch: make(chan struct{}, countRequestLimit),
		t:  time.NewTicker(time.Second / time.Duration(countRequestLimit)),
	}
	go limiter.resetCountRequestInSecond()
	return limiter
}

func (l *Limiter) resetCountRequestInSecond() {
	for {
		_, ok := <-l.t.C
		if !ok {
			return
		}
		l.ch <- struct{}{}
	}
}

func (l *Limiter) Wait(ctx context.Context) error {

	select {
	case <-l.ch:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (l *Limiter) Close() {
	close(l.ch)
	l.t.Stop()
}
