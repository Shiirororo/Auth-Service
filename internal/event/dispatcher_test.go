package event

import (
	"context"
	"sync"
	"testing"
	"time"
)

type noopHandler struct{}

func (h noopHandler) Handle(ctx context.Context, event Event) error {
	return nil
}

func TestDispatchConcurrentRegister(t *testing.T) {
	d := NewDispatcher(make(chan Event))
	d.Register(LoginEvent, noopHandler{})

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				d.Register(LoginEvent, noopHandler{})
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				d.Dispatch(context.Background(), Event{Type: LoginEvent})
			}
		}
	}()

	wg.Wait()
}

func TestStartStopsWhenQueueClosed(t *testing.T) {
	queue := make(chan Event)
	d := NewDispatcher(queue)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{})
	go func() {
		defer close(done)
		d.Start(ctx)
	}()

	close(queue)

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("dispatcher did not stop after queue was closed")
	}
}
