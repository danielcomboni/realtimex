package realtimex

import (
	"context"
	"time"
)

type Job func()

// type Scheduler struct {
// 	jobs []Job
// }

// func (s *Scheduler) Every(d time.Duration, job Job) {
// 	go func() {
// 		ticker := time.NewTicker(d)
// 		defer ticker.Stop()

// 		for range ticker.C {
// 			job()
// 		}
// 	}()
// }

// Example usage
/*
scheduler.Every(2*time.Second, func() {
	manager.Broadcast("chat", event)
	manager.Broadcast("stream", event)
})
*/

type Scheduler struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewScheduler() *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())

	return &Scheduler{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *Scheduler) Every(d time.Duration, fn func()) {
	go func() {
		t := time.NewTicker(d)
		defer t.Stop()

		for {
			select {
			case <-t.C:
				fn()
			case <-s.ctx.Done():
				return
			}
		}
	}()
}
