package boxutil

import (
	"context"
	"errors"
	"log"
	"time"
)

// ErrCanceled is returned when the context cancels or times out
// an operation.
var ErrCanceled = errors.New("context is done")

// readyCheckInterval is the interval to wait between checking
// the status in StatusChan.
// Unexported because 1 second is sensible, but configurable to make
// tests run quicker.
var readyCheckInterval = 1 * time.Second

// StatusChan gets a channel that periodically gets the box info
// and sends a message whenever the status changes.
func StatusChan(ctx context.Context, i Box) <-chan string {
	log.Println("boxutil: DEPRECATED use github.com/machinebox/sdk-go/boxutil instead")
	statusChan := make(chan string)
	go func() {
		var lastStatus string
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(readyCheckInterval)
				status := "unavailable"
				info, err := i.Info()
				if err == nil {
					status = info.Status
				}
				if status != lastStatus {
					lastStatus = status
					statusChan <- status
				}
			}
		}
	}()
	return statusChan
}

// WaitForReady blocks until the Box is ready.
func WaitForReady(ctx context.Context, i Box) error {
	log.Println("boxutil: DEPRECATED use github.com/machinebox/sdk-go/boxutil instead")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	statusChan := StatusChan(ctx, i)
	for {
		select {
		case <-ctx.Done():
			return ErrCanceled
		case status := <-statusChan:
			if IsReady(status) {
				return nil
			}
		}
	}
}

// IsReady gets whether the box info status is ready or not.
func IsReady(status string) bool {
	log.Println("boxutil: DEPRECATED use github.com/machinebox/sdk-go/boxutil instead")
	return status == "ready"
}