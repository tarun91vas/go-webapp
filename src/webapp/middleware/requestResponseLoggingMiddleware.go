package middleware

import (
	"log"
	"net/http"
)

// RequestResponseLoggingMiddleware logs all the requests and responses
type RequestResponseLoggingMiddleware struct {
	Next http.Handler
}

func (rrl *RequestResponseLoggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if rrl.Next == nil {
		rrl.Next = http.DefaultServeMux
	}

	//log the request
	log.Println("Received request: ", r.URL.Path)

	rrl.Next.ServeHTTP(w, r)

	//log the response
	log.Println("Response sent: ", w)
}
