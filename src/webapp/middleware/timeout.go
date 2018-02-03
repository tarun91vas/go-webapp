package middleware

import (
	"context"
	"net/http"
	"time"
)

// TimeoutMiddleware timeouts a request
type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm *TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}

	//Get exiting context
	ctx := r.Context()

	//Create a context with timeout
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	//Set new context in the request
	r.WithContext(ctx)

	//Handle request
	ch := make(chan struct{})
	go func() {
		tm.Next.ServeHTTP(w, r)
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		return
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	}

}
